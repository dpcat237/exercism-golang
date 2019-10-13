package pythagorean

import (
	"math"
	"sort"
)

type Triplet [3]int

// Improved by: https://exercism.io/tracks/go/exercises/pythagorean-triplet/solutions/b69a478e90ef463c8eb78bc683a8af82
// Range returns all triplets where each side lies in [min..max]
func Range(min, max int) (trips []Triplet) {

	for m := minM(min); m <= maxM(max); m++ {
		// n should start at 1 if m is even or 2 if m is odd and maintain the same even/odd parity
		for n := 1 + m%2; n < m; n += 2 {
			if isCoprime(m, n) {
				a, b, c := euclidPrimitive(m, n)
				for k := minK(min, a); k <= maxK(max, c); k++ {
					trips = append(trips, Triplet{k * a, k * b, k * c})
				}
			}
		}
	}
	return trips
}

func validSum(m, s, p int) bool {
	return p%(m*s) == 0
}

// Sum returns all triplets where x+y+z=p, solution adapted from
// https://stackoverflow.com/a/2835053/4304664
func Sum(p int) []Triplet {

	var trips []Triplet

	if p%2 == 1 {
		return nil
	}
	p /= 2

	limit := sqrt(p)
	factors := getFactors(p)
	for _, m := range factors {
		if m > limit {
			continue
		}

		for _, s := range factors {
			if s >= 2*m || s <= m {
				continue
			}

			if validSum(m, s, p) {
				n := s - m
				k := p / m / s
				trips = append(trips, euclidK(k, m, n))
			}
		}
	}

	return dropDuplicates(trips)
}
func dropDuplicates(triplets []Triplet) (newTrips []Triplet) {
	sort.Slice(triplets, func(i, j int) bool { return triplets[i][0] < triplets[j][0] })
	newTrips = append(newTrips, triplets[0])
	for i := 1; i < len(triplets); i++ {
		if triplets[i] != triplets[i-1] {
			newTrips = append(newTrips, triplets[i])
		}
	}
	return newTrips
}

// euclidK generates a triplet multiple of a primitive given k
func euclidK(k, m, n int) Triplet {
	a, b, c := euclidPrimitive(m, n)
	return Triplet{k * a, k * b, k * c}
}

// genEuclid uses Euclid's formula to generate primitives given m > n and exactly one of m,n are even.
func euclidPrimitive(m, n int) (a, b, c int) {
	a, b, c = m*m-n*n, 2*m*n, m*m+n*n

	if a > b {
		a, b = b, a
	}

	return
}

func getFactors(n int) (factors []int) {

	for i := 1; i <= sqrt(n); i++ {
		if n%i == 0 {
			f1, f2 := i, n/i
			if f1 == f2 {
				factors = append(factors, f1)
			} else {
				factors = append(factors, f1, f2)
			}
		}
	}
	return
}

// GCD uses Dijkstra's algorithm to calculate greatest common divisor
func GCD(m, n int) int {
	if m == n {
		return m
	} else if m > n {
		return GCD(m-n, n)
	}

	return GCD(m, n-m)
}

func isCoprime(m, n int) bool {
	return GCD(m, n) == 1
}

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func minimum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Functions below deal with satisfying the constraints min <= a < b < c <=max when choosing values for k,m,n

// minK returns the largest value of k meeting constraints.
// Given the smallest side a, we need k*a >= min  => k >= min/a
func minK(min, a int) (k int) {
	if min%a == 0 {
		k = min / a
	} else {
		k = (min / a) + 1
	}

	return
}

// maxK returns the largest value of k meeting constraints.
// Given m,n, we need k*c <= max.  Therefore k <= max/c
func maxK(max, c int) (k int) {
	return max / c
}

// minM returns the largest value of m meeting constraints.
// Given a triplet T with a < b < c, we need < a < b < c <= max
// Since c = k(m*m + n*n), the value of m is maximized at k=1 subject to m*m + n*n <= max
// m > n implies  m*m + n*n <= m*m + m*m <= 2m*m, thus satisfying 2*m*m <= max ensures m*m+n*n <= max.
// Therefore we need m <= sqrt(max/2)
func maxM(max int) int {
	return sqrt(max / 2)
}

// minM returns the smallest value of m meeting constraints.
// Given a triplet T with a < b < c, we need min <= a < b < c.
// a = minimum(2*m*n, m*m - n*n), so we must mimimize both equations. Given m > n, we have
// 2*m*m >= 2*m*n >= min   ==>  m >= sqrt(min/2)  , and
//
// The smallest value of n is 1, so we have
// m*m - n*n >= m*m - 1*1 >= min  ==>  m >= sqrt(min-1)
// so minM should be minimum( sqrt(min/2), sqrt(min-1))
func minM(min int) int {
	m1, m2 := sqrt(min/2), sqrt(min-1)
	return minimum(m1, m2)
}
