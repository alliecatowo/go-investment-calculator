package investment_calculator

import (
	"fmt"
	"math"
)

func CalculateInvestment() {
	// I was suggested to use 2.0% - as if!
	const inflationRate = 3.2
	var investmentAmount, years, expectedReturnRate float64

	fmt.Println("Enter the amount you want to invest: ")
	// Store the results of fmt.Scan into the value at the adress of investmentAmount
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter how many years you want to invest for: ")
	fmt.Scan(&years)

	fmt.Println("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInflatedValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Your expected value after %v years is $%.2f\n", years, futureValue)
	fmt.Printf("Accounting for a %v%% inflation rate, your expected value after %v years is $%.2f\n", inflationRate, years, futureInflatedValue)
}
