package main

import (
	"DiskScanner"
	"fmt"
	"strings"
)

func testSSTF() {
	diskScanner := DiskScanner.New()
	addresses := []string{"98", "183", "37", "122", "14", "124", "65", "67"}
	totalCylinderMovement := diskScanner.ScanByOption(addresses, 2)
	fmt.Printf("Total cylinder movement was: %d \n", totalCylinderMovement)
}
func testCLook() {
	diskScanner := DiskScanner.New()
	addresses := []string{"11","34","50","62","64","95","119","123","180"}
	totalCylinderMovement := diskScanner.ScanByOption(addresses, 5)
	fmt.Printf("Total cylinder movement was: %d \n", totalCylinderMovement)
}

func testCscan() {
	diskScanner := DiskScanner.New()
	addresses := []string{"95", "180", "34", "119", "11", "123", "62", "64"}
	totalCylinderMovement := diskScanner.ScanByOption(addresses, 4)
	fmt.Printf("Total cylinder movement was: %d \n", totalCylinderMovement)
}

func main() {
  // testSSTF()
	testCscan()
  for {
		fmt.Print("Insert the string of addresses to seek from 0 to 199 separated by '-': ")
		var addressesString string
		fmt.Scanf("%s", &addressesString)
		addressesArray := strings.Split(addressesString, "-")
		fmt.Print("Scan Options:\n1. FCFS.\n2. SSTF.\n3. SCAN.\n4. C-SCAN.\n5 .C-LOOK.\nInsert your option:")
		var scanOption int
		fmt.Scanf("%d", &scanOption)

		diskScanner := DiskScanner.New()
		totalCylinderMovement := diskScanner.ScanByOption(addressesArray, scanOption)
		fmt.Printf("Total cylinder movement was: %d \n", totalCylinderMovement)

		fmt.Print("Continue Y/N: ")
		var continueScanning string
		fmt.Scanf("%s", &continueScanning)

		if strings.ToLower(continueScanning) == "n" {
			break
		}
	}
}
