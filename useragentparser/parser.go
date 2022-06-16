package useragentparser

import (
	"net/url"
	"regexp"
	"strings"
	"sync"
)

const (
	brandVivo   = "vivo"
	brandXiaomi = "Xiaomi"
	brandOppo   = "OPPO"
	brandHuawei = "Huawei"
	brandApple  = "Apple"
	brandOther  = "Other"
)

type UserAgent struct {
	Device *Device
	Os     *Os
	Bot    *Bot
}

type deviceModelParser struct {
	Reg   *regexp.Regexp
	Expr  string `yaml:"regex"`
	Model string `yaml:"model"`
}

type botParser struct {
	Reg    *regexp.Regexp
	Expr   string `yaml:"regex"`
	Family string `yaml:"family"`
}

type userAgentParser struct {
	OsParsers     []*osParser     `yaml:"oss"`
	BotParsers    []*botParser    `yaml:"bots"`
	DeviceParsers []*deviceParser `yaml:"devices"`
	ModelMapping  map[string]string
	FuzzyMapping  map[string]string
}

func NewUserAgentParser() UserAgentParser {
	return &userAgentParser{
		OsParsers:     defaultOsParsers(),
		BotParsers:    defaultBotParsers(),
		DeviceParsers: defaultDeviceParsers(),
		ModelMapping:  modelMapping(),
		FuzzyMapping:  fuzzyBrandMapping(),
	}
}

func (parser *userAgentParser) Parse(userAgentString string) (userAgent *UserAgent) {
	if strings.Contains(userAgentString, "%") {
		originAgentString, err := url.QueryUnescape(userAgentString)
		if err == nil {
			userAgentString = originAgentString
		}
	}

	userAgent = parser.preDetect(userAgentString)
	// userAgent = &UserAgent{}

	var wg sync.WaitGroup

	if userAgent.Device == nil {
		wg.Add(1)

		go func() {
			defer wg.Done()

			userAgent.Device = parser.parseDevice(userAgentString)
		}()
	}

	if userAgent.Os == nil {
		wg.Add(1)

		go func() {
			defer wg.Done()

			userAgent.Os = parser.parseOs(userAgentString)
		}()
	}

	wg.Wait()

	return
}

func (parser *userAgentParser) parseDevice(agentString string) *Device {
	dvc := new(Device)

	found := false

	for _, dvcPattern := range parser.DeviceParsers {
		dvcPattern.Match(agentString, dvc)

		if len(dvc.Brand) > 0 {
			found = true

			break
		}
	}

	if !found {
		dvc.Brand = "Other"
	}

	return dvc
}

func (parser *userAgentParser) parseOs(agentString string) *Os {
	os := new(Os)
	found := false

	for _, osPattern := range parser.OsParsers {
		osPattern.Match(agentString, os)

		if len(os.Family) > 0 {
			found = true

			break
		}
	}

	if !found {
		os.Family = "Other"
	}

	return os
}

