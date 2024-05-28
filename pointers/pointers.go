package main

import "fmt"

func main() {
	age := 32
	var agePointer *int = &age
	fmt.Println("Age: ", *agePointer, "Pointer: ", agePointer)
	getAdultYears(agePointer)
	fmt.Println("Adult years: ", age, "Pointer: ", &age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
