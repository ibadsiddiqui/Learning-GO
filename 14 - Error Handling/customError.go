package main

import (
	"errors"
	"fmt"
)

func Increment(n int) (int, error) {
	if n < 0 {
		// return error object
		return 0, errors.New("math: cannot process negative number")
	} else {
		return (n + 1), nil
	}
}


func main() {
	num := -1
	inc, err := Increment(num)

	if err != nil {
		fmt.Printf("Failed Number: %v, error message: %v", num, err)
	} else {
		fmt.Printf("Incremented Number: %v", inc)
	}
}
