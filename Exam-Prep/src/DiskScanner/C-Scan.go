package DiskScanner

/*import (
  "math"
  "strings"
  "strconv"
)*/

import "math"

// const LEFT_LIMIT int = 0
// const RIGHT_LIMIT int = 199
// const HEAD_START int = 53

func cScan(addressesArray []string, headStart int) int {
  var totalHeadMovement int = 0
  var currentHeadPosition int = headStart

  for i := 0; i < len(addressesArray); i++ {
    newHeadPos := 0
    totalHeadMovement += int(math.Abs(float64(newHeadPos - currentHeadPosition)))
  }

  return totalHeadMovement
}
