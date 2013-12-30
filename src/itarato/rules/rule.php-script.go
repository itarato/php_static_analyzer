package rules

import (
	"itarato/types"
)

type PHPScriptRule Rule

var (
	function_rule *FunctionRule = &FunctionRule{}
)

func (rule *PHPScriptRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	if types.IsEqualWithAClosure(*buffer, "<?php", types.WHITESPACE_ALL) {
		// @todo use proper states.
		context_state.Push(rule)

		// Clear buffer.
		*buffer = ""

		return true
	}

	// @todo add later closing php tag.

	return false
}

func (rule *PHPScriptRule) Read(char byte, buffer *string, context_state *types.Stack) {
	function_rule.TryOn(char, buffer, context_state)
}

func (rule *PHPScriptRule) GetName() string {
	return "PHP script"
}
