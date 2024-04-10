package main

import "fmt"

func GetInput(prompt string) (input float64) {
	fmt.Println(prompt)

	// Store the results of fmt.Scan into the value at the address of `input`
	fmt.Scan(&input)

	return input
}
