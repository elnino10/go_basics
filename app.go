package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var years float64
	var investmentAmount float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment:", futureValue)
	fmt.Println("Future Real Value of Investment:", futureRealValue)
}
