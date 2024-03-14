package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := []string{
		"here is a string string for test test",
		"друга строка для перевірки перевірки",
		"слова слова для анализа анализа",
		"еще одна для тестирования тестирования",
		"последняя строка для примера строка примера",
	}

	index := make(map[string][]string)

	for _, line := range lines {

		words := strings.Split(line, " ")

		for _, word := range words {

			word = strings.TrimSpace(strings.ToLower(word))
			if word != "" &&
				(len(index[word]) == 0 || index[word][len(index[word])-1] != line) {
				index[word] = append(index[word], line)
			}
		}
	}

	searchWord := "строка"
	fmt.Printf("Strings containing searchword: '%s':\n", searchWord)
	for _, line := range index[searchWord] {
		fmt.Println(line)
	}
}
