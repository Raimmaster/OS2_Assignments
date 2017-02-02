package Disk

import(
	"fmt"
	"math"
)

func RoundToPowerOfTwo(number int) int{
	powerOfTwoOfNumber := math.Log2(float64(number))
	powerOfTwoOfNumber = math.Ceil(powerOfTwoOfNumber)
	comparisonNumber := math.Pow(2, powerOfTwoOfNumber)

	differenceOfNumbers := int(comparisonNumber) - number
	fmt.Printf("Power of two: %f Comparison number: %f, diff: %d \n", powerOfTwoOfNumber, comparisonNumber, differenceOfNumbers)
	if differenceOfNumbers > 0 {
		number += int(differenceOfNumbers)
	}

	return number
}
