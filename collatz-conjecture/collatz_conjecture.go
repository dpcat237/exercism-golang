package collatzconjecture

import "github.com/pkg/errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Wrong number")
	}
	return calculate(n, 0), nil
}

func calculate(n, steps int) int {
	if n == 1 {
		return steps
	}

	if isEven(n) {
		return calculate(n/2, steps+1)
	}
	return calculate(n*3+1, steps+1)
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
