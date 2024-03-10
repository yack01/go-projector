package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		randomNumber := rand.Intn(100)
		fmt.Printf("Random digit: %d, i= %d Остаток от деления на i: %d, на 3: %d, на 9: %d\n",
			randomNumber, i, randomNumber%2, randomNumber%3, randomNumber%9)
	}
}
