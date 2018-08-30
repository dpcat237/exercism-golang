package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) < 1 {
		return 0, nil
	}

	if len(a) != len(b) {
		return -1, errors.New("DNAs with different lengths")
	}

	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c, nil
}
