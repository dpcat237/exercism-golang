package pythagorean

import (
	"fmt"
	"sort"
	"strings"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	var dto []Triplet
	if max == 0 || min >= max {
		return dto
	}

	var a, b, c int
	m := 2
	for c < max {
		for n := 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if c > max {
				break
			}
			if a <= min || b <= min || c <= min {
				continue
			}
			nCh := []int{a, b, c}
			sort.Ints(nCh)
			dto = append(dto, Triplet{nCh[0], nCh[1], nCh[2]})
		}
		m++
	}
	return dto
}

func Sum(n int) []Triplet {
	var dto []Triplet
	if n == 0 {
		return dto
	}
	added := make(map[string]bool)
	arrayToString := func(arr []int, del string) string {
		return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", del, -1), "[]")
	}
	for i := 1; i <= n/3; i++ {
		for j := 1; j <= n/2; j++ {
			k := n - i - j
			if i*i+j*j != k*k {
				continue
			}
			nCh := []int{i, j, k}
			sort.Ints(nCh)
			key := arrayToString(nCh, ",")
			if _, exists := added[key]; !exists {
				added[key] = true
				dto = append(dto, Triplet{nCh[0], nCh[1], nCh[2]})
			}
		}
	}
	return dto
}
