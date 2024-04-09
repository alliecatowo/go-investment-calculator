package profit_calculator

import (
	"fmt"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func Prompt() {
	revenue := user_input.GetInput("Enter Your Revenue:")
	expenses := user_input.GetInput("Enter Your Expenses:")
	tax_rate := user_input.GetInput("Enter Your Tax Rate:")

	ebt, profit, ratio := calculate_profits(revenue, expenses, tax_rate)

	fmt.Printf("Your Earnings Before Tax (Gross Profit) is: $%.2f\n", ebt)
	fmt.Printf("Your Earnings After Tax (Net Profit) is: $%.2f\n", profit)
	fmt.Printf("Your Net Profit as a percentage of Revenue is: %.2f%%\n", ratio)

}

func calculate_profits(revenue, expenses, tax_rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = profit / revenue * 100

	return ebt, profit, ratio
}
