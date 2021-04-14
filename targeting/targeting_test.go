package targeting_test

import (
	"testing"

	"github.com/one-piece-official/Nimbus/targeting"
	"github.com/stretchr/testify/assert"
)

// 定向基础 string.
func TestTargetingSingleContain(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{"a": "a", "b": "banana"}

	rule := &targeting.Rule{
		AggregateType: "and",
		Rules: []targeting.Rule{
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "b",
					Operator: "contain",
					Values:   []string{"bana"},
				},
			},
		},
		RuleItem: nil,
	}
	match, err := rule.Match(input)
	assert.Equal(t, true, match)
	assert.Nil(t, err)
}

// 定向两个 key 各自包含某 string.
func TestTargetingMultipleAndContain(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{"a": "apple", "b": "banana"}
	rule := &targeting.Rule{
		AggregateType: "and",
		Rules: []targeting.Rule{
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "b",
					Operator: "contain",
					Values:   []string{"bana"},
				},
			},
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "a",
					Operator: "contain",
					Values:   []string{"pp"},
				},
			},
		},
		RuleItem: nil,
	}
	match, err := rule.Match(input)
	assert.Equal(t, true, match)
	assert.Nil(t, err)
}

// 定向两个 key 其中一个包含某 string.
func TestTargetingMultipleOrContain(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{"a": "apple", "b": "apple"}
	rule := &targeting.Rule{
		AggregateType: "and",
		Rules: []targeting.Rule{
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "b",
					Operator: "contain",
					Values:   []string{"bana"},
				},
			},
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "a",
					Operator: "contain",
					Values:   []string{"pp"},
				},
			},
		},
		RuleItem: nil,
	}
	match, err := rule.Match(input)
	assert.Equal(t, false, match)
	assert.Nil(t, err)
}

// 定向两个 key 其中之一==某 string.
func TestTargetingMultipleOrEqual(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{"a": "apple", "b": "banana"}
	rule := &targeting.Rule{
		AggregateType: "or",
		Rules: []targeting.Rule{
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "b",
					Operator: "contain",
					Values:   []string{"banana"},
				},
			},
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "equal",
					Key:      "a",
					Operator: "contain",
					Values:   []string{"apple"},
				},
			},
		},
		RuleItem: nil,
	}
	match, err := rule.Match(input)
	assert.Equal(t, true, match)
	assert.Nil(t, err)

	input["a"] = ""
	match, err = rule.Match(input)
	assert.Equal(t, true, match)
	assert.Nil(t, err)
}

// 定向两个 key 其中一个包含某 string.
func TestTargetingMultipleAndWithMultipleAndContain(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{"a": "apple", "b": "bana", "c": "tt"}
	rule := &targeting.Rule{
		AggregateType: "and",
		Rules: []targeting.Rule{
			{
				AggregateType: "and",
				Rules: []targeting.Rule{
					{
						AggregateType: "",
						Rules:         nil,
						RuleItem: &targeting.RuleItem{
							Type:     "include",
							Key:      "b",
							Operator: "contain",
							Values:   []string{"bana"},
						},
					},
					{
						AggregateType: "",
						Rules:         nil,
						RuleItem: &targeting.RuleItem{
							Type:     "include",
							Key:      "c",
							Operator: "contain",
							Values:   []string{"tt"},
						},
					},
				},
			},
			{
				AggregateType: "",
				Rules:         nil,
				RuleItem: &targeting.RuleItem{
					Type:     "include",
					Key:      "a",
					Operator: "contain",
					Values:   []string{"pp"},
				},
			},
		},
		RuleItem: nil,
	}
	match, err := rule.Match(input)
	assert.Equal(t, true, match)
	assert.Nil(t, err)

	delete(input, "c")
	match, err = rule.Match(input)
	assert.Equal(t, false, match)
	assert.NotNil(t, err)
}
