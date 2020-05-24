package pangram

func IsPangram(str string) bool {
	if len(str) < 24 {
		return false
	}
	pgm := make(map[int]bool)
	for _, r := range str {
		n := toLower(int(r))
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
	return len(pgm) == (123 - 97)
}

func toLower(n int) int {
	if n >= 65 && n <= 90 {
		return n + 32
	}
	return n
}
