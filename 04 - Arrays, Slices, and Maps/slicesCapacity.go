package main

import (
	"fmt"
)

func main() {
	number2 := make([]int, 15)
	// copy the original slice to new slice
	copy(number2, number)
	fmt.Println(numbers)
}
