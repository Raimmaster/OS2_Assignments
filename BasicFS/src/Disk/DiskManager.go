package Disk

import(
	"os"
	"fmt"
	"bufio"
)

type DiskManager struct {
	bufferedWriter *Writer
}

func CreateDiskManager() *DiskManager {
	dMang := new (DiskManager)

	return dMang
}

func (* DiskManager) CreateDisk(diskName string, diskSize int, blockSize int) {
	newFile, err := os.Create(diskName)
    if err != nil {
        log.Fatal(err)
    }
    newFile.Close()

    truncErr := os.Truncate(diskName, diskSize)    
	amountOfBlocks := diskSize / blockSize
	fmt.Printf("DS: %d BS: %d BC: %d", diskSize, blockSize, amountOfBlocks)
}

