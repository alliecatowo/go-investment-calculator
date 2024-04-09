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

	fmt.Println("Your expected value after", years, "years is $", futureValue)
	fmt.Println("Accounting for a", inflationRate, "% inflation rate, your expected value after", years, "years is $", futureInflatedValue)
}
