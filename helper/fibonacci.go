package helper

import "fmt"

// Fibonacci returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
func Fibonacci() func() int {
	fa := 0
	fb := 1
	return func() int {
		result := fa
		fa = fb
		fb = result + fb
		return result
	}
}

// PrintFibonacci print Fibonacci array of 10
func PrintFibonacci() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
