package maps

import "fmt"

func main() {
	// create a map with string keys and string values
	websites := map[string]string{
		"google": "https://google.com",
		"meli":   "https://mercadolivre.com.br",
	}
	fmt.Printf("websites: %v\n", websites)
	fmt.Printf("websites[\"meli\"]: %v\n", websites["meli"])
	// add new item to map - no need to slice
	websites["linkedin"] = "https://linkedin.com"
	fmt.Printf("websites[\"linkedin\"]: %v\n", websites["linkedin"])
	// remove an element
	delete(websites, "google")
	fmt.Printf("websites: %v\n", websites)
}
