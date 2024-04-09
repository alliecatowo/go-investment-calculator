package main

import (
	"fmt"
	"os"

	"github.com/alliecatowo/go-investment-calculator/investment_calculator"
	"github.com/alliecatowo/go-investment-calculator/profit_calculator"
)

func main() {
	fmt.Println("Welcome to the Investment Calculator!")

	fmt.Println("Please select an option:")
	fmt.Println("1. Profit Calculator")
	fmt.Println("2. Investment Calculator")

	var option int
	// catch invalid input
	_, err := fmt.Scanln(&option)
	// nil, because google had to be quirky!
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		os.Exit(1)
	}

	switch option {
	case 1:
		profit_calculator.CalculateProfit()
	case 2:
		investment_calculator.CalculateInvestment()
	default:
		fmt.Println("Invalid option. Please try again.")
		os.Exit(1)
	}
}
