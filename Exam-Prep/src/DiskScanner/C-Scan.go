package DiskScanner

import (
  "math"
  "strconv"
)

// const LEFT_LIMIT int = 0
// const RIGHT_LIMIT int = 199
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
  currentHeadPosition := headStart
  intAddressesArr := stringArrayToIntArray(addressesArray)

  for {
    var differencesArr [] int
    for i := 0; i < len(intAddressesArr); i++ {
      valueToSubstract := intAddressesArr[i]
      distance := int(math.Abs(float64(currentHeadPosition - valueToSubstract)))
      differencesArr = append(differencesArr, distance)
    }
    closestHeadPos := getMinAbsPos(differencesArr)
    totalHeadMovement += differencesArr[closestHeadPos]

    intAddressesArr = removeIntIndex(intAddressesArr, closestHeadPos)

    if len(intAddressesArr) == 0 {
      break
    }
  }

  return totalHeadMovement
}
