package main

import (
	"fmt"
	"math"
)

func calculate() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
