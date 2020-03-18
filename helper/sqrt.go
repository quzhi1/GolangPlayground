package helper

import "fmt"

// Sqrt Calculate square root
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// CalculateSqrt prints sqrt(2)
func CalculateSqrt() {
	fmt.Println(Sqrt(2))
}
