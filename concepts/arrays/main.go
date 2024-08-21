package main

import "fmt"

// DYNAMIC ARRAY
func main() {
	prices := []float64{1.99, 5.99, 3.50}

	prices = append(prices, 4.66) // returns an array
	fmt.Println(prices)

	discountedPrices := []float64{1.50, 5, 50, 2.90}
	prices = append(prices, discountedPrices...)

	fmt.Println(prices)

}

// STATIC ARRAY
// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 43.99, 90.50, 21.25}

// 	productNames[2] = "A Carpet"

// 	featuredPrices := prices[1:3]

// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)
// 	fmt.Println(featuredPrices)

// 	fmt.Println("Len:", len(featuredPrices))      // Number of elements in an array
// 	fmt.Println("Capacity:", cap(featuredPrices)) // Number of elements available in an array from the starting index to the left-end of the array.
// }
