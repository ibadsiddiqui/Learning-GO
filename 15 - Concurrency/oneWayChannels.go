// There are cases where we want a Go routine to receive data
// via the channel but not send data, and also vice versa.
// For this, we can also create a one-way channel.
// Let’s look into a simple example:

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	// concurrent function
	go sc(ch)
	fmt.Println(<-ch)
}

func sc(ch chan<- string) {
	ch <- "hello"
}

// In the above example,
// sc is a Go routine which can only send messages to the channel
// but cannot receive messages.
