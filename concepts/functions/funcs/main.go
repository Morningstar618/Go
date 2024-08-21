package funcs

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformValue(&numbers, double)
	tripled := transformValue(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformValue(&numbers, transformFn1)
	moreTransformedNumbers := transformValue(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformValue(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func getTransformerFunction(number *[]int) transformFn {
	if (*number)[0] == 1 {
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
