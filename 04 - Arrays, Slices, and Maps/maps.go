package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["route"] = 66
	m["clearity"] = 2
	m["simplicity"] = 3
	// delete(m, "route")
	i, ok := m["route"]
	i, ok := m["clearity"]
	i, ok := m["simplicity"]

	fmt.Println(m)
	fmt.Println(i)

	fmt.Println(ok)
}
