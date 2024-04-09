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

	option := user_input.GetInput("Your choice: ")

	fmt.Println("You chose option", option)
}
