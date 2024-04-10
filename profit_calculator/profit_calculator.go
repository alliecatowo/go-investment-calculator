package profit_calculator

import (
	"fmt"
	"os"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func writeProfit(profit float64) {
	profitBytes := []byte(fmt.Sprintf("%.2f", profit))
	err := os.WriteFile("profit.txt", profitBytes, 0644)
	if err != nil {
		panic("Can't write profit to file! Aborting application!")
	}
}

func Prompt() {
	for {
		revenue := user_input.GetInput("Enter Your Revenue:")
		expenses := user_input.GetInput("Enter Your Expenses:")
		tax_rate := user_input.GetInput("Enter Your Tax Rate:")

		ebt, profit, ratio := calculate_profits(revenue, expenses, tax_rate)

		fmt.Printf("Your Earnings Before Tax (Gross Profit) is: $%.2f\n", ebt)
		fmt.Printf("Your Earnings After Tax (Net Profit) is: $%.2f\n", profit)
		fmt.Printf("Your Net Profit as a percentage of Revenue is: %.2f%%\n", ratio)

		choice := user_input.GetInput("Do you want to calculate again? (1 for Yes, 0 for No):")

		switch choice {
		case 1:
			continue
		case 0:
			return
		default:
			fmt.Println("Invalid option. Exiting to main menu. Please launch the Profit Calculator again.")
			return
		}
	}

}

func calculate_profits(revenue, expenses, tax_rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = profit / revenue * 100
	writeProfit(profit)
	return ebt, profit, ratio
}
