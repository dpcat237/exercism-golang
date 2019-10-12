package pangram

import "strings"

func IsPangram(str string) bool {
	if len(str) < 24 {
		return false
	}
	pgm := make(map[int]bool)
	for _, r := range strings.ToLower(str) {
		n := int(r)
		if isValid(n) {
			pgm[n] = true
		}
	}
	return hasAll(pgm)
}

func isValid(n int) bool {
	return n >= 97 && n <= 122
}

func hasAll(pgm map[int]bool) bool {
	for i := 97; i <= 122; i++ {
		if !pgm[i] {
			return false
		}
	}
	return true
}
