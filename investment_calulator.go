package main

import (
	"fmt"
	"math"
)

func main() {
	// I was suggested to use 2.0% - as if!
	const inflationRate = 3.2
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5

	// Store the results of fmt.Scan into the value at the adress of investmentAmount
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInflatedValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureInflatedValue)
}
