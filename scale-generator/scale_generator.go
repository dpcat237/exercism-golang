package scale

import "strings"

const (
	typeF = "flat"
	typeS = "sharp"
)

var (
	filterF = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	scales  = map[string][]string{
		typeF: {"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
		typeS: {"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
	}
)

func Scale(t, int string) []string {
	if int == "" {
		return getFullScale(t)
	}

	var strs []string
	scl := getFullScale(t)
	salt := 0
	strs = append(strs, scl[0])
	for i := 0; i < len(int)-1; i++ {
		switch int[i] {
		case 'M':
			salt += 2
			strs = append(strs, scl[salt])
		case 'm':
			salt++
			strs = append(strs, scl[salt])
		case 'A':
			salt += 3
			strs = append(strs, scl[salt])
		}
	}
	return strs
}

func getPosition(t string, values []string) int {
	t = getTonic(t)
	for p, v := range values {
		if v == t {
			return p
		}
	}
	return -1
}

func getPositionType(t string) (int, string) {
	if useFlats(t) {
		return getPosition(t, scales[typeF]), typeF
	}
	return getPosition(t, scales[typeS]), typeS
}

func getFullScale(t string) []string {
	var strs []string
	p, tp := getPositionType(t)
	scales := scales[tp]
	strs = append(strs, scales[p:]...)
	strs = append(strs, scales[:p]...)
	return strs
}

func getTonic(t string) string {
	to := strings.ToUpper(string(t[0]))
	if len(t) == 2 {
		to = to + string(t[1])
	}
	return to
}

func useFlats(t string) bool {
	for _, char := range filterF {
		if char == t {
			return true
		}
	}
	return false
}
