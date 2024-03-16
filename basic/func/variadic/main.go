package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum("hello gopher", nums...)
	sum("nothing")
}

func sum(msg string, nums ...int) {
	var result int
	for _, num := range nums {
		result += num
	}

	fmt.Printf("msg: %s, result: %d\n", msg, result)
}
