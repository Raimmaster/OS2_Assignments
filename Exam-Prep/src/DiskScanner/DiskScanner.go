package DiskScanner

import (
  "math"
  "strings"
  "strconv"
)

type DiskScanner struct {

}

func New() *DiskScanner {
  return &DiskScanner{}
}
const LEFT_LIMIT int = 0
const RIGHT_LIMIT int = 199
const HEAD_START int = 53

func (diskScanner *DiskScanner) ScanByOption(addressesArray []string, scanOption int) int {
  switch scanOption {
    case 1:
      return fcfs(addressesArray, HEAD_START)
    case 2:
      return sstf(addressesArray, HEAD_START)
    case 3:
      return 0//scan(addressesArray, HEAD_START)
    case 4:
      return cScan(addressesArray, HEAD_START)
    case 5:
      return 0//cLook(addressesArray, HEAD_START)
  }
  return 0
}

func fcfs(addressesArray []string, headStart int) int {
  var totalHeadMovement int = 0
  var currentHeadPosition int = headStart
  for i := 0; i < len(addressesArray); i++ {
    newHeadPos, _ := strconv.Atoi(strings.TrimSpace(addressesArray[i]))
    totalHeadMovement += int(math.Abs(float64(newHeadPos - currentHeadPosition)))
    currentHeadPosition = newHeadPos
  }

  return totalHeadMovement
}

func scan(addressesArray []string, headStart int) int {
  var totalHeadMovement int =0
  var currentHeadPosition int = headStart
  var newPathArray []string
  if (currentHeadPosition - LEFT_LIMIT) < (RIGHT_LIMIT - currentHeadPosition) {

  }
}
