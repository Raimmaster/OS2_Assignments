package DiskScanner

import (
  "math"
  "strings"
  "strconv"
  "sort"
  //"fmt"
)

type DiskScanner struct {

}

func New() *DiskScanner {
  return &DiskScanner{}
}
const LEFT_LIMIT int = 0
const RIGHT_LIMIT int = 199
const HEAD_START int = 50

func (diskScanner *DiskScanner) ScanByOption(addressesArray []string, scanOption int) int {
  switch scanOption {
    case 1:
      return fcfs(addressesArray, HEAD_START)
    case 2:
      return sstf(addressesArray, HEAD_START)
    case 3:
      return scan(addressesArray, HEAD_START)
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
  var newPathArray []int
  var newEntries []int
  for _,element := range addressesArray {
    //fmt.Println(element)
    value, _ := strconv.Atoi(strings.TrimSpace(element))
    newEntries = append(newEntries, value)
  }
  sort.Ints(newEntries)
  if (currentHeadPosition - LEFT_LIMIT) < (RIGHT_LIMIT - currentHeadPosition) {
    for i :=len(newEntries)-1; i>=0; i-- {
      value := newEntries[i]
      if(value <= currentHeadPosition){
        newPathArray = append(newPathArray,value)
        newEntries = append(newEntries[:i], newEntries[i+1:]...)
      }
    }
    newPathArray = append(newPathArray,LEFT_LIMIT)
    for i :=0; i< len(newEntries); i++{
        newPathArray = append(newPathArray,newEntries[i])
    }
  }else{
    for i :=0; i<len(newEntries); i++ {
      value := newEntries[i]
      if(value >= currentHeadPosition){
        newPathArray = append(newPathArray,value)
        newEntries = append(newEntries[:i], newEntries[i+1:]...)
      }
    }
    newPathArray = append(newPathArray,RIGHT_LIMIT)
    for i :=len(newEntries)-1; i>=0 ; i--{
        newPathArray = append(newPathArray,newEntries[i])
    }
  }
  for _,element := range newPathArray {
    newHeadPos := element
    totalHeadMovement += int(math.Abs(float64(newHeadPos - currentHeadPosition)))
    currentHeadPosition = newHeadPos
  }
  return totalHeadMovement
}
