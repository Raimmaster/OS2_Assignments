package Disk

import(
	"os"
	"fmt"
	"log"
	"unsafe"
	"encoding/binary"
)

const BEGINNING_OF_FILE int = 0
const DISK_DIR string = "./disks/"
const POINTER_SIZE int64 = 8
const METADATA_SIZE int64 = 56

type DiskManager struct {
	mountedDisk *Disk
	mountedDFile *os.File
	MountedDiskName string
}

type Disk struct {
	diskSize int64
	blockSize int64
	freeSpace int64
	blockQuantity int64
	freeBlocks int64
	headBlock int64
	tailBlock int64
	diskName string
}

func CreateDiskManager() *DiskManager {
	dMang := new (DiskManager)

	return dMang
}

const STRING_POINTER_SIZE int64 = 16

func (disk *Disk) toBytesArray() []byte{
	sizeOfDisk := int64(unsafe.Sizeof(*disk)) - STRING_POINTER_SIZE
	byteSlice := make([]byte, sizeOfDisk)

	binary.PutVarint(byteSlice[0:8], disk.diskSize)
	binary.PutVarint(byteSlice[8:16], disk.blockSize)
	binary.PutVarint(byteSlice[16:24], disk.freeSpace)
	binary.PutVarint(byteSlice[24:32], disk.blockQuantity)
	binary.PutVarint(byteSlice[32:40], disk.freeBlocks)
	binary.PutVarint(byteSlice[40:48], disk.headBlock)
	binary.PutVarint(byteSlice[48:56], disk.tailBlock)

	return byteSlice
}

func (disk *Disk) obtainMetaFromDisk(byteSlice []byte){
	disk.diskSize, _ = binary.Varint(byteSlice[0:8])
	disk.blockSize, _ = binary.Varint(byteSlice[8:16])
	disk.freeSpace, _ =  binary.Varint(byteSlice[16:24])
	disk.blockQuantity, _ = binary.Varint(byteSlice[24:32])
	disk.freeBlocks, _ = binary.Varint(byteSlice[32:40])
	disk.headBlock, _ = binary.Varint(byteSlice[40:48])
	disk.tailBlock, _ = binary.Varint(byteSlice[48:56])
}

func (diskManager *DiskManager) writeDiskMetadata(disk *Disk, diskName string) {
	diskFile, _ := os.OpenFile(DISK_DIR + diskName, os.O_WRONLY, 0666)
	defer diskFile.Close()
	setDiskFilePointer(diskFile)
	disk.freeSpace = disk.diskSize - disk.blockSize
	disk.blockQuantity = disk.diskSize / disk.blockSize
	disk.freeBlocks = disk.blockQuantity - 1
	disk.headBlock = 1
	disk.tailBlock = disk.blockQuantity - 1

	bytesWritten, err := diskFile.Write(disk.toBytesArray())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes written. \n", bytesWritten)
}

func setDiskFilePointer(diskFile *os.File){
	var offset int64 = 0
	var whence int = 0
	diskFile.Seek(offset, whence)
}

func (diskManager *DiskManager) initBlockList(disk *Disk, diskPath string) {
	diskFile, _ := os.OpenFile(diskPath, os.O_WRONLY, 0666)
	defer diskFile.Close()
	setDiskFilePointer(diskFile)
	diskManager.mountedDisk = disk
	var BLOCK_POINTER_OFFSET int64 = disk.blockSize - POINTER_SIZE
	var blockIndex int64
	pointerOfNextBlock := make([]byte, POINTER_SIZE)
	for blockIndex = 1; blockIndex < disk.blockQuantity - 1; blockIndex++ {
		binary.PutVarint(pointerOfNextBlock[0:8], (blockIndex + 1))
		diskManager.WriteBlock(blockIndex, BLOCK_POINTER_OFFSET, pointerOfNextBlock)
	}
	//write final block to point to nil (-1)
	binary.PutVarint(pointerOfNextBlock[0:8], (-1))
	diskManager.WriteBlock(blockIndex, BLOCK_POINTER_OFFSET, pointerOfNextBlock)
	diskManager.mountedDisk = nil
}

