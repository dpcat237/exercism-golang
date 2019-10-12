package protein

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidBase = errors.New("Invalid base")
	STOP           = errors.New("STOP")
	transData      = []Transform{
		{[]string{"AUG"}, "Methionine"},
		{[]string{"UUU", "UUC"}, "Phenylalanine"},
		{[]string{"UUA", "UUG"}, "Leucine"},
		{[]string{"UCU", "UCC", "UCA", "UCG"}, "Serine"},
		{[]string{"UAU", "UAC"}, "Tyrosine"},
		{[]string{"UGU", "UGC"}, "Cysteine"},
		{[]string{"UGG"}, "Tryptophan"},
	}
	transStop = Transform{[]string{"UAA", "UAG", "UGA"}, STOP.Error()}
)

type Transform struct {
	Codon   []string
	Protein string
}

// FromRNA transforms from RNA to protein
func FromRNA(in string) ([]string, error) {
	var rsl []string
	reg := regexp.MustCompile("\\S{3}")
	rna := reg.FindAllStringSubmatch(in, -1)

	for i := 0; i < len(rna); i++ {
		c := rna[i][0]
		r, err := FromCodon(c)
		if err == ErrInvalidBase {
			return rsl, err
		}
		if err == STOP {
			return rsl, nil
		}
		rsl = append(rsl, r)
	}
	return rsl, nil
}

// FromCodon transforms from codon to protein
func FromCodon(in string) (string, error) {
	isStop := func(in string) bool {
		for _, s := range transStop.Codon {
			if s == in {
				return true
			}
		}
		return false
	}

	getProtein := func(in string) string {
		for _, td := range transData {
			for _, s := range td.Codon {
				if s == in {
					return td.Protein
				}
			}

		}
		return ""
	}

	if isStop(in) {
		return "", STOP
	}

	p := getProtein(in)
	if p == "" {
		return "", ErrInvalidBase
	}

	return p, nil
}
