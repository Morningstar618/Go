package anonfuncs

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformers(2)
	triple := createTransformers(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(transformed)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func createTransformers(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