func (diskManager *DiskManager) CreateDisk(diskName string, diskSize int, blockSize int) {
	diskPath := DISK_DIR + diskName
	newFile, err := os.Create(diskPath)
	if err != nil {
      log.Fatal(err)
  }
  newFile.Close()
	disk := new (Disk)
	disk.diskName = diskName
	fmt.Printf("DS: %d BS: %d\n", diskSize, blockSize)
	disk.diskSize = int64(diskSize)
	disk.blockSize = int64(blockSize)
	diskManager.writeDiskMetadata(disk, diskName)
  truncErr := os.Truncate(diskPath, int64(diskSize))
	diskManager.initBlockList(disk, diskPath)

  if truncErr != nil {
  	log.Fatal(truncErr)
  }
}

func (diskManager *DiskManager) MountDisk(diskName string) {
	var err error
	diskManager.mountedDFile, err = os.Open(DISK_DIR + diskName)
	diskManager.mountedDisk = new (Disk)
	diskManager.mountedDisk.diskName = diskName
	diskManager.MountedDiskName = diskName
	if err != nil {
		log.Printf("Failed to mount disk. %s", err)
		return
	}
	byteSliceOfMeta := make([]byte, METADATA_SIZE)
	diskManager.ReadBlock(0, 0, byteSliceOfMeta)
	diskManager.mountedDisk.obtainMetaFromDisk(byteSliceOfMeta)
	log.Printf("Disk %s mounted", diskName)
}

func (diskManager *DiskManager) PrintDiskInfo() {
	fmt.Printf("Disk Name: %s\n", diskManager.mountedDisk.diskName)
	fmt.Printf("Disk Size: %d\n", diskManager.mountedDisk.diskSize)
	fmt.Printf("Free disk space: %d\n", diskManager.mountedDisk.freeSpace)
	fmt.Printf("Block Amount: %d\n", diskManager.mountedDisk.blockQuantity)
	fmt.Printf("Free blocks: %d\n", diskManager.mountedDisk.freeBlocks)
}

func (diskManager *DiskManager) UnmountDisk(diskName string) {
	diskManager.mountedDFile.Close()
	diskManager.mountedDisk = nil
	diskManager.mountedDFile = nil
	diskManager.MountedDiskName = ""

	fmt.Printf("Disk %s successfully unmounted. \n", diskName)
}

func (diskManager *DiskManager) HasMountedDisk() bool{
	return diskManager.mountedDisk != nil
}

func (disk *Disk) Seek(offset int64, whence int, diskFile *os.File){
	diskFile.Seek(offset, whence)
}

func (disk *Disk) Write(buffer []byte, diskFile *os.File){
	bytesWritten, err := diskFile.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written in block: %d", bytesWritten)
}

func (disk *Disk) Read(buffer []byte, diskFile *os.File){
	_, err := diskFile.Read(buffer)
	if err != nil {
		log.Printf("Failed to read: %s", err)
		return
	}
}

func (diskManager *DiskManager) WriteBlock(blockNumber int64, blockOffset int64, buffer []byte) {
	offset := blockNumber * diskManager.mountedDisk.blockSize + blockOffset
	diskFile, _ := os.OpenFile(DISK_DIR + diskManager.mountedDisk.diskName, os.O_WRONLY, 0666)
	defer diskFile.Close()
	diskManager.mountedDisk.Seek(offset, BEGINNING_OF_FILE, diskFile)
	diskManager.mountedDisk.Write(buffer, diskFile)
}

func (diskManager *DiskManager) ReadBlock(blockNumber int64, blockOffset int64, buffer []byte) {
	offset := blockNumber * diskManager.mountedDisk.blockSize + blockOffset
	diskFile, _ := os.OpenFile(DISK_DIR + diskManager.mountedDisk.diskName, os.O_RDONLY, 0666)
	defer diskFile.Close()
	diskManager.mountedDisk.Seek(offset, BEGINNING_OF_FILE, diskFile)
	diskManager.mountedDisk.Read(buffer, diskFile)
}
