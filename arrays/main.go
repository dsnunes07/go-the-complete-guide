package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Printf("m: %v\n", m)
}

func main() {
	// tells go to create a slice with len 2 and cap 5 from the start, to avoid array recreation when appending
	userNames := make([]string, 2, 5)
	// then we can assign elements to the pre created positions
	userNames[0] = "Dan"
	userNames[1] = "May"
	// after len is reached, need to append
	userNames = append(userNames, "Lalo")
	fmt.Printf("userNames: %v\n", userNames)

	// make also works for maps
	// type alias
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5
	courseRatings.output()

	// for loop to iterate through values with indexed
	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}

}
