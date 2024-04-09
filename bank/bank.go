package bank

import (
	"fmt"

	"github.com/alliecatowo/go-investment-calculator/user_input"
)

func Prompt() {
	fmt.Println("Welcome to the GO Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	opt := int(user_input.GetInput("Your choice: "))

	switch opt {
	case 1:
		fmt.Println("Your balance is $1000")
	case 2:
		fmt.Println("You deposited $100")
	case 3:
		fmt.Println("You withdrew $100")
	case 4:
		fmt.Println("Goodbye!")
	}

	fmt.Println("You chose option", opt)
}
