package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 5, 10)
	numbers = append(numbers, 1, 2, 3, 4)
	fmt.Println(numbers)
}
