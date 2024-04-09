package bank

import (
	"fmt"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func Prompt() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to the GO Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	opt := int(user_input.GetInput("Your choice: "))
	fmt.Println("You chose option", opt)

	switch opt {
	case 1:
		fmt.Printf("Your balance is %.2f\n", accountBalance)
	case 2:
		deposit := user_input.GetInput("Enter the amount you want to deposit: ")
		accountBalance += deposit
		fmt.Printf("Your new balance is %.2f\n", accountBalance)
	case 3:
		withdraw := user_input.GetInput("Enter the amount you want to withdraw: ")
		accountBalance -= withdraw
		fmt.Printf("Your new balance is %.2f\n", accountBalance)
	case 4:
		fmt.Println("Goodbye!")
	}

}
