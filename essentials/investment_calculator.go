package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment amount: ")
	// Scan receives a pointer!
	fmt.Scan(&investmentAmount)
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years to invest: ")
	fmt.Scan(&years)

	futureValue, adjustedFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	output := fmt.Sprintf(`
	Future value: %.2f
	Future value adjusted by inflation: %.2f`,
		futureValue,
		adjustedFutureValue)
	fmt.Print(output)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, adjustedFutureValue float64) {
	const inflationRate float64 = 2.5
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	adjustedFutureValue = futureValue / math.Pow((1+inflationRate/100), years)
	return futureValue, adjustedFutureValue
	// or just "return", since the variables are defined in the function signature
}
