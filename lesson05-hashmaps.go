package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "это строка строка для теста теста"

	words := strings.Split(str, " ")

	wordCount := make(map[string]int)

	for _, word := range words {
		word = strings.TrimSpace(strings.ToLower(word))
		if word != "" {
			wordCount[word]++
		}
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
