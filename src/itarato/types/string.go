package types

import (
	"regexp"
	"strings"
)

const (
	WHITESPACE_ALL = "\n\r\t\f\v "
)

func ReplaceWhitespaces(s string) string {
	// Replace whitespaces to something visual. Eg.: \n -> [n].
	for _, pattern := range []string{"n", "r", "t", "v", "f", "s"} {
		rx, _ := regexp.Compile("\\" + pattern)
		s = rx.ReplaceAllString(s, "["+pattern+"]")
	}
	return s
}

func IsWhitespace(char byte) bool {
	return char == '\n' || char == '\r' || char == '\t' || char == ' ' || char == '\v' || char == '\f'
}

func IsWhitespaces(s string) bool {
	rx, _ := regexp.Compile("^[" + WHITESPACE_ALL + "]*$")
	return rx.MatchString(s)
}

func IsEqualWithAClosure(original string, test string, closure_chars string) bool {
	// Trim.
	original_trimmed := strings.Trim(original, WHITESPACE_ALL)

	// Remove closure too.
	rx, _ := regexp.Compile("[" + closure_chars + "]$")
	original_trimmed = rx.ReplaceAllString(original_trimmed, "")

	return (original_trimmed == test) && (rx.MatchString(original))
}
