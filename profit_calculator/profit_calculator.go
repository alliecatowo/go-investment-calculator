package profit_calculator

import "fmt"

func CalculateProfit() {
	var revenue, expenses, tax_rate float64
	fmt.Println("Enter Your Revenue:")
	fmt.Scan(&revenue)

	fmt.Println("Enter Your Expenses:")
	fmt.Scan(&expenses)

	fmt.Println("Enter Your Tax Rate:")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := profit / revenue * 100

	fmt.Printf("Your Earnings Before Tax (Gross Profit) is: $%.2f\n", ebt)
	fmt.Printf("Your Earnings After Tax (Net Profit) is: $%.2f\n", profit)
	fmt.Printf("Your Net Profit as a percentage of Revenue is: %.2f%%\n", ratio)

}
