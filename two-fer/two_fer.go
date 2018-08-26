package twofer

import "fmt"

const (
	baseYou = "you"
	baseSentence = "One for %s, one for me."
)

func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf(baseSentence, name)
	}
	return fmt.Sprintf(baseSentence, baseYou)
}
