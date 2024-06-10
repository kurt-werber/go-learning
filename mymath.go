package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(*e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, &err
	}
	z := 1.
	prior := 0.
	i := 1
	for !(abs(z-prior) < 0.0000000001) {
		prior = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: %f\n", i, z)
		i += 1
	}
	return z, nil
}

func abs(z float64) float64 {
	if z < 0 {
		return -z
	} else {
		return z
	}
}
