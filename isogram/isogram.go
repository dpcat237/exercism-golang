package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram check if string is isogram
func IsIsogram(str string) bool {
	for i, char := range strings.ToLower(str) {
		if !unicode.IsLetter(char) {
			continue
		}

		if lastI := strings.LastIndexAny(str, string(char)); lastI > i {
			return false
		}
	}
	return true
}
