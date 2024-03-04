package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(4)
	fmt.Printf("n: %d\n", n)
	foo(n)
	foo2(n)
	bar(n)
}

func foo(n int) {
	switch {
	case n > 1:
		fmt.Println("n is greater than 1")
	case n == 1:
		fmt.Println("n is equal 1")
	case n < 1:
		fmt.Println("n is less than 1")
	default:
		fmt.Println("default case - no matched")
	}
}

func foo2(n int) {
	switch n {
	case 0:
		fmt.Println("n is 0")
	case 1:
		fmt.Println("n is 1")
	case 2:
		fmt.Println("n is 2")
	default:
		fmt.Println("default case - no matched")
	}
}

func bar(n int) {
	switch n {
	case 0:
		fmt.Println("Case: 0")
		fallthrough // Case 1 will also execute
	case 1:
		fmt.Println("Case: 1")
	case 2:
		fmt.Println("Case: 2")
	default:
		fmt.Println("default")
	}
}
