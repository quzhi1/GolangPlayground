package helper

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of word => occurence
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word]++
	}
	return result
}

// WordCountTest verify WordCount function
func WordCountTest() {
	wc.Test(WordCount)
}
