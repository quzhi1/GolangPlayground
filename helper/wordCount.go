package helper

import "strings"

// WordCount returns a map of word => occurence
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word]++
	}
	return result
}
