package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the GO Finance CLI!")

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Bank")
		fmt.Println("2. Profit Calculator")
		fmt.Println("3. Investment Calculator")
		fmt.Println("4. Exit")

		option := GetInput("Your choice: ")

		switch option {
		case 1:
			Bank()
		case 2:
			Profit()
		case 3:
			Investment()
		case 4:
			fmt.Println("Goodbye! Please come again :)")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
			continue
		}
	}
}
