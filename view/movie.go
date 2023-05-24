package view

import (
	"NeetFlix/controller"
	"NeetFlix/database"
	"NeetFlix/model"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

// ✅ Fungsi Input Data Movie
func InsertMovie() {
	scanner := bufio.NewScanner(os.Stdin)

	var Release, Id int
	var Title, Studio, Agerate string
	var Category, Genre []string
	var Rating float64

	// Id
	database.IDMovie += 1
	Id = database.IDMovie

	// Judul
	fmt.Println("=================================")
	fmt.Print("| Judul       : ")
	scanner.Scan()
	Title = scanner.Text()

	// Tahun Rilis
	fmt.Println("=================================")
	fmt.Print("| Tahun Rilis : ")
	fmt.Scan(&Release)

	// Genre
	fmt.Println("=================================")
	fmt.Println("| Genre	")
	fmt.Println("|--------------------------------")
	for i, genre := range database.GenreList {
		fmt.Printf("| %d. %v \n", i+1, genre)
	}
	fmt.Printf("| %d. Exit \n", len(database.GenreList)+1)
	fmt.Println("|--------------------------------")
	var Gen int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Gen)
		reInput()

		if Gen < 1 || Gen > len(database.GenreList)+1 {
			fmt.Println("|   ! Invalid Input !")
		} else if Gen == len(database.GenreList)+1 {
			if len(Genre) == 0 {
				fmt.Println("|   ! Tambahkan Minimal 1 Genre !")
			} else {
				break
			}
		} else {
			selectedGenre := database.GenreList[Gen-1]
			if strings.Contains(strings.Join(Genre, ""), selectedGenre) {
				fmt.Println("| Genre Sudah Dipilih, Silahkan Tambahkan Genre Lain !")
			} else {
				fmt.Println("| -", selectedGenre, "Telah Ditambahkan")
				Genre = append(Genre, selectedGenre)
			}
		}
	}

	// Kategori
	fmt.Println("=================================")
	fmt.Println("| Kategori")
	fmt.Println("|--------------------------------")
	for k, category := range database.CategoryList {
		fmt.Printf("| %d. %v \n", k+1, category)
	}
	fmt.Printf("| %d. Exit \n", len(database.CategoryList)+1)
	fmt.Println("|--------------------------------")
	var Cat int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Cat)
		reInput()

		if Cat < 1 || Cat > len(database.CategoryList)+1 {
			fmt.Println("| ! Invalid Input !")
		} else if Cat == len(database.CategoryList)+1 {
			if len(Category) == 0 {
				fmt.Println("| ! Tambahkan Minimal 1 Category !")
			} else {
				break
			}
		} else {
			selectedCategory := database.CategoryList[Cat-1]
			if strings.Contains(strings.Join(Category, ""), selectedCategory) {
				fmt.Println("| ! Category Sudah Dipilih, Silahkan Tambahkan Category Lain !")
			} else {
				fmt.Println("| -", selectedCategory, "Telah Dipilih")
				Category = append(Category, selectedCategory)
			}
		}
	}

	scanner.Scan()

	// Studio
	fmt.Println("=================================")
	fmt.Print("| Studio      : ")
	scanner.Scan()
	Studio = scanner.Text()

	// Age Rate
	fmt.Println("=================================")
	fmt.Println("| Batasan Umur")
	fmt.Println("|--------------------------------")
	for i := range database.AgeRate {
		fmt.Printf("| %d. %v \n", i+1, database.AgeRate[i])
	}
	fmt.Println("|--------------------------------")
	var Age int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Age)
		reInput()
		if Age > len(database.AgeRate) || Age < 1 {
			fmt.Println("| ! Invalid Input !")
		} else {
			fmt.Println("| - ", database.AgeRate[Age-1], "Telah Dipilih")
			Agerate = database.AgeRate[Age-1]
			break
		}
	}

	// Rating
	fmt.Println("=================================")
	for {
		fmt.Print("| Rating      : ")
		fmt.Scan(&Rating)
		if Rating > 100 || Rating < 1 {
			reInput()
			fmt.Println("|  Invalid Input !")
		} else {
			break
		}
	}
	controller.InsertMovie(Id, Title, Release, Category, Studio, Genre, Agerate, Rating)
}

// ██    ██ ██████  ██████   █████  ████████ ███████
// ██    ██ ██   ██ ██   ██ ██   ██    ██    ██
// ██    ██ ██████  ██   ██ ███████    ██    █████
// ██    ██ ██      ██   ██ ██   ██    ██    ██
//  ██████  ██      ██████  ██   ██    ██    ███████

