package accumulate

// Accumulate performs required action of a collection of strings
func Accumulate(cll []string, f func(string) string) []string {
	var r []string
	for _, s := range cll {
		r = append(r, f(s))
	}
	return r
}
