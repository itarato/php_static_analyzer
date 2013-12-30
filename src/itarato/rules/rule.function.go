package rules

import (
	"itarato/types"
	"strings"
)

type FunctionRule Rule

func (rule *FunctionRule) TryOn(char byte, buffer *string, context_state *types.Stack) bool {
	buffer_trimmed := strings.Trim(*buffer, "\n\r\t ")

	if buffer_trimmed == "function" && types.IsWhitespace(char) {
		// It's a function definition.
		context_state.Push(rule)

		*buffer = ""
	}

	return false
}

func (rule *FunctionRule) Read(char byte, buffer *string, context_state *types.Stack) {

}

func (rule *FunctionRule) GetName() string {
	return "Function"
}
