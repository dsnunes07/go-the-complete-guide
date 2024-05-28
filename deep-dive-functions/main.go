package main

import (
	"fmt"

	"example.com/functions/anonymous_functions"
	"example.com/functions/recursion"
	"example.com/functions/variadic_functions"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	transformFunction := getTransformFunction(&numbers)
	doubledNumbers := transformNumbers(&numbers, transformFunction)
	fmt.Printf("doubledNumbers: %v\n", doubledNumbers)
	tripledNumbers := transformNumbers(&numbers, transformFunction)
	fmt.Printf("tripledNumbers: %v\n", tripledNumbers)
	anonymous_functions.AnonymousDemo()
	recursion.Recursion()
	variadic_functions.Variadic()
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

// functions can be returned as well
func getTransformFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
