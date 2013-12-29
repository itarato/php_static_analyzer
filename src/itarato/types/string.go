package types

import (
	"regexp"
)

func ReplaceWhitespaces(s string) string {
	// Replace whitespaces to something visual. Eg.: \n -> [n].
	for _, pattern := range []string{"n", "r", "t", "s"} {
		rx, _ := regexp.Compile("\\" + pattern)
		s = rx.ReplaceAllString(s, "["+pattern+"]")
	}
	return s
}

func IsWhitespace(char byte) bool {
	return char == '\n' || char == '\r' || char == '\t' || char == ' '
}
