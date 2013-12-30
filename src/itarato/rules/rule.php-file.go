package rules

import (
	"itarato/types"
)

type PHPFileRule struct {
	Rule
	php_script_rule *PHPScriptRule
}

func NewPHPFileRule() *PHPFileRule {
	rule := new(PHPFileRule)
	rule.php_script_rule = NewPHPScriptRule()
	return rule
}

func (rule *PHPFileRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	return true
}

func (rule *PHPFileRule) Read(char byte, buffer *string, context_state *types.Stack) {
	rule.php_script_rule.TryOn(char, buffer, context_state)
}

func (rule *PHPFileRule) GetName() string {
	return "PHP file"
}
