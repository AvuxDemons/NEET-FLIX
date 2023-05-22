package main

import (
	"NeetFlix/model"
	"NeetFlix/view"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("INSERT MOVIE")

	// ✅ Input & Insert Data
	var loop string
	for {
		if loop != "t" {
			view.InsertMovie()
			fmt.Print("| Ulang ? : ")
			fmt.Scan(&loop)
			scanner.Scan()
		} else {
			break
		}
	}

	// ✅ Delete Data
	model.DeleteData(2)

	// ✅ View Raw Data
	fmt.Println(model.ViewAllMovie())

	// ✅ View Data
	fmt.Println("\n---- VIEW ALL DATA ----")
	view.PrintAllMovie()
}
