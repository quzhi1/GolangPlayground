package main

import (
	"GolangPlayground/helper"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
)

func main() {
	displayRot13()
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

func printExampleIPAddr() {
	hosts := map[string]helper.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func checkSqrtWithError() {
	fmt.Println(helper.SqrtWithError(2))
	fmt.Println(helper.SqrtWithError(-2))
}

func verifyMyReader() {
	reader.Validate(helper.MyReader{})
}

func displayRot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := helper.Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)
}
