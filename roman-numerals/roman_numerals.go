package romannumerals

import "errors"

type romanArabic struct {
	roman  string
	arabic int
}

var numbers = []romanArabic{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral convert arabic numbers to roman
func ToRomanNumeral(ar int) (string, error) {
	resl := ""
	if ar < 0 || ar > 3000 {
		return resl, errors.New("Not converted")
	}

	for _, ra := range numbers {
		for ar >= ra.arabic {
			ar -= ra.arabic
			resl += ra.roman
		}
	}

	if len(resl) < 1 {
		return resl, errors.New("Not converted")
	}
	return resl, nil
}
