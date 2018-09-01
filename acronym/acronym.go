package acronym

import "strings"

// Abbreviate
func Abbreviate(s string) string {
	var rsl string
	take := true

	for _, char := range s {
		if take && string(char) != "" {
			rsl += strings.ToUpper(string(char))
			take = false
			continue
		}

		if string(char) == " " || string(char) == "-" {
			take = true
		}
	}
	return rsl
}
