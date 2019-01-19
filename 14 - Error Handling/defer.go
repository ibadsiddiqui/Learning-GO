// A defer statement pushes a function call onto a list.
// The list of saved calls is executed after the surrounding function returns.
// Defer is commonly used to simplify functions that perform various clean-up actions.

// The function g takes the int i, and panics if i is greater than 3,
// or else it calls itself with the argument i+1.
// The function f defers a function that calls recover and prints the recovered
// value (if it is non-nil). Try to picture what the output of this program might be before reading on.

package main

import "fmt"

func main() {
	// calls function f
	f()

	// prints after f() is executed
	fmt.Println("Returned normally from f.")
}

func f() {

	// creates a nested defer function which is called during panic
	
	defer func() {
		// it captures the value at which panicking started
		// and returns to main()
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// prints as during normal execution
	fmt.Println("Calling g.")

	// calls function g
	
	g(0)
	// g(5)
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
