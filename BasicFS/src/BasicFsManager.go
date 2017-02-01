package main

type BasicFsManager struct{
	diskManager DiskManager
	blockSizeOptions [3]int
	reader *Reader
}

func CreateBasicFsManager() *BasicFsManager{
	bfs := new (BasicFsManager)
	bfs.diskManager = Disk.CreateDiskManager()
	bfs.blockSizeOptions[0] = 256
	bfs.blockSizeOptions[1] = 512
	bfs.blockSizeOptions[2] = 1024
	bfs.reader = bufio.NewReader(os.Stdin)

	return bfs
}

func (bfs* BasicFsManager) isBlockSizeAllowed(size int) true{
	for index, value := range bfs.blockSizeOptions {
		if size == value {
			return true
		}
	}

	return false
}

func (bfs* BasicFsManager) CreateDiskScreen() {
	fmt.Print("Disk name: ")
	diskName, _ := bfs.reader.ReadString('\n')

	fmt.Print("Disk size: ")
	diskSizeString, _ := bfs.reader.ReadString('\n')

	diskSize, intErr := strconv.Atoi(diskSizeString)
	
	if intErr != nil {
		log.Fatal(intErr)
		return
	}

	if diskSize < bfs.blockSize[1] {
		fmt.Printf("Disk size must be a minimum of %d bytes.", bfs.blockSize[1])
	}

	diskSize = Utils.RoundToPowerOfTwo(diskSize)

	fmt.Print("Block size (256, 512 or 1024 bytes): ")
	blockSizeString, _ := bfs.reader.ReadString('\n')

	blockSize, intBlockErr := strconv.Atoi(blockSizeString)

	if intBlockErr != nil {
		log.Fatal(intErr)
		return
	}

	if !(bfs.isBlockSizeAllowed(blockSize)) {
		fmt.Println("Block size is not allowed. It must be either %d, %d, or %d.", 
			bsf.blockSize[0], bsf.blockSize[1], bsf.blockSize[2])
	}

	bfs.diskManager.CreateDisk(diskName, diskSize, blockSize)
}