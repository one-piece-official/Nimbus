package useragentparser

import (
	"regexp"
	"strings"
)

type Os struct {
	Family  string
	Version string
}

type osParser struct {
	Reg     *regexp.Regexp
	Expr    string `yaml:"regex"`
	Family  string `yaml:"family"`
	Version string `yaml:"version"`
}

func (parser *osParser) Match(userAgentString string, os *Os) {
	matches := parser.Reg.FindStringSubmatchIndex(userAgentString)

	if len(matches) == 0 {
		return
	}

	os.Family = parser.Family
	os.Version = string(parser.Reg.ExpandString(nil, parser.Version, userAgentString, matches))
	os.Version = strings.TrimSpace(os.Version)
}
