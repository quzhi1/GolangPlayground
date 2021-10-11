package iteration

import "strings"

const repeatCount = 5

func Repeat(to_repeat string) string {
	// var result string
	// for i := 0; i < repeatCount; i++ {
	// 	result += to_repeat
	// }
	// return result
	return strings.Repeat(to_repeat, repeatCount)
}
