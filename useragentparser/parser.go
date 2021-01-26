package useragentparser

import (
	"regexp"
	"strings"
	"sync"
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

type deviceParser struct {
	Reg          *regexp.Regexp
	Expr         string               `yaml:"regex"`
	Brand        string               `yaml:"brand"`
	ModelParsers []*deviceModelParser `yaml:"models"`
}

type osParser struct {
	Reg     *regexp.Regexp
	Expr    string `yaml:"regex"`
	Family  string `yaml:"family"`
	Version string `yaml:"version"`
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
}

func NewUserAgentParser() UserAgentParser {
	return &userAgentParser{
		OsParsers:     defaultOsParsers(),
		BotParsers:    defaultBotParsers(),
		DeviceParsers: defaultDeviceParsers(),
	}
}

func (parser *userAgentParser) Parse(userAgentString string) (userAgent *UserAgent) {
	userAgent = &UserAgent{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		userAgent.Device = parser.parseDevice(userAgentString)
	}()
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

func (parser *deviceParser) Match(userAgentString string, device *Device) {
	matches := parser.Reg.FindStringSubmatchIndex(userAgentString)

	if len(matches) == 0 {
		return
	}

	device.Brand = parser.Brand

	for _, modelParser := range parser.ModelParsers {
		modelMatches := modelParser.Reg.FindStringSubmatchIndex(userAgentString)
		if len(modelMatches) > 0 {
			device.Model = string(parser.Reg.ExpandString(nil, modelParser.Model, userAgentString, matches))
			device.Model = strings.TrimSpace(device.Model)

			break
		}
	}
}

type UserAgentParser interface {
	Parse(userAgentString string) (userAgent *UserAgent)
}
