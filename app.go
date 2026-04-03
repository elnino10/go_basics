package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	years := 10.0
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println("Future Value of Investment:", futureValue)
}
