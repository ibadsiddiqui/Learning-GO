package main

import (
	"fmt"
)

func main() {
	// create a new slice
	number2 := []int{1,2,3,4,5}
	// copy the original slice to new slice
	// number2 = append(number2, 1, 2, 3, 4)
	number := make([]int, len(number2))
	copy(number, number2)
	fmt.Println(number2)
	fmt.Println(number)
}
