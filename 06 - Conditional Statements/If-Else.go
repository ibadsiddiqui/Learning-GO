package main

import (
	"fmt"
)

func main() {
	num := 1
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num >= 10 && num < 100 {
		fmt.Println(num, "has 2 digits")
	} else {
		fmt.Println(num, "has 1 digit")
	}
}
