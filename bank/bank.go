package bank

import "fmt"

func Prompt() {
	fmt.Println("Welcome to the GO Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var option int
	fmt.Print("Your choice: ")
	fmt.Scan(&option)

	fmt.Println("You chose option", option)
}
