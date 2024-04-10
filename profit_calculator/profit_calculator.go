package profit_calculator

import (
	"fmt"
	"os"
	"strconv"
	"errors"
	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func writeProfit(profit float64) {
	profitBytes := []byte(fmt.Sprintf("%.2f", profit))
	err := os.WriteFile("profit.txt", profitBytes, 0644)
	if err != nil {
		panic("Can't write profit to file! Aborting application!")
	}
}

func getProfit() (float64, error) {
	data, err := os.ReadFile("profit.txt")

	if err != nil {
		writeProfit(100.0)
		return 100.0, errors.New("profit file not found. Default profit of $100.00 set")
	}

	profitText := string(data)
	profit, err := strconv.ParseFloat(profitText, 64)

	if err != nil {
		writeProfit(100.0)
		return 100.0, errors.New("failed to parse profit. Default profit of $100.00 set")
	}
	return profit, nil
}

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
