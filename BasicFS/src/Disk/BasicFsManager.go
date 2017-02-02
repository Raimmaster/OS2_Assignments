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

}

func (bfs *BasicFsManager) ListFiles() string{//ls
	var files_names string
	files_names = " \n"

	files, err := ioutil.ReadDir(".")

	if err != nil{
		return " \n"
	}

	for _, file := range files {
		files_names += file.Name() + "\n"
	}

	return files_names
}
