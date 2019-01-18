package main

import (
	"fmt"
)

func main() {
	i := 0
	sum := 0
	for i < 10 {
		sum += 1
		fmt.Println(sum)

		i++
	}
	fmt.Println(sum)

}
