package DiskScanner

import (
	"fmt"
	"strconv"
)

func removeIntIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func removeStringIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func absolute(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func getMinAbsPos(s []int) int {
	min := 0
	for i := 1; i < len(s); i++ {
		if absolute(s[i]) < absolute(s[min]) {
			min = i
		}
	}
	return min
}

func sstf(addressesArray []string, headStart int) int {
	var differences []int
	movement := 0
	for i := 0; i < len(addressesArray); i++ {
		arrayValue, _ := strconv.Atoi(addressesArray[i])
		differences = append(differences, arrayValue-headStart)
		fmt.Printf("Difference[%d]: %d\n", i, differences[i])
	}

	for {

		nextRequestPos := getMinAbsPos(differences)
		nextRequestDif := differences[nextRequestPos]

		requestValue, _ := strconv.Atoi(addressesArray[nextRequestPos])

		fmt.Printf("nextRequestPos: %d nextRequestDif: %d RequestValue: %d\n", nextRequestPos, nextRequestDif, requestValue)

		movement += absolute(nextRequestDif)

		for i := 0; i < len(differences); i++ {
			differences[i] = differences[i] - nextRequestDif
		}

		differences = removeIntIndex(differences, nextRequestPos)
		addressesArray = removeStringIndex(addressesArray, nextRequestPos)

		if len(differences) == 0 {
			break
		}
	}

	return movement
}
