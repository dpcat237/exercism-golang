package lsproduct

import "errors"

type Product struct {
	result uint64
	digits []uint64
}

func LargestSeriesProduct(dgtsStr string, sp int) (int, error) {
	if sp == 0 {
		return 1, nil
	}
	if sp < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	if sp > len(dgtsStr) {
		return -1, errors.New("span larger than digits length")
	}

	circBuff := make([]int, sp)
	product := 1
	var cInd, max int
	for _, dgtStr := range dgtsStr {
		if dgtStr < '0' || dgtStr > '9' {
			return 0, errors.New("strings in input")
		}
		val := int(dgtStr - '0')

		if val == 0 {
			cInd = 0
			product = 1
			continue
		}

		product *= val
		circBuff[cInd%sp] = val
		cInd++

		if cInd < sp {
			continue
		}

		if product > max {
			max = product
		}
		product = product / circBuff[cInd%sp]
	}

	return max, nil
}
