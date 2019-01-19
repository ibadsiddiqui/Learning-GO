// Go is built with concurrency in mind. Concurrency in Go can be achieved by Go routines which are lightweight threads.

// Go routine

// Go routines are the function which can run in parallel or concurrently 
// with another function. 
// Creating a Go routine is very simple. 
// Simply by adding a keyword Go in front of a function, 
// we can make it execute in parallel. 
// Go routines are very lightweight, so we can create thousands of them. Let’s look into a simple example:
package main

import (
	"fmt"
	"time"
)

func main() {
	go c()
	fmt.Println("I am main")
	time.Sleep(time.Second * 2)
}
func c() {
	fmt.Println("I am concurrent")
	time.Sleep(time.Second * 2)
}
