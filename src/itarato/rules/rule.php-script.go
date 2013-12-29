package rules

import (
	"itarato/types"
)

type PHPScriptRule Rule

func (rule *PHPScriptRule) Read(char byte, buffer *string, context_state *types.Stack) {

}

func (rule *PHPScriptRule) GetName() string {
	return "PHP script"
}
