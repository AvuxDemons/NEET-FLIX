package view

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ✅ Fungsi Clear Terminal
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ✅ Fungsi Delete 1 Baris
func reInput() {
	fmt.Print("\033[F")
	fmt.Print("\033[K")
}

func MainMenu() {
	var input int
	for input != 3 {

		fmt.Println("========================================")
		fmt.Println("███╗   ██╗███████╗███████╗████████╗███████╗██╗     ██╗██╗  ██╗")
		fmt.Println("████╗  ██║██╔════╝██╔════╝╚══██╔══╝██╔════╝██║     ██║╚██╗██╔╝")
		fmt.Println("██╔██╗ ██║█████╗  █████╗     ██║   █████╗  ██║     ██║ ╚███╔╝ ")
		fmt.Println("██║╚██╗██║██╔══╝  ██╔══╝     ██║   ██╔══╝  ██║     ██║ ██╔██╗ ")
		fmt.Println("██║ ╚████║███████╗███████╗   ██║   ██║     ███████╗██║██╔╝ ██╗")
		fmt.Println("╚═╝  ╚═══╝╚══════╝╚══════╝   ╚═╝   ╚═╝     ╚══════╝╚═╝╚═╝  ╚═╝")
		fmt.Println("========================================")
		fmt.Println("| 1. Manage User  ")
		fmt.Println("| 2. Manage Movie  ")
		fmt.Println("| 3. Exit  ")
		fmt.Println("========================================")
		fmt.Print("| Input : ")
		fmt.Scan(&input)
		fmt.Println("========================================")

		if input < 3 && input > 0 {
			clearScreen()
			switch input {
			case 1:
				MenuUser()
			case 2:
				MenuMovie()
			}
		} else if input != 3 {
			fmt.Println("|  INVALID INPUT! ")
		} else {
			fmt.Println("|  Exiting Program ")
			break
		}
	}
}