func UpdateMovie() {
	scanner := bufio.NewScanner(os.Stdin)

	var Temp int

	fmt.Println("Masukkan ID movie yang akan diubah : ")
	fmt.Scan(&Temp)
	scanner.Scan()

	var Release, Id int
	var Title, Studio, Agerate string
	var Category, Genre []string
	var Rating float64

	if model.CheckData(Temp) {

		Params := strconv.Itoa(Temp)
		PrintMovie(Params)

		// EDIT DATA

		// Judul
		fmt.Println("=================================")
		fmt.Print("| Judul       : ")
		scanner.Scan()
		Title = scanner.Text()
		if Title == "-" {
			fmt.Println("| Data Tidak Berubah")
		}

		// Tahun Rilis
		var releases string
		fmt.Println("=================================")
		fmt.Print("| Tahun Rilis : ")
		fmt.Scan(&releases)
		if releases != "-" {
			x, _ := strconv.Atoi(releases)
			Release = x
		} else {
			fmt.Println("| Data Tidak Berubah")
			Release = 0
		}

		// Genre
		fmt.Println("=================================")
		fmt.Println("| Genre	")
		fmt.Println("|--------------------------------")
		for i, genre := range database.GenreList {
			fmt.Printf("| %d. %v \n", i+1, genre)
		}
		fmt.Printf("| %d. Exit \n", len(database.GenreList)+1)
		fmt.Println("|--------------------------------")
		var genres string
		for {
			fmt.Print("|  Pilih : ")
			fmt.Scan(&genres)
			reInput()

			x, _ := strconv.Atoi(genres)
			if x == len(database.GenreList)+1 || genres == "-" {
				if genres == "-" {
					fmt.Println("| Data Tidak Berubah")
				}
				break
			} else if x > len(database.GenreList)+1 {
				fmt.Println("| ! Invalid Input !")
			} else {
				if strings.Contains(strings.Join(Genre, ""), database.GenreList[x-1]) {
					fmt.Println("| ! Genre Sudah Dipilih , Silahkan Tambahkan Genre Lain !")
				} else {
					fmt.Println("| -", database.GenreList[x-1], "Telah Ditambahkan")
					Genre = append(Genre, database.GenreList[x-1])
				}
			}
		}

		// Kategori
		fmt.Println("=================================")
		fmt.Println("| Kategori")
		fmt.Println("|--------------------------------")
		for k, category := range database.CategoryList {
			fmt.Printf("| %d. %v \n", k+1, category)
		}
		fmt.Printf("| %d. Exit \n", len(database.CategoryList)+1)
		fmt.Println("|--------------------------------")
		for {
			var cats string

			fmt.Print("| Pilih : ")
			fmt.Scan(&cats)

			reInput()

			if cats != "-" {
				x, _ := strconv.Atoi(cats)
				if x < 1 || x > len(database.CategoryList)+1 {
					fmt.Println("| ! Invalid Input !")
				} else if x == len(database.CategoryList)+1 {
					if len(Category) == 0 {
						fmt.Println("| ! Tambahkan Minimal 1 Category !")
					} else {
						break
					}
				} else {
					selectedCategory := database.CategoryList[x-1]
					if strings.Contains(strings.Join(Category, ""), selectedCategory) {
						fmt.Println("| ! Category Sudah Dipilih, Silahkan Tambahkan Category Lain !")
					} else {
						fmt.Println("| -", selectedCategory, "Telah Dipilih")
						Category = append(Category, selectedCategory)
					}
				}
			} else {
				fmt.Println("| Data Tidak Berubah")
				break
			}
		}

		scanner.Scan()

		// Studio
		fmt.Println("=================================")
		fmt.Print("| Studio      : ")
		scanner.Scan()
		Studio = scanner.Text()
		if Studio == "-" {
			fmt.Println("| Data Tidak Berubah")
		}

		// Age Rate
		fmt.Println("=================================")
		fmt.Println("| Batasan Umur")
		fmt.Println("|--------------------------------")
		for i := range database.AgeRate {
			fmt.Printf("| %d. %v \n", i+1, database.AgeRate[i])
		}
		fmt.Println("|--------------------------------")
		var ages string
		for {
			fmt.Print("| Pilih : ")
			fmt.Scan(&ages)

			reInput()

			if ages != "-" {
				x, _ := strconv.Atoi(ages)
				if x > len(database.AgeRate) || x < 1 {
					fmt.Println("| ! Invalid Input !")
				} else {
					fmt.Println("| -", database.AgeRate[x-1], "Telah Dipilih")
					Agerate = database.AgeRate[x-1]
					break
				}
			} else {
				Agerate = "-"
				fmt.Println("| Data Tidak Berubah")
				break
			}
		}

		// Rating
		fmt.Println("=================================")
		for {
			fmt.Print("| Rating      : ")
			var ratings string
			fmt.Scan(&ratings)

			reInput()

			if ratings != "-" {
				x, _ := strconv.ParseFloat(ratings, 64)
				if x > 100 || x < 1 {
					fmt.Println("| ! Invalid Input !")
				} else {
					Rating = x
					break
				}
			} else {
				Rating = 0
				fmt.Println("| Data Tidak Berubah")
				break
			}
		}

		controller.UpdateMovie(Id, Title, Release, Category, Studio, Genre, Agerate, Rating)

	} else {
		fmt.Println("| Data Masih Kosong")
	}

}

