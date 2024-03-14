package main

import (
	"fmt"
)

func calcAvg(values []int) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := 0
	for _, mark := range values {
		sum += mark
	}

	avg := float64(sum) / float64(len(values))
	return avg
}

func main() {
	rating := []int{8, 7, 9, 10, 1, 2}
	result := calcAvg(rating)
	fmt.Printf("Средняя оценка: %f\n", result)
}
