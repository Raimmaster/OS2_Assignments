package main

import (
    "os"
    "log"
  	"unsafe"
    "encoding/binary"
)


type Disk struct {
	diskSize int64
	blockSize int64
	freeSpace int64
	blockQuantity int64
	freeBlocks int64
	headBlock int64
	tailBlock int64
	// diskPath string
}

func main() {
    // Open file for reading
    disk := new (Disk)
    sizeOfDisk := int64(unsafe.Sizeof(*disk))
    log.Printf("%d   ", sizeOfDisk)
    file, err := os.Open("disks/love2.dk")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    diskSize := 8192
    byteSlice := make([]byte, diskSize)
    bytesRead, err := file.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", bytesRead)
    var value int64
    const incrementor int = 8
    // var initialVal int64 = 0
    for i := 0; i < diskSize; i+=incrementor {
      value, _ = binary.Varint(byteSlice[i:i + incrementor])
      log.Printf("Data read at byte %d: %d\n", i, value)
    }
}
