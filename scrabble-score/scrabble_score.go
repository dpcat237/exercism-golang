package scrabble

import "strings"

var marks = []Mark{
	{Letters: "AEIOULNRST", Mark: 1},
	{Letters: "DG", Mark: 2},
	{Letters: "BCMP", Mark: 3},
	{Letters: "FHVWY", Mark: 4},
	{Letters: "K", Mark: 5},
	{Letters: "JX", Mark: 8},
	{Letters: "QZ", Mark: 10},
	{Letters: "AEIOULNRST", Mark: 1},
}

// Mark is a structure for letters mark
type Mark struct {
	Letters string
	Mark    int
}

// Score calculates score of passed string
func Score(w string) int {
	sum := 0
	if len(w) < 1 {
		return sum
	}

	w = strings.ToUpper(w)
	ltrs := make(map[string]int)
	for _, l := range w {
		ltrs[string(l)]++
	}
	for l, c := range ltrs {
		sum += c * getMark(l)
	}
	return sum
}

func getMark(ltr string) int {
	for _, m := range marks {
		for _, l := range m.Letters {
			if ltr == string(l) {
				return m.Mark
			}
		}
	}
	return 0
}
