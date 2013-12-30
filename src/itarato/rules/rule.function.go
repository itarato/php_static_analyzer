package rules

import (
	"itarato/types"
	"log"
	"strings"
)

type FunctionRule struct {
	Rule
	php_script_rule *PHPScriptRule
	state           string
	name            string
}

const (
	STATE_INIT        = "init"
	STATE_NAME        = "name"
	STATE_ARGS        = "args"
	STATE_BLOCK_BEGIN = "block_begin"
	STATE_BLOCK       = "block"
)

func NewFunctionRule() *FunctionRule {
	rule := new(FunctionRule)
	rule.state = STATE_INIT
	return rule
}

func (rule *FunctionRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	if types.IsEqualWithAClosure(*buffer, "function", types.WHITESPACE_ALL) {
		// It's a function definition.
		context_state.Push(rule)
		*buffer = ""
		rule.state = STATE_NAME

		rule.php_script_rule = NewPHPScriptRule()

		return true
	}

	return false
}

func (rule *FunctionRule) Read(char byte, buffer *string, context_state *types.Stack) {
	if rule.state == STATE_NAME {
		if char == '(' {
			rule.name = strings.Trim(*buffer, types.WHITESPACE_ALL+"(")
			log.Println("Function name:", rule.name)
			rule.state = STATE_ARGS
			*buffer = ""

			return
		}
	}

	if rule.state == STATE_ARGS {
		if char == ')' {
			*buffer = ""
			rule.state = STATE_BLOCK_BEGIN
		}
	}

	if rule.state == STATE_BLOCK_BEGIN {
		if char == '{' {
			*buffer = ""
			rule.state = STATE_BLOCK
			context_state.Push(rule.php_script_rule)
		}
	}

	if rule.state == STATE_BLOCK {
		rule.php_script_rule.TryOn(char, buffer, context_state)
	}
}

func (rule *FunctionRule) GetName() string {
	return "Function (" + rule.state + ")"
}
