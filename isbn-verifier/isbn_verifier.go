package isbn

import (
	"math"
	"strings"
)

func IsValidISBN(txt string) bool {
	txt = clean(txt)
	if len(txt) != 10 {
		return false
	}

	pls := 10
	rst := 0
	for _, num := range txt {
		var n int
		if pls == 1 && num == 88 {
			n = 10
		} else {
			if num < 48 || num > 57 {
				return false
			}
			n = int(num - '0')
		}

		rst += n*pls
		pls--
	}
	return math.Mod(float64(rst), 11) == 0
}

func clean(txt string) string {
	if strings.Contains(txt, "-") {
		return strings.Replace(txt, "-", "", -1)
	}
	return txt
}