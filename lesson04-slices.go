package main

import (
	"fmt"
)

func main() {

	var mySlice []int

	for i := 0; i < 40; i++ {
		mySlice = append(mySlice, i)

		fmt.Printf("Element: %v (address: %p, size: %d, capacity: %d\n", mySlice[i], &mySlice[i], len(mySlice), cap(mySlice))
	}
}
