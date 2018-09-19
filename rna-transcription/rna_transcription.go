package strand

// ToRNA transform DNA to RNA
func ToRNA(dna string) string {
	var rna string
	for _, l := range dna {
		rna += string(transform(l))
	}
	return rna
}

func transform(l rune) rune {
	switch l {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	}
	return 0
}
