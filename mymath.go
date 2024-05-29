package main

import (
	"fmt"
)

func Sqrt(x float64) (z float64) {
	z = 1.
	prior := 0.
	i := 1
	for !(abs(z-prior) < 0.0000000001) {
		prior = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: %f\n", i, z)
		i += 1
	}
	return
}

func abs(z float64) float64 {
	if z < 0 {
		return -z
	} else {
		return z
	}
}
