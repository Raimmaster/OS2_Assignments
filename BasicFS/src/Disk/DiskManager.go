package Disk

import(
	"os"
	"fmt"
	"log"
	"bufio"
)

type DiskManager struct {
	bufferedWriter *bufio.Writer
}

func CreateDiskManager() *DiskManager {
	dMang := new (DiskManager)

	return dMang
}

func (diskManager *DiskManager) CreateDisk(diskName string, diskSize int, blockSize int) {
	newFile, err := os.Create(diskName)
    if err != nil {
        log.Fatal(err)
    }
    diskManager.bufferedWriter = bufio.NewWriter(newFile)
    newFile.Close()

    truncErr := os.Truncate(diskName, int64(diskSize))
    if truncErr != nil {
    	log.Fatal(truncErr)
    }    
	amountOfBlocks := diskSize / blockSize
	fmt.Printf("DS: %d BS: %d BC: %d", diskSize, blockSize, amountOfBlocks)
}

