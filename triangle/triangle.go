package triangle

import "math"

const (
	NaT = Kind(iota)
	Equ
	Iso
	Sca
)

type Kind int

// KindFromSides calculate a type of the triangle
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if areAllEqual(a, b, c) {
		return Equ
	} else if areTwoEqual(a, b, c) {
		return Iso
	}
	return Sca
}

func isTriangle(a, b, c float64) bool {
	return areSidesPositive(a, b, c) && areTwoSumGreaterThanThird(a, b, c) && !areTwoInfinite(a, b, c)
}

func areAllEqual(a, b, c float64) bool {
	return a == b && b == c
}

func areSidesPositive(a, b, c float64) bool {
	return !(a <= 0 || math.IsNaN(a) || b <= 0 || math.IsNaN(b) || c <= 0 || math.IsNaN(c))
}

func areTwoEqual(a, b, c float64) bool {
	return a == b || a == c || b == c
}

func areTwoInfinite(a, b, c float64) bool {
	return (math.IsInf(a, 1) && math.IsInf(b, 1)) || (math.IsInf(b, 1) && math.IsInf(c, 1)) || (math.IsInf(a, 1) && math.IsInf(c, 1))
}

func areTwoSumGreaterThanThird(a, b, c float64) bool {
	return !(a+b < c || a+c < b || b+c < a)
}
