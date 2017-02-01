package Disk

import(
	"math"
)

func RoundToPowerOfTwo(number int) int{
	powerOfTwoOfNumber := math.Log2(float64(number))
	powerOfTwoOfNumber = math.Floor(powerOfTwoOfNumber)
	comparisonNumber := math.Pow(2, powerOfTwoOfNumber)

	differenceOfNumbers := int(comparisonNumber) - number
	if differenceOfNumbers > 0 {
		number += int(differenceOfNumbers)
	}
	
	return number
}