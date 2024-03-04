package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
		age       int
	}

	p1 := person{
		firstName: "ed",
		lastName:  "huang",
		age:       34,
	}

	p2 := person{
		firstName: "mary",
		lastName:  "lin",
		age:       21,
	}

	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p2: %#v\n", p2)

	fmt.Printf("p1 type: %T\n", p1)
}
