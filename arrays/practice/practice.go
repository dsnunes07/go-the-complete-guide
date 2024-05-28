package practice

import "fmt"

type product struct {
	id    int
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Piano", "Running", "Gaming"}
	fmt.Printf("hobbies: %v\n", hobbies)
	// 2
	fmt.Printf("hobbies[0]: %v\n", hobbies[0])
	fmt.Printf("hobbies[1:]: %v\n", hobbies[1:])
	// 3
	mainHobbies := hobbies[:2]
	fmt.Printf("mainHobbies: %v cap: %d\n", mainHobbies, cap(mainHobbies))
	// 4
	mainHobbies = mainHobbies[1:3]
	fmt.Printf("mainHobbies: %v\n", mainHobbies)
	// 5
	var courseGoals = []string{"Learn go", "Learn all the basics"}
	fmt.Printf("goals: %v\n", courseGoals)

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Printf("courseGoals: %v\n", courseGoals)

	// in an array of structs, dont need to specify the struct name in each element!
	var products = []product{
		{
			1,
			"A book",
			30.99,
		},
		{
			2,
			"A carpet",
			12.00,
		},
	}
	fmt.Printf("products: %v\n", products)
	newProduct := product{
		3,
		"A mug",
		4.99,
	}
	products = append(products, newProduct)
	fmt.Printf("products: %v\n", products)
}
