package user_input

import "fmt"

func GetInput(prompt string) (input float64) {
	fmt.Print(prompt)
	// Store the results of fmt.Scan into the value at the address of `input`
	fmt.Scan(&input)

	return input
}
