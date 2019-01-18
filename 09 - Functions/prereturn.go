package main

import (
	"fmt"
)

// Here result is defined as the return variable. So the variable result 
// defined would be automatically returned without needing to be defined at the return statement at the end.

func addWithPreReturn(value1 int, value2 int) (result int) {
	result = value1 + value2
	return
}

func main() {
	fmt.Println(addWithPreReturn(1, 2))
}
