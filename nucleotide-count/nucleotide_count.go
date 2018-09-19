package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts count nucleotide in a string
func (d DNA) Counts() (Histogram, error) {
	h := generateHistogram()
	for _, l := range d {
		if !isHistogram(l) {
			return h, errors.New("Not a histogram")
		}
		h[l]++
	}
	return h, nil
}

func generateHistogram() Histogram {
	return Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
}

func isHistogram(l rune) bool {
	return l == 'A' || l == 'C' || l == 'G' || l == 'T'
}