// nolint
func (parser *userAgentParser) preDetect(agentString string) (userAgent *UserAgent) {
	//
	if strings.Contains(agentString, "iPhone") {
		return &UserAgent{
			Device: &Device{Brand: brandApple, Model: ""},
			Os:     &Os{},
		}
	}

	if strings.Contains(agentString, "Windows NT") {
		return &UserAgent{
			Device: &Device{Brand: brandOther, Model: ""},
			Os:     &Os{},
		}
	}

	userAgent = &UserAgent{}

	var mainPart, system, model string
	if firstPart := strings.Split(agentString, "("); len(firstPart) > 1 {
		mainPart = strings.Split(firstPart[1], ")")[0]
	}
	var items []string
	if strings.HasPrefix(agentString, "AliXAdSDK") || strings.HasPrefix(agentString, "Youku") {
		mainPart = agentString
		items = strings.Split(agentString, ";")
		model = strings.ToLower(items[len(items)-1])
		system = "android " + items[len(items)-2]
	} else {
		mainPart = strings.Replace(mainPart, "Linux;", "", 1)
		mainPart = strings.Replace(mainPart, "U;", "", 1)
		mainPart = strings.ToLower(mainPart)
		mainPart = strings.Replace(mainPart, "en-us;", "", 1)
		mainPart = strings.Replace(mainPart, "zh-cn;", "", 1)
		mainPart = strings.Replace(mainPart, "zh-hans-cn;", "", 1)
		mainPart = strings.Replace(mainPart, "zh-tw;", "", 1)
		mainPart = strings.Replace(mainPart, "zh-mo;", "", 1)
		mainPart = strings.Replace(mainPart, "zh-hk;", "", 1)
		mainPart = strings.TrimSpace(mainPart)

		items = strings.Split(mainPart, ";")

		if len(items) < 2 {
			model = items[0]
		} else if strings.HasPrefix(agentString, "MomoChat") {
			system = items[1]
			model = strings.TrimSpace(strings.Split(items[0], " build")[0])
		} else if items[0] == "android" {
			if len(items) > 2 {
				system = "android " + strings.Split(items[1], ".")[0]
				model = strings.TrimSpace(strings.Split(items[2], " build")[0])
			}
		} else {
			system = items[0]
			model = strings.TrimSpace(strings.Split(items[1], " build")[0])
		}

		model = strings.Split(model, ",")[0]
	}

	if strings.HasPrefix(agentString, "ting") {
		if strings.Split(mainPart, "android")[1] >= "29" {
			system = "android 10"
		} else {
			system = "android 9"
		}
	}

	if systemItems := strings.Split(system, "android"); len(systemItems) > 1 {
		userAgent.Os = &Os{
			Family:  "Android",
			Version: strings.Split(strings.TrimSpace(systemItems[1]), ".")[0],
		}
	}

	userAgent.Device = parser.modelStrToDevice(model)

	//if userAgent.Device == nil {
	//	fmt.Println(model, mainPart, agentString)
	//}

	return
}

// nolint
func (parser *userAgentParser) modelStrToDevice(model string) *Device {
	if brand, ok := parser.ModelMapping[model]; ok {
		return &Device{
			Brand: brand,
			Model: model,
		}
	}

	for k, v := range parser.FuzzyMapping {
		if strings.Contains(model, k) {
			return &Device{
				Brand: v,
				Model: model,
			}
		}
	}

	model = strings.Split(model, " ")[0]

	if len(model) >= 6 && len(model) <= 7 && model[0] == 'v' &&
		isNum(model[1]) && isNum(model[2]) && isEn(model[len(model)-1]) {
		return &Device{
			Brand: brandVivo,
			Model: model,
		}
	}

	if len(model) == 6 && model[0] == 'p' && isNum(model[5]) {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	if len(model) == 5 && model[0] == 'r' && isNum(model[1]) && model[2] == '0' && model[3] == '0' {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	if strings.HasPrefix(model, "x90") || strings.HasPrefix(model, "n5") {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	if len(model) == 7 && strings.HasPrefix(model, "rmx") && isNum(model[len(model)-1]) {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	if len(model) == 3 && model[0] == 'p' && isNum(model[1]) && model[2] == '0' {
		return &Device{
			Brand: brandHuawei,
			Model: model,
		}
	}

	if len(model) >= 8 && len(model) <= 10 && model[0] == 'm' && model[1] == '2' && isNum(model[2]) && isEn(model[len(model)-1]) {
		return &Device{
			Brand: brandXiaomi,
			Model: model,
		}
	}

	if len(model) >= 7 && len(model) <= 10 && model[0] == '2' && model[1] <= '2' && isEn(model[len(model)-1]) {
		return &Device{
			Brand: brandXiaomi,
			Model: model,
		}
	}

	if len(model) >= 7 &&
		((model[3] == '-' && isNum(model[6])) || (len(model) > 7 && model[4] == '-' && isNum(model[7]))) {
		return &Device{
			Brand: brandHuawei,
			Model: model,
		}
	}

	if len(model) == 6 && isEn(model[0]) && isEn(model[1]) && isNum(model[2]) && isNum(model[5]) {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	if strings.HasPrefix(model, "cph") {
		return &Device{
			Brand: brandOppo,
			Model: model,
		}
	}

	return nil
}

func isNum(i byte) bool {
	return i <= '9' && i >= '0'
}

func isEn(i byte) bool {
	return i <= 'z' && i >= 'a'
}

type UserAgentParser interface {
	Parse(userAgentString string) (userAgent *UserAgent)
}
