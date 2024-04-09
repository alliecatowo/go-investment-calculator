package investment_calculator

import (
	"fmt"
	"math"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func Prompt() {
	// I was suggested to use 2.0% - as if!
	const inflationRate = 2.0
	var years, expectedReturnRate float64

	investmentAmount := user_input.GetInput("Enter the amount you want to invest: ")
	years = user_input.GetInput("Enter how many years you want to invest for: ")
	expectedReturnRate = user_input.GetInput("Enter the expected return rate: ")

	futureValue, futureInflatedValue := future_investments(investmentAmount, years, expectedReturnRate, inflationRate)

	fmt.Printf("Your expected value after %v years is $%.2f\n", years, futureValue)
	fmt.Printf("Accounting for a %v%% inflation rate, your expected value after %v years is $%.2f\n", inflationRate, years, futureInflatedValue)
}

func future_investments(investmentAmount, years, expectedReturnRate, inflationRate float64) (futureValue, futureInflatedValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInflatedValue = futureValue / math.Pow(1+inflationRate/100, years)

	// you could use 'return' but I prefer to be explicit for claritys sake
	return futureValue, futureInflatedValue
}
