package DiskScanner

import (
	"sort"
	"strconv"
	"strings"
	"math"
)

func cLook(addressesArray []string, headStart int) int {
	totalHeadMovements := 0
	currentHeadPosition := headStart
	intIndexes := toIntArray(addressesArray)

	i := getIndexNearTo(intIndexes,headStart)
	startedAt := i
	for {
		newHeadPos := intIndexes[i]
	    totalHeadMovements += int(math.Abs(float64(newHeadPos - currentHeadPosition)))
	    currentHeadPosition = newHeadPos

	    i = (i+1)%len(intIndexes)
	    if i == startedAt {
	    	totalHeadMovements -= int(math.Abs(float64(intIndexes[0] - intIndexes[len(intIndexes)-1])))
	    	break
	    }
	}

	return totalHeadMovements
}

func toIntArray(arr []string) []int {
	ints := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ints[i],_ = strconv.Atoi(strings.TrimSpace(arr[i]))
	}
	sort.Ints(ints)
	return reverseInts(ints)
}

func reverseInts(input []int) []int {
    if len(input) == 0 {
        return input
    }
    return append(reverseInts(input[1:]), input[0])
}

func getIndexNearTo(add []int, headStart int) int {
	for i := 0; i < len(add); i++ {
		if add[i] <= headStart {
			return i
		}
	}
	return 0
}
