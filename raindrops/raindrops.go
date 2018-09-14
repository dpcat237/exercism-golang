package raindrops

import (
	"sort"
	"strconv"
)

var factors = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string {
	var str string

	var keys []int
	for k := range factors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, fac := range keys {
		if n%fac == 0 {
			str = str + factors[fac]
		}
	}

	if len(str) == 0 {
		return strconv.Itoa(n)
	}
	return str
}
