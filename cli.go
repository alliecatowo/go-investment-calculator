package main

import (
	"fmt"

	"github.com/alliecatowo/go-investment-calculator/bank"
	"github.com/alliecatowo/go-investment-calculator/investment_calculator"
	"github.com/alliecatowo/go-investment-calculator/profit_calculator"
	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func main() {
	fmt.Println("Welcome to the GO Finance CLI!")

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Bank")
		fmt.Println("2. Profit Calculator")
		fmt.Println("3. Investment Calculator")
		fmt.Println("4. Exit")

		option := user_input.GetInput("Your choice: ")

		switch option {
		case 1:
			bank.Prompt()
		case 2:
			profit_calculator.Prompt()
		case 3:
			investment_calculator.Prompt()
		case 4:
			fmt.Println("Goodbye! Please come again :)")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			continue
		}
	}
}
