package main

import (
	"GolangPlayground/helper"
	"fmt"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	printFibonacci()
}

func showPic() {
	pic.Show(helper.Pic)
}

func calculateSqrt() {
	fmt.Println(helper.Sqrt(2))
}

func wordCountTest() {
	wc.Test(helper.WordCount)
}

func printFibonacci() {
	f := helper.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
