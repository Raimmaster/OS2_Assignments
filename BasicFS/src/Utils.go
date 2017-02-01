package main

import(
	"math"
)

funct RoundToPowerOfTwo(number int) int{
	powerOfTwoOfNumber := math.Log2(number)
	powerOfTwoOfNumber = math.Floor(powerOfTwoOfNumber)
	comparisonNumber := math.Pow(2, powerOfTwoOfNumber)

	differenceOfNumbers = comparisonNumber - number
	if differenceOfNumbers > 0 {
		number += int(differenceOfNumbers)
	}
	
	return number
}