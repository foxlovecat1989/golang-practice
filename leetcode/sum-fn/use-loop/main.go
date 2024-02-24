package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4}
	results1 := calculate(nums1)
	fmt.Println(results1) // [1 3 6 10]

	nums2 := []int{1, 1, 1, 1}
	results2 := calculate(nums2)
	fmt.Println(results2) // [1 2 3 4]
}

func calculate(nums []int) []int {
	var results []int
	for i := range nums {
		n := addAll(nums[:i+1])
		results = append(results, n)
	}

	return results
}

func addAll(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	return sum
}
