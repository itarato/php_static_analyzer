package rules

import (
	"itarato/types"
)

type PHPScriptRule struct {
	Rule
	function_rule *FunctionRule
}

func NewPHPScriptRule() *PHPScriptRule {
	rule := new(PHPScriptRule)
	return rule
}

func (rule *PHPScriptRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	if types.IsEqualWithAClosure(*buffer, "<?php", types.WHITESPACE_ALL) {
		// @todo use proper states.
		context_state.Push(rule)

		// Clear buffer.
		*buffer = ""

		rule.function_rule = NewFunctionRule()

		return true
	}

	return false
}

func (rule *PHPScriptRule) Read(char byte, buffer *string, context_state *types.Stack) {
	rule.function_rule.TryOn(char, buffer, context_state)

	if char == '}' {
		// Maybe it should be a variable escape token: "}" or "?>" ...
		*buffer = ""
		context_state.Pop()
		context_state.Pop()
	}

	// @todo add later closing php tag.
}

func (rule *PHPScriptRule) GetName() string {
	return "PHP script"
}
