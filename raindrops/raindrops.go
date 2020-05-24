package raindrops

import (
	"strconv"
)

var factors = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string {
	var rst string
	for _, k := range [3]int{3, 5, 7} {
		if n%k == 0 {
			rst += factors[k]
		}
	}
	if len(rst) == 0 {
		return strconv.Itoa(n)
	}
	return rst
}
