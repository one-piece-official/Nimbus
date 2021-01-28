package useragentparser

import (
	"regexp"
	"strings"
)

type Device struct {
	Brand string
	Model string
}

type deviceParser struct {
	Reg          *regexp.Regexp
	Expr         string               `yaml:"regex"`
	Brand        string               `yaml:"brand"`
	ModelParsers []*deviceModelParser `yaml:"models"`
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
