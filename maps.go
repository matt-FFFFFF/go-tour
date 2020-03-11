package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount counts words in a string
func WordCount(s string) map[string]int {

	result := make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, present := result[word]; !present {
			result[word] = 1
		} else {
			result[word]++
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
