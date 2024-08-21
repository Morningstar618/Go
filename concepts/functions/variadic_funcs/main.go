package main

import "fmt"

func main() {
	numbers := []int{1, 5, 15}

	sum := sumUp(100, 2, 20, -5, 12)
	anotherSum := sumUp(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)

}

// Variadic functions are functions that work with any amount of parameters
func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}
