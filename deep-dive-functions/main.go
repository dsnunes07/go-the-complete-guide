package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Printf("doubledNumbers: %v\n", doubledNumbers)
	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Printf("tripledNumbers: %v\n", tripledNumbers)
}

// in go, we can pass functions as parameter values
// a function can be value typed as well (with transformFn instead of full function signature)
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
