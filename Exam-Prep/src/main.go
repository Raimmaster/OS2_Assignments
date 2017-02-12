package main

import (
  "fmt"
  "strings"
  "DiskScanner"
)

func main(){
  for {
    fmt.Print("Insert the string of addresses to seek from 0 to 199 separated by '-': ")
    var addressesString string
    fmt.Scanf("%s", &addressesString)
    addressesArray := strings.Split(addressesString, "-")
    fmt.Print("Scan Options:\n1. FCFS.\n2. SSTF.\n3. SCAN.\n4. C-SCAN.\n5.C-LOOK.\nInsert your option:")
    var scanOption int
    fmt.Scanf("%d", &scanOption)

    var diskScanner DiskScanner
    totalCylinderMovement := diskScanner.ScanByOption(addressesArray, scanOption)
    fmt.Printf("Total cylinder movement was: %d \n", totalCylinderMovement)

    fmt.Print("Continue Y/N: ")
    var continueScanning string
    fmt.Scanf("%s", &continueScanning)

    if strings.ToLower(continueScanning) == "n" {
      break;
    }
  }
}
