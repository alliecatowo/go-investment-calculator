package bank

import (
	"fmt"
	"os"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func writeBalance(balance float64) {
	balanceBytes := []byte(fmt.Sprintf("%.2f", balance))
	os.WriteFile("balance.txt", balanceBytes, 0644)
}
func Prompt() {
	fmt.Println("Welcome to the GO Bank!")
	var accountBalance float64 = 1000
loop: //label the loop to allow break to exit from switch statement
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		opt := int(user_input.GetInput("Your choice: "))

		switch opt {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			deposit := user_input.GetInput("Enter the amount you want to deposit: ")
			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0. Please try again.")
				continue
			}
			accountBalance += deposit
			writeBalance(accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		case 3:
			withdraw := user_input.GetInput("Enter the amount you want to withdraw: ")
			if accountBalance < withdraw {
				fmt.Println("Insufficient funds.")
				continue
			}
			if withdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0. Please try again.")
				continue
			}
			accountBalance -= withdraw
			writeBalance(accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		case 4:
			break loop
		default:
			fmt.Println("Invalid option. Please try again.")
			continue
		}
	}
	fmt.Println("Goodbye!")

}
