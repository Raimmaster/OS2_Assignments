package DiskScanner

type DiskScanner struct {

}

const HEAD_START int = 53

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
      return cLook(addressesArray, HEAD_START)
  }
}

func fcfs(addressesArray []string, headStart int) int {
  
}
