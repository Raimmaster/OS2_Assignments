package DiskScanner

import (
  "fmt"
  "math"
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
  var totalHeadMovements int = 0
  currentHeadPosition := headStart
  addressesArray = append(addressesArray, "0")
  addressesArray = append(addressesArray, "199")

  intAddressesArr := toIntArray(addressesArray)

  index := getIndexNearTo(intAddressesArr, headStart)
  fmt.Printf("Initial index: %d \n", index)
  startIndex := index
  for i := 0; i < len(intAddressesArr); i++ {
    fmt.Printf("Values: %d \n", intAddressesArr[i])
  }
  for {
    newHeadPos := intAddressesArr[index]

    totalHeadMovements += int(math.Abs(float64(newHeadPos - currentHeadPosition)))
    currentHeadPosition = newHeadPos

    index = (index + 1)%len(intAddressesArr)
    fmt.Printf("Index: %d \n", index)

    if index == startIndex {
      totalHeadMovements -= int(math.Abs(float64(intAddressesArr[0] - intAddressesArr[len(intAddressesArr)-1])))
      break
    }
  }

  return totalHeadMovements
}
