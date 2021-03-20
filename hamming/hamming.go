package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	differenceCount := 0

	if len(a) != len(b) {
		return differenceCount, errors.New("Irregular DNA structure (length)")
	}

	for i, v := range a {
		if v != rune(b[i]) {
			differenceCount++
		}
	}

	return differenceCount, nil
}
