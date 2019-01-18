// Go provides pointers. Pointers are the place to hold the address of a value.
// A pointer is defined by *. A pointer is defined according to the type of data. Example:

package main

import (
	"fmt"
)

func main() {
	var app *int

	a := 12
	app = &a

	// prints the address of variable "a"
	fmt.Println(app)  
	
	// prints value to stored at '0xc0000120b0" -> it is the location where variable "a" is stored
	fmt.Println(*app) 
}
