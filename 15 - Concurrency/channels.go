// We can pass data between two Go routines using channels.
// While creating a channel it is necessary to specify what
// kind of data the channel receives.
// Let’s create a simple channel with string type as follows:
// With this channel,
// we can send string type data.
// We can both send and receive data in this channel:
package main

import "fmt"

func main() {

	// creates a data channel of string and stores it in variable C
	c := make(chan string)
	go func() { c <- "ibad" }()
	msg := <-c
	fmt.Println(printC(msg))
}

func printC(c string) string {
	return c
}

// Go prefers not sharing the variables of one thread with another
// because this adds a chance of deadlock and resource waiting.
