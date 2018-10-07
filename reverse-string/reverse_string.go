package reverse

import "strings"

func String(str string) string {
	if str == "" {
		return ""
	}

	var wb []string
	wb = strings.Split(str, "")
	for i := len(wb)/2 - 1; i >= 0; i-- {
		opp := len(wb) - 1 - i
		wb[i], wb[opp] = wb[opp], wb[i]
	}

	return strings.Join(wb[:], "")
}
