package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, run(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		oldZ := z
		z -= (z*z - x) / (2*z)
		if oldZ == z {
			return z, nil
		}
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func run(x float64) error {
	return ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
