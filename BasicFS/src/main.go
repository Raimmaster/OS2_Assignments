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
            //basicFsManager.AllocateBlockScreen()
            break;
        case 4:
            //basicFsManager.FreeBlockScreen()
        case 5:
            //basicFsManager.PrintBlocksInfo()
        default:
            fmt.Println("Wrong option selected.")
        }
}

func main() {
    basicFsManager := Disk.CreateBasicFsManager()
    for {
        menuOptions := "\n***BASIC FS***\n1. Create disk.\n2. Mount or dismount disk."
        menuOptions += "\n3. Allocate block. \n4. Liberar block."
        menuOptions += "\n5.Imprimir bloques y espacio libre."
        fmt.Println(menuOptions)
        fmt.Print("Your choice: ")
        reader := bufio.NewReader(os.Stdin)
        optionString, _ := reader.ReadString('\n')
        selectedOption, _ := strconv.Atoi(strings.TrimSpace(string(optionString)))
        showMenuSelection(selectedOption, basicFsManager)
    }
}
