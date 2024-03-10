package main

import "fmt"

func printBiography(born, school, institute string) {
	biography := fmt.Sprintf("Biography:\n%s\n%s\n%s", born, school, institute)
	fmt.Println(biography)
}

func main() {
	born := "I was born in 1980 in Sevastopol"
	school := "Visited school in Lviv and Kyiv"
	institute := "Got education in Kyiv Polytechnical institute"

	printBiography(born, school, institute)
}
