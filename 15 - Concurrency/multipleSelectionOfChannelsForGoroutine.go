// Organizing multiple channels for a Go routine using select

// There may be multiple channels that a function is waiting on.
// For this, we can use a select statement.
// Let us take a look at an example for more clarity:
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// concurrent processing is done here
	go speed1(c1)
	go speed2(c2)
	fmt.Println("The first to arrive is:")
	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	}
}

// In the above example, the main is waiting on two channels, c1 and c2.
// With select case statement the main function prints,
// the message sends from the channel whichever it receives first.

func speed1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "speed 2"
}
