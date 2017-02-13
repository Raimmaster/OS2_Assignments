package DiskScanner

import (
  "fmt"
  "math"
  "sort"
  "strconv"
)

// const HEAD_START int = 53

func stringArrayToIntArray(addressesArray []string) []int{
  var intArr []int
  for i := 0; i < len(addressesArray); i++ {
    intVal, _ := strconv.Atoi(addressesArray[i])
    intArr = append(intArr, intVal)
  }

  return intArr
}

func cScan(addressesArray []string, headStart int) int {
  var totalHeadMovement int = 0
  isMovingLeft := true
  currentHeadPosition := headStart
  intAddressesArr := stringArrayToIntArray(addressesArray)
  sort.Ints(intAddressesArr)

  var differencesArr [] int
  for i := 0; i < len(intAddressesArr); i++ {
    valueToSubstract := intAddressesArr[i]
    distance := int(math.Abs(float64(valueToSubstract - currentHeadPosition)))
    differencesArr = append(differencesArr, distance)
  }

  closestHeadPos := getMinAbsPos(differencesArr)
  newHeadPosition := intAddressesArr[closestHeadPos]
  isMovingLeft = newHeadPosition < currentHeadPosition

  for {
    // fmt.Printf("Got differences: %d \n", len(differencesArr))
    differencesArr = differencesArr[:0]
    for i := 0; i < len(intAddressesArr); i++ {
      valueToSubstract := intAddressesArr[i]
      fmt.Printf("New val to add to substract: %d \n", valueToSubstract)
      distance := int(math.Abs(float64(valueToSubstract - currentHeadPosition)))
      fmt.Printf("New distance VAL: %d \n", distance)
      differencesArr = append(differencesArr, distance)
    }

    closestHeadPos = getMinAbsPos(differencesArr)
    newHeadPosition = intAddressesArr[closestHeadPos]
    fmt.Printf("Current: %d ", currentHeadPosition)
    var pausing int
    fmt.Scanf("%d", &pausing)
    if isMovingLeft {
      // fmt.Printf("Moving left with NH: %d and CH: %d \n", newHeadPosition, currentHeadPosition)
      if newHeadPosition > currentHeadPosition {
        isMovingLeft = false
        totalHeadMovement += currentHeadPosition
        currentHeadPosition = RIGHT_LIMIT
        continue
      }
    } else {
      // fmt.Printf("Moving right with New: %d and Current: %d \n", newHeadPosition, currentHeadPosition)
      if newHeadPosition < currentHeadPosition {
        isMovingLeft = true
        totalHeadMovement += int(math.Abs(float64(currentHeadPosition - RIGHT_LIMIT)))
        currentHeadPosition = LEFT_LIMIT
        continue
      }
    }

    totalHeadMovement += differencesArr[closestHeadPos]
    currentHeadPosition = newHeadPosition
    intAddressesArr = removeIntIndex(intAddressesArr, closestHeadPos)

    fmt.Printf("Current int arr length: %d \n", len(intAddressesArr))
    if len(intAddressesArr) == 0 {
      break
    }

  }

  return totalHeadMovement
}
