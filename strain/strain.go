package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns array of validated integers
func (ints Ints) Keep(f func(int) bool) Ints {
	var resl Ints
	for _, i := range ints {
		if f(i) {
			resl = append(resl, i)
		}
	}
	return resl
}

// Keep returns array of discarded integers
func (ints Ints) Discard(f func(int) bool) Ints {
	var resl Ints
	for _, i := range ints {
		if f(i) {
			continue
		}
		resl = append(resl, i)
	}
	return resl
}

// Keep returns array of validated arrays of integers
func (lts Lists) Keep(f func([]int) bool) Lists {
	var resl Lists
	for _, l := range lts {
		if f(l) {
			resl = append(resl, l)
		}
	}
	return resl
}

// Keep returns array of validated strings
func (strs Strings) Keep(f func(string) bool) Strings {
	var resl Strings
	for _, s := range strs {
		if f(s) {
			resl = append(resl, s)
		}
	}
	return resl
}
