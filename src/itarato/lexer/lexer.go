package lexer

import (
	"itarato/types"
)

const (
	T_PHP_START int = 0
	T_KEYWORD
	T_NAME
	T_VARIABLE
	T_OPERATION
	T_BRACKET
	T_WHITESPACE
	T_UNKNOWN
)

type LexerResult struct {
	Found    bool
	TypeCode int
	// Word     string
}

func Interpret(word *string, char byte) *LexerResult {
	// Current token is only whitespaces.
	if types.IsWhitespaces(*word) {
		if types.IsWhitespace(char) {
			return &LexerResult{Found: false}
		}

		return &LexerResult{true, T_WHITESPACE}
	}

	if types.IsWhitespace(char) {
		return &LexerResult{true, T_UNKNOWN}
	}

	return &LexerResult{Found: false}
}
