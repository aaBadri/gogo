package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	map_words := map[string]int{}
	for _, word := range words {
		map_words[word] = map_words[word] + 1
	}
	return map_words
}

func main() {
	wc.Test(WordCount)
}
