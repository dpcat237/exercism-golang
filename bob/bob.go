package bob

import (
	"strings"
	"unicode"
)

const (
	questionReply     = "Sure."
	yellReply         = "Whoa, chill out!"
	yellQuestionReply = "Calm down, I know what I'm doing!"
	emptyReply        = "Fine. Be that way!"
	otherReply        = "Whatever."
)

// Reply with predefined sentences when request has specific characters
func Hey(txt string) string {
	txt = strings.TrimSpace(txt)
	if isQuestion(txt) && isYell(txt) {
		return yellQuestionReply
	} else if isQuestion(txt) {
		return questionReply
	} else if isYell(txt) {
		return yellReply
	} else if isEmpty(txt) {
		return emptyReply
	}
	return otherReply
}

func isEmpty(txt string) bool {
	if txt == "" {
		return true
	}

	for _, char := range txt {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func isQuestion(txt string) bool {
	if len(txt) < 1 {
		return false
	}
	return txt[len(txt)-1:] == "?"
}

func isYell(txt string) bool {
	if len(txt) < 1 {
		return false
	}
	hasExla := false
	if txt[len(txt)-1:] == "!" {
		hasExla = true
	}

	cLow := 0
	for _, char := range txt {
		if unicode.IsLetter(char) && !unicode.IsLower(char) {
			cLow++
		} else {
			cLow = 0
		}
		if cLow > 1 && hasExla || cLow > 3 {
			return true
		}
	}
	return false
}
