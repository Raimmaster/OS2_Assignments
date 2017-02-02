package main

import (
    "os"
    "fmt"
    "Disk"
    "bufio"
    "strings"
    "strconv"
)

func showMenuSelection(selectedOption int, basicFsManager *Disk.BasicFsManager){
    switch selectedOption {
        case 1:
            basicFsManager.CreateDiskScreen()
            break;
        case 2:
            basicFsManager.MountOrDismountDiskScreen()
            break;
        case 3:
            basicFsManager.AllocateBlockScreen()
            break;
        case 4:
            basicFsManager.FreeBlockScreen()
            break;
          case 5:
            basicFsManager.PrintBlocksInfoScreen()
            break;
        default:
            fmt.Println("Wrong option selected.")
        }
}

func main() {
    basicFsManager := Disk.CreateBasicFsManager()
    for {
        menuOptions := "\n***BASIC FS***\n1. Create disk.\n2. Mount or dismount disk."
        menuOptions += "\n3. Allocate block. \n4. Free block."
        menuOptions += "\n5. Print disk info."
        fmt.Println(menuOptions)
        fmt.Print("Your choice: ")
        reader := bufio.NewReader(os.Stdin)
        optionString, _ := reader.ReadString('\n')
        selectedOption, _ := strconv.Atoi(strings.TrimSpace(string(optionString)))
        fmt.Printf("\n\n\n")
        showMenuSelection(selectedOption, basicFsManager)
    }
}
