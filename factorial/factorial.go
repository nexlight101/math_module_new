package factorial

import "errors"

// Factorial returns the slice of intergers
func Factorial(n int) ([]int, error) {
	if n < 1 {
		return nil, errors.New("Cannot factorize 0")
	}
	f := 1
	fx := []int{1}
	for i := 2; i <= n; i++ {
		f *= i
		fx = append(fx, f)
	}
	return fx, nil
}
