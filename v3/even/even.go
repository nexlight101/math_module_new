package even

import "errors"

//Even determens if n is even and returns bool
func Even(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("cannot use negative number")
	}
	if n%2 == 0 {
		return true, nil
	}
	return false, nil
}
