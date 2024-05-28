package list

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Printf("prices[0:1]: %v\n", prices[0:1])
	// to increase array dynamically, this wont work.
	// instead, should use "append"
	// prices[2] = 11.99
	// fmt.Printf("prices: %v\n", prices)

	// append creates another slice
	updatedPrices := append(prices, 11.99)
	fmt.Println(updatedPrices)
	fmt.Println(len(updatedPrices), cap(updatedPrices))
	fmt.Println(&updatedPrices)
	updatedPrices = append(updatedPrices, 8.99)
	fmt.Printf("updatedPrices: %v\n", updatedPrices)
	fmt.Println(len(updatedPrices), cap(updatedPrices))

	discountPrices := []float64{101.99, 80.99, 20.59}
	// unpack list values
	prices = append(prices, discountPrices...)
	fmt.Printf("prices: %v\n", prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames[2] = "A Carpet"
// 	fmt.Printf("prices: %v\n", prices)
// 	fmt.Printf("productNames: %v\n", productNames)
// 	fmt.Printf("prices[3]: %v\n", prices[3])

// 	// slice: first index included, last excluded
// 	featuredPrices := prices[1:3]
// 	fmt.Printf("featuredPrices: %v\n", featuredPrices)

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Printf("highlightedPrices: %v\n", highlightedPrices)

// 	// slice: all elements except the last:
// 	fmt.Printf("featuredPrices: %v\n", prices[:3])

// 	// slice: all elements except the first:
// 	fmt.Printf("featuredPrices: %v\n", prices[1:])

// 	// a slice does not create another array in memory
// 	featuredPrices[0] = 42.0
// 	fmt.Printf("featuredPrices[0]: %v\n", featuredPrices[0])
// 	fmt.Printf("prices[1]: %v\n", prices[1])

// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	featuredPrices = featuredPrices[:3]
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
