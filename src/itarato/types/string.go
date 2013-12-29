package types

import (
	"regexp"
)

func ReplaceWhitespaces(s string) string {
	for _, pattern := range []string{"n", "r", "s", "t"} {
		rx, _ := regexp.Compile("\\" + pattern)
		s = rx.ReplaceAllString(s, "["+pattern+"]")
	}
	return s
}

func IsWhitespace(char byte) bool {
	return true
}
