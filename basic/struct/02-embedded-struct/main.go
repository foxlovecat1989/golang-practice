package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type agent struct {
	firstName string
	license   bool
	person
}

func main() {
	a := agent{
		firstName: "secret",
		license:   true,
		person: person{
			firstName: "ed",
			lastName:  "huang",
			age:       34,
		},
	}
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("lastName: %s\n", a.lastName)
	fmt.Printf("firstName: %s\n", a.firstName)
	fmt.Printf("firstName from person: %s\n", a.person.firstName)
}
