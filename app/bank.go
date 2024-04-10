package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFileName = "balance.txt"

func getBalance() (float64, error) {
	data, err := os.ReadFile(accountFileName)

	if err != nil {
		writeBalance(100.0)
		return 100.0, errors.New("account balance file not found. Default balance of $100.00 set")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		writeBalance(100.0)
		return 100.0, errors.New("failed to parse account balance. Default balance of $100.00 set")
	}
	return balance, nil
}

func writeBalance(balance float64) {
	balanceBytes := []byte(fmt.Sprintf("%.2f", balance))
	err := os.WriteFile(accountFileName, balanceBytes, 0644)
	if err != nil {
		// panic causes a hard exit as opposed to graceful error handling
		panic("Can't write balance to file! Aborting application!")
	}
}

func Bank() {
	var accountBalance, err = getBalance()

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		fmt.Println("\n-------------")
	}

	fmt.Println("Welcome to the GO Bank!")

loop: //label the loop to allow break to exit from switch statement
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		opt := int(GetInput("Your choice: "))

		switch opt {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			deposit := GetInput("Enter the amount you want to deposit: ")
			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0. Please try again.")
				continue
			}
			accountBalance += deposit
			writeBalance(accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		case 3:
			withdraw := GetInput("Enter the amount you want to withdraw: ")
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
