package scrabble

import "strings"

// Score calculates score of passed string
func Score(w string) int {
	sum := 0
	if len(w) < 1 {
		return sum
	}

	w = strings.ToUpper(w)
	for _, l := range w {
		sum += getMark(l)
	}
	return sum
}

func getMark(ltr rune) int {
	switch ltr {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}
