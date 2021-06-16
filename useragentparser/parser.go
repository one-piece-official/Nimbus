package useragentparser

import (
	"net/url"
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
	wg.Add(1)

	go func() {
		defer wg.Done()

		userAgent.Os = parser.parseOs(userAgentString)
	}()
	wg.Wait()

	return
}

func (parser *userAgentParser) parseDevice(agentString string) *Device {
	if strings.Contains(agentString, "%") {
		originAgentString, err := url.QueryUnescape(agentString)
		if err == nil {
			agentString = originAgentString
		}
	}

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

type UserAgentParser interface {
	Parse(userAgentString string) (userAgent *UserAgent)
}
