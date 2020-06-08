package sum

// Sum sums the given numbers and returns the sum
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}

// AddF sums float numbers
func AddF(n ...float64) float64 {
	s := 0.
	for _, v := range n {
		s += v
	}
	return s
}