func ViewByMenu() {
	fmt.Println("| PILIH BERDASARKAN")
	fmt.Println("| 1. ID ")
	fmt.Println("| 2. Judul ")
	fmt.Println("| 3. Kategori ")
	fmt.Println("| 4. Genre ")

	var x int
	fmt.Println("========================================")
	fmt.Print("| PILIH : ")
	fmt.Scan(&x)
	ViewBy(x)
}

func ViewBy(selector int) {
	scanner := bufio.NewScanner(os.Stdin)

	switch selector {
	case 3:
		for i, category := range database.CategoryList {
			fmt.Printf("| %d. %v \n", i+1, category)
		}
	case 4:
		for i, genre := range database.GenreList {
			fmt.Printf("| %d. %v \n", i+1, genre)
		}
	}
	fmt.Println("========================================")
	fmt.Print("| Search Query :  ")
	scanner.Scan()
	scanner.Scan()
	Params := scanner.Text()

	switch selector {
	case 3:
		x, _ := strconv.Atoi(Params)
		Params = database.CategoryList[x-1]
	case 4:
		x, _ := strconv.Atoi(Params)
		Params = database.GenreList[x-1]
	}

	PrintMovie(Params)
}

func PrintMovie(params string) {
	data := model.ViewBy(params)
	for _, temp := range data {
		fmt.Println("| ID \t\t:", temp.Id)
		fmt.Println("| Title\t\t:", temp.Title)
		fmt.Println("| Release\t:", temp.Release)
		fmt.Println("| Categoty\t:", temp.Category)
		fmt.Println("| Studio\t:", temp.Studio)
		fmt.Println("| Genre\t\t:", temp.Genre)
		fmt.Println("| Age Rate\t:", temp.Agerate)
		fmt.Println("| Rating\t:", temp.Rating)
	}
}

// ✅
func PrintAllMovie() {
	data := model.ViewAllMovie()
	length := len(data)
	for x, temp := range data {
		fmt.Println("| ID 			:", temp.Id)
		fmt.Println("| Title 		:", temp.Title)
		fmt.Println("| Release 		:", temp.Release)
		fmt.Println("| Categoty 	:", temp.Category)
		fmt.Println("| Studio\t\t:", temp.Studio)
		fmt.Println("| Genre 		:", temp.Genre)
		fmt.Println("| Age Rate 	:", temp.Agerate)
		fmt.Println("| Rating\t\t:", temp.Rating)
		if x != length {
			fmt.Println("-----------------------")
		}
	}
}

func MainMenu() {
	var input int
	for input != 3 {

		fmt.Println("========================================")
		fmt.Println("|  _  _ ___ ___ _____ ___ _    _____  __")
		fmt.Println("| | \\| | __| __|_   _| __| |  |_ _\\ \\/ /")
		fmt.Println("| | .` | _|| _|  | | | _|| |__ | | >  < ")
		fmt.Println("| |_|\\_|___|___| |_| |_| |____|___/_/\\_\\")
		fmt.Println("|")
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
			ChooseMenu(input)
		} else if input != 3 {
			fmt.Println("|  INVALID INPUT! ")
		} else {
			fmt.Println("|  Exiting Program ")
		}
	}
}

func ChooseMenu(selector int) {
	scanner := bufio.NewScanner(os.Stdin)

	var input int
	masterData := "Movie"
	exitcase := 6

	for input != exitcase {
		fmt.Println("========================================")
		fmt.Printf("|\tPengolahan Data %s       \n", masterData)
		fmt.Println("========================================")
		fmt.Printf("| 1. Add %s                   \n", masterData)
		fmt.Printf("| 2. Delete %s                \n", masterData)
		fmt.Printf("| 3. Edit %s                  \n", masterData)
		fmt.Printf("| 4. View All %s              \n", masterData)
		fmt.Printf("| 5. View %s By      \n", masterData)
		fmt.Println("| 6. Back                      ")
		fmt.Println("========================================")
		fmt.Print("| Pilih Menu : ")
		fmt.Scan(&input)
		scanner.Scan()
		// clearScreen()

		switch input {
		case 1:
			fmt.Println("========================================")
			fmt.Println("|\t\tData Movie")
			var loop string
			for {
				if loop != "t" {
					InsertMovie()
					fmt.Print("| Ulang ? : ")
					fmt.Scan(&loop)
					scanner.Scan()
				} else {
					break
				}
			}
		case 2:
			fmt.Println("========================================")
			fmt.Println("|  Masukkan MovieID yang ingin dihapus ")
			fmt.Println("========================================")

			model.DeleteData(1)
		case 3:
			fmt.Println("========================================")
			fmt.Println("|  Masukkan Movie yang ingin diupdate ")
			fmt.Println("========================================")

			UpdateMovie()
		case 4:
			PrintAllMovie()
		case 5:
			ViewByMenu()
		}

		if input > exitcase {
			fmt.Println(" INVALID INPUT!")
		}
	}

}
