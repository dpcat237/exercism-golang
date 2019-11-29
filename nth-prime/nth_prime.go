package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	i := 2
	prm := 0
	for prm < n {
		if isPrime(i) {
			prm++
			if prm == n {
				return i, true
			}
		}
		i++
	}
	return 0, false
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
