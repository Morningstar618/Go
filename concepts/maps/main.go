package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	// Using make function to reserve memory for data structures like arrays in advance.
	//-
	// This makes appending data to that array more efficient as go won't have to assign new memory
	// prior to appending any data to the array by making a new array in that memory and copying
	// data to that new array with the appended data.
	//-
	// Additionally we can initialize an empty slice (slice and array are different) and assign the values
	// through indexes later on as creating array using the make function will reservce the space for the
	// number of values we tell it to.
	// --
	// Make function syntax - make(type of array, length to reserve, capacity)

	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Ayush")
	userNames = append(userNames, "Joshi")

	userNames[0] = "Names"

	fmt.Println(userNames)

	// The same functionality can be achieved in Maps too by doing the following.
	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7 // Here go will have to allot new space in the memory as the reserved space
	// was for 2 strings when creating the map using the make command.

	courseRatings.output()

	// For loops - Array
	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	fmt.Println("------------------- ")

	// For loops - Maps || Index = Key
	for index, value := range courseRatings {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

}
