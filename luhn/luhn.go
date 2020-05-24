package luhn

import "strconv"

func Valid(txt string) bool {
	total := len(txt)
	if total < 2 {
		return false
	}

	c := 1
	cn := 0
	sum := 0
	for i := total - 1; i >= 0; i-- {
		r := txt[i]
		if !isValidRune(rune(r)) {
			return false
		}
		if r == 32 {
			continue
		}

		if c == 2 {
			sum += doublingNumber(runeToNum(r))
			c = 1
			cn++
			continue
		}
		sum += runeToNum(r)
		cn++
		c++
	}
	if cn < 2 {
		return false
	}
	return sum%10 == 0
}

func doublingNumber(num int) int {
	rst := num + num
	if rst < 10 {
		return rst
	}
	rstStr := strconv.Itoa(num + num)
	return runeToNum(rstStr[0]) + runeToNum(rstStr[1])
}

func isValidRune(r rune) bool {
	return r == 32 || (r >= 48 && r <= 57)
}

func runeToNum(r uint8) int {
	return int(r) - 48
}
