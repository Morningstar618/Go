package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investedAmount float64
	var years float64
	var expectedReturn float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investedAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturn)

	// futureValue := investedAmount * math.Pow(1+expectedReturn/100, years)
	// realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	var fv, rfv = calculateFutureValues(investedAmount, expectedReturn, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", fv)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Infation): %.2f\n", rfv)

	fmt.Print(formattedFV, formattedRFV)

}

func calculateFutureValues(investedAmount, expectedReturn, years float64) (float64, float64) {
	fv := investedAmount * math.Pow(1+expectedReturn/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
