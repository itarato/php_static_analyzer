package rules

import (
	"itarato/types"
)

type PHPFileRule Rule

var (
	php_script_rule *PHPScriptRule = &PHPScriptRule{}
)

func (rule *PHPFileRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	return true
}

func (rule *PHPFileRule) Read(char byte, buffer *string, context_state *types.Stack) {
	php_script_rule.TryOn(char, buffer, context_state)
}

func (rule *PHPFileRule) GetName() string {
	return "PHP file"
}
