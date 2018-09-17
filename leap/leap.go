package leap

// IsLeapYear calculates if is a leap year
func IsLeapYear(y int) bool {
	divisibleBy := func(d, by int) bool {
		if d%by == 0 {
			return true
		}
		return false
	}

	switch {
	case divisibleBy(y, 4) && !divisibleBy(y, 100):
		return true
	case divisibleBy(y, 400):
		return true
	}
	return false
}
