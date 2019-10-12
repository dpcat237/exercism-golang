package summultiples

func SumMultiples(n int, divisors ...int) int {
	if n == 0 || len(divisors) == 0 {
		return 0
	}
	isMult := func(num int, vls []int) bool {
		for _, vl := range vls {
			if vl==0 {
				continue
			}
			if num%vl == 0 {
				return true
			}
		}
		return false
	}

	var sum int
	for i := 1; i < n; i++ {
		if isMult(i, divisors) {
			sum+=i
		}
	}
	return sum
}
