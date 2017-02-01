package main

import (
    "log"
    "fmt"
    "bufio"
    "strconv"
)

func showMenuSelection(){
    switch selectedOption {
        case 1:
            basicFsManager.CreateDiskScreen()
            break;
        case 2:
            //basicFsManager.MountOrDismountDiskScreen()
            break;
        case 3:
            //basicFsManager.WriteBlockScreen()
            break;
        case 4:
            //basicFsManager.ReadBlockScreen()
            break;
        case 5:
            //basicFsManager.AllocateBlockScreen()
            break;
        case 6:
            //basicFsManager.FreeBlockScreen()
        case 7:
            //basicFsManager.PrintBlocksInfo()
        default: 
            fmt.Println("Wrong option selected.")
        }
}

func main() {
    basicFsManager, errDiskManager := new (BasicFsManager)
    for {
        menuOptions := "***BASIC FS***\n1. Create disk.\n2. Mount or dismount disk.
            \n3. Write block.\n4. Read block.
            \n5. Allocate block. \n6. Liberar block. 
            \n7.Imprimir bloques y espacio libre."
        fmt.Println(menuOptions)
        fmt.Print("Your choice: ")
        reader := bufio.NewReader(os.Stdin)
        optionString, bufErr := reader.ReadString('\n')
        selectedOption, err := strconv.Atoi(optionString)

        showMenuSelection(selectedOption)
    }    
}