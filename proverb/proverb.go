package proverb

import "fmt"

const (
	fraseTemplate = "For want of a %s the %s was lost."
	endTemplate   = "And all for the want of a %s."
)

// Proverb generates relevant proverbs
func Proverb(rhyme []string) []string {
	var generated []string
	total := len(rhyme)
	if total < 1 {
		return generated
	}

	for i := 0; i < total; i++ {
		if i < (total - 1) {
			generated = append(generated, fmt.Sprintf(fraseTemplate, rhyme[i], rhyme[i+1]))
			continue
		}
		generated = append(generated, fmt.Sprintf(endTemplate, rhyme[0]))
	}

	return generated
}
