package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return -1, errors.New("DNAs with different lengths")
	}

	var dist int
	for _, r1 := range a {
		r2, size := utf8.DecodeRuneInString(b)
		if r1 != r2 {
			dist++
		}
		b = b[size:]
	}

	return dist, nil
}
