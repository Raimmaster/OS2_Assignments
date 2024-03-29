package Disk

import(
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

type BasicFsManager struct{
	diskManager *DiskManager
	blockSizeOptions [3]int
	reader *bufio.Reader
}

func CreateBasicFsManager() *BasicFsManager{
	bfs := new (BasicFsManager)
	bfs.diskManager = CreateDiskManager()
	bfs.blockSizeOptions[0] = 256
	bfs.blockSizeOptions[1] = 512
	bfs.blockSizeOptions[2] = 1024
	bfs.reader = bufio.NewReader(os.Stdin)

	return bfs
}

func (bfs *BasicFsManager) isBlockSizeAllowed(size int) bool{
	for _, value := range bfs.blockSizeOptions {
		if size == value {
			return true
		}
	}

	return false
}

func (bfs *BasicFsManager) CreateDiskScreen() {
	fmt.Print("Disk name: ")
	diskName, _ := bfs.reader.ReadString('\n')
	diskName = strings.TrimSpace(string(diskName))

	fmt.Print("Disk size: ")
	diskSizeString, _ := bfs.reader.ReadString('\n')

	diskSize, intErr := strconv.Atoi(strings.TrimSpace(string(diskSizeString)))

	if intErr != nil {
		log.Fatal(intErr)
		return
	}

	if diskSize < bfs.blockSizeOptions[1] {
		fmt.Printf("Disk size must be a minimum of %d bytes.", bfs.blockSizeOptions[1])
		return
	}

	diskSize = RoundToPowerOfTwo(diskSize)

	fmt.Print("Block size (256, 512 or 1024 bytes): ")
	blockSizeString, _ := bfs.reader.ReadString('\n')

	blockSize, intBlockErr := strconv.Atoi(strings.TrimSpace(string(blockSizeString)))

	if intBlockErr != nil {
		log.Fatal(intErr)
		return
	}

	if !(bfs.isBlockSizeAllowed(blockSize)) {
		fmt.Println("Block size is not allowed. It must be either %d, %d, or %d.",
			bfs.blockSizeOptions[0], bfs.blockSizeOptions[1], bfs.blockSizeOptions[2])
	}

	bfs.diskManager.CreateDisk(diskName, diskSize, blockSize)
}

func (bfs *BasicFsManager) MountOrDismountDiskScreen(){
	canMountDisk := true
	if bfs.diskManager.HasMountedDisk() {
		fmt.Printf("Cannot mount another disk; disk %s is already mounted. Dismount to mount another.\n", bfs.diskManager.MountedDiskName)
		canMountDisk = false
	}

	fmt.Print("List disks: ")
	fmt.Println(ListFiles())
	fmt.Print("Type disk name to mount if none is mounted, or unmount if mounted: ")
	diskName, _ := bfs.reader.ReadString('\n')
	diskName = strings.TrimSpace(string(diskName))
	if canMountDisk{
		bfs.diskManager.MountDisk(diskName)
		return
	}

	if diskName == bfs.diskManager.MountedDiskName {
		bfs.diskManager.UnmountDisk(diskName)
	}
}

func (bfs *BasicFsManager) PrintBlocksInfoScreen() {
	if(bfs.diskManager.HasMountedDisk()){
		bfs.diskManager.PrintDiskInfo()
	}else {
		fmt.Println("No disk is mounted. Mount one to print its info.")
	}
}

func (bfs *BasicFsManager) AllocateBlockScreen() {
	fmt.Printf("Allocating block from disk %s\n", bfs.diskManager.MountedDiskName)
	bfs.diskManager.AllocateBlock()
}

func (bfs *BasicFsManager) FreeBlockScreen(){
	blockQuant := bfs.diskManager.mountedDisk.blockQuantity
	fmt.Printf("Block quantity: %d\n", blockQuant)
	fmt.Printf("Which block from 1 to %d will you free? ", (blockQuant - 1))
	blockToBeFreedString, _ := bfs.reader.ReadString('\n')
	blockToBeFreedString = strings.TrimSpace(string(blockToBeFreedString))

	blockToBeFreed, _ := strconv.Atoi(blockToBeFreedString)
	bfs.diskManager.FreeBlock(blockToBeFreed)
}

func ListFiles() string{
	var files_names string
	files_names = " \n"

	files, err := ioutil.ReadDir("./disks")

	if err != nil{
		fmt.Println("ERROR!")
		return " \n"
	}

	for _, file := range files {
		files_names += "* " + file.Name() + "\n"
	}

	return files_names
}
