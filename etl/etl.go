package etl

import "strings"

// Transform from legacy system to new
func Transform(given map[int][]string) map[string]int {
	rsl := make(map[string]int)

	for n, lts := range given {
		for _, l := range lts {
			rsl[strings.ToLower(string(l))] = n
		}
	}

	return rsl
}
