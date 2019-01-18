package main

import "fmt"

func ReturnMultiValues(a int, b int) (int, string) {
	c := a + b
	return c, "successfully added"
}
func main() {
	sum, message := ReturnMultiValues(2, 1)
	// fmt.Println(message)
	fmt.Println( message, "result is", sum)
}
