package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResults(ebt, profit, ratio float64) {
	var resultsText = fmt.Sprintf(`
	EBT: %.2f
	EAT: %.2f
	Ratio: %.2f
	`, ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}

func main() {
	revenue, revenueErr := promptUserInput("Revenue: ")
	expense, expensesErr := promptUserInput("Expenses: ")
	taxRate, taxRateErr := promptUserInput("Tax rate: ")
	if revenueErr != nil || expensesErr != nil || taxRateErr != nil {
		panic(revenueErr)
	}

	earningsBeforeTax := calculateEarningsBeforeTax(revenue, expense)
	earningsAfterTax := calculateEarningsAfterTax(earningsBeforeTax, taxRate)
	ratio := calculateRatio(earningsBeforeTax, earningsAfterTax)

	fmt.Println("EBT:", formatValue(earningsBeforeTax))
	fmt.Println("EAT:", formatValue(earningsAfterTax))
	fmt.Println("Ratio: ", formatValue(ratio))
	storeResults(earningsAfterTax, earningsAfterTax, ratio)
}

func promptUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid input. only positive values allowed")
	}

	return userInput, nil
}

func calculateEarningsBeforeTax(revenue, expense float64) float64 {
	return revenue - expense
}

func calculateEarningsAfterTax(earningsBeforeTax, taxRate float64) float64 {
	return earningsBeforeTax * (1 - taxRate/100)
}

func calculateRatio(valueBefore, valueAfter float64) float64 {
	return valueBefore / valueAfter
}

func formatValue(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
