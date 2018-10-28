package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	for _, candidate := range candidates {
		if len(subject) != len(candidate) {
			continue
		}

		lSub := strings.ToLower(subject)
		lCan := strings.ToLower(candidate)
		if lSub == lCan {
			continue
		}
		
		if equalInts(normalize(lSub), normalize(lCan)) {
			result = append(result, candidate)
		}
	}
	return result
}

func normalize(subject string) []int {
	var nums []int
	for _, r := range subject {
		nums = append(nums, int(r))
	}
	sort.Ints(nums)
	return nums
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
