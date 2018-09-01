package raindrops

import "strconv"

var factors = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(n int) string {
	var str string
	for fac, s := range factors {
		if n%fac == 0 {
			str = str + s
		}
	}

	if len(str) == 0 {
		return strconv.Itoa(n)
	}
	return str
}
