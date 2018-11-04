package wordcount

import (
	"log"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(txt string) Frequency {
	rst := make(Frequency)
	if len(txt) < 1 {
		return rst
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	regAp, err := regexp.Compile("[^a-zA-Z0-9'-]+")
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(strings.ToLower(strings.Replace(txt, ",", " ", -1)))
	for _, str := range words {
		rst[cleanWord(reg, regAp, str)]++
	}

	return rst
}

func cleanWord(reg, regAp *regexp.Regexp, str string) string {
	if strings.Count(str, "'") > 1 {
		return reg.ReplaceAllString(str, "")
	}
	return regAp.ReplaceAllString(str, "")
}
