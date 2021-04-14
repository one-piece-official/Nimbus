package targeting

import (
	"fmt"
	"strings"
)

const (
	valueOperatorEqual   = "equal"
	valueOperatorContain = "contain"
	ruleTypeInclude      = "include"
	ruleTypeExclude      = "exclude"
	ruleAggregateTypeAnd = "and"
	ruleAggregateTypeOr  = "or"
)

var (
	errWrongAggregateType = fmt.Errorf("wrong aggregate type")
	errNoInputKeyValue    = fmt.Errorf("no input key value")
	errInputNotString     = fmt.Errorf("input is not string")
)

type RuleItem struct {
	Type     string   `json:"type"`
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type Rule struct {
	AggregateType string    `json:"aggregate_type"` // 列表聚合类型 AND OR
	Rules         []Rule    `json:"rules"`          // 子列表
	RuleItem      *RuleItem `json:"rule"`           // 具体规则
}

func (l Rule) Match(input map[string]interface{}) (match bool, err error) {
	if l.RuleItem != nil {
		return l.RuleItem.Match(input)
	} else if l.Rules != nil {
		for _, rule := range l.Rules {
			match, err = rule.Match(input)
			switch l.AggregateType {
			case ruleAggregateTypeAnd:
				if !match {
					return match, err
				}
			case ruleAggregateTypeOr:
				if match {
					return match, err
				}
			default:
				return false, errWrongAggregateType
			}
		}
	}

	return true, nil
}

func (lt RuleItem) Match(input map[string]interface{}) (match bool, err error) {
	value, ok := input[lt.Key]
	if !ok {
		return false, errNoInputKeyValue
	}

	strValue, ok := value.(string)
	if !ok {
		return false, errInputNotString
	}

	for _, target := range lt.Values {
		if lt.MatchValue(target, strValue) {
			return lt.Type == ruleTypeInclude, nil // 在白名单中，可以投放
		}
	}

	return lt.Type == ruleTypeExclude, nil // 不在黑名单中，可以投放
}

func (lt RuleItem) MatchValue(target, value string) bool {
	switch lt.Operator {
	case valueOperatorEqual:
		return value == target
	case valueOperatorContain:
		return strings.Contains(value, target)
	default:
		return false
	}
}
