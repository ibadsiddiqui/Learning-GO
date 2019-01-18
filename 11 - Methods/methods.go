package main

import "fmt"

// struct defination
type Person struct {
	name   string
	age    int
	gender string
}

// method definition of Stuct Person
func (p Person) describe() {
	fmt.Printf("%v is %v years old.", p.name, p.age)
}
func (p *Person) setAge(age int) {
	p.age = age
}

func (p *Person) setName(name string) {
	p.name = name
}

func main() {
	pp := Person{name: "Bob", age: 42, gender: "Male"}
	pp.describe()
	// => Bob is 42 years old
	pp.setAge(45)
	fmt.Printf("\n%v", pp.age)
	//=> 45
	pp.setName("Hari")
	fmt.Printf("\n%v\n", pp.name)
	//=> Bob

	pp.describe()

}
