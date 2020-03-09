package helper

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
