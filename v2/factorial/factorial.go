package factorial

import "errors"

// Factorial returns the factorial of n
func Factorial(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Cannot factorize 0")
	}
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f, nil
}
