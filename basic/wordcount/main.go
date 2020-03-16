package main

import (
	"strings"

	wc "github.com/golang/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
