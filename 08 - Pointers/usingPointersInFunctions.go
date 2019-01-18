// While passing value, the value is actually copied which means more memory
// With the pointer passed, 
// The value is changed by the function is reflected back in the method/function caller

package main

import (
	"fmt"
)

func increment(pointer *int)  {
	*pointer++
}


func main(){ 
	value := 10

	// when using pointers in function param
	// u need to pass the address of the value that needs to be changed
	increment(&value)
	fmt.Println(value)
}