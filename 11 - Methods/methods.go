// Methods are a special type of function with a receiver.
// A receiver can be both a value or a pointer.
// Let’s create a method called describe which has a receiver type person we created in the above example:

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

// As we can see in the above example,
// the method now can be called using a dot operator as pp.describe.
// Note that the receiver is a pointer.
// With the pointer we are passing a reference to the value,
// so if we make any changes in the method it will be reflected in the receiver pp.
// It also does not create a new copy of the object, which saves memory.

// Note that in the above example the value of age is changed,
// whereas the value of name is not changed because the method setName
// is of the receiver type whereas setAge is of type pointer.