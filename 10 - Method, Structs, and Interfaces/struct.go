package main

import "fmt"

// struct declaration and definition
type person struct {
	name   string
	age    int
	gender string
}

func main() {

	//way 1: specifying attribute and value
	// struct initialization

	p := person{name: "Bob", age: 42, gender: "Male"}
	//way 2: specifying only value
	// person{"Bob", 42, "Male"}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.gender)
	fmt.Println(p.age)
}
