package view

import (
	"NeetFlix/controller"
	"NeetFlix/database"
	"NeetFlix/entity"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

var GenreList = []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Historical", "Horror", "Mistery", "Romance", "Sci-fi", "Sport", "Supranatural", "Thriller"}
var CategoryList = []string{"Movie", "Anime", "Series", "Cartoon", "Documenter"}
var AgeRate = []string{"KIDS | 6-11", "TEEN | 12-18", "ADULT | 18+"}

// ✅ Fungsi Input Data Movie
func InsertMovie() {
	scanner := bufio.NewScanner(os.Stdin)

	var Release, Id int
	var Title, Studio, Agerate string
	var Category, Genre []string
	var Rating float32

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
	for i, genre := range GenreList {
		fmt.Printf("| %d. %v \n", i+1, genre)
	}
	fmt.Printf("| %d. Exit \n", len(GenreList)+1)
	fmt.Println("|--------------------------------")
	var Gen int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Gen)
		reInput()

		if Gen < 1 || Gen > len(GenreList)+1 {
			fmt.Println("|   ! Invalid Input !")
		} else if Gen == len(GenreList)+1 {
			if len(Genre) == 0 {
				fmt.Println("|   ! Tambahkan Minimal 1 Genre !")
			} else {
				break
			}
		} else {
			selectedGenre := GenreList[Gen-1]
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
	for k, category := range CategoryList {
		fmt.Printf("| %d. %v \n", k+1, category)
	}
	fmt.Printf("| %d. Exit \n", len(CategoryList)+1)
	fmt.Println("|--------------------------------")
	var Cat int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Cat)
		reInput()

		if Cat < 1 || Cat > len(CategoryList)+1 {
			fmt.Println("| ! Invalid Input !")
		} else if Cat == len(CategoryList)+1 {
			if len(Category) == 0 {
				fmt.Println("| ! Tambahkan Minimal 1 Category !")
			} else {
				break
			}
		} else {
			selectedCategory := CategoryList[Cat-1]
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
	for i := range AgeRate {
		fmt.Printf("| %d. %v \n", i+1, AgeRate[i])
	}
	fmt.Println("|--------------------------------")
	var Age int
	for {
		fmt.Print("| Pilih : ")
		fmt.Scan(&Age)
		reInput()
		if Age > len(AgeRate) || Age < 1 {
			fmt.Println("| ! Invalid Input !")
		} else {
			fmt.Println("| - ", AgeRate[Age-1], "Telah Dipilih")
			Agerate = AgeRate[Age-1]
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
	controller.InsertData(Id, Title, Release, Category, Studio, Genre, Agerate, Rating)
}

// ❌
func printMovie(temp entity.Movie) {
	fmt.Println("Title 		:", temp.Title)
	fmt.Println("Release 	:", temp.Release)
	fmt.Println("Categoty	:", temp.Category)
	fmt.Println("Studio 	:", temp.Studio)
	fmt.Println("Genre 		:", temp.Genre)
	fmt.Println("Age Rate 	:", temp.Agerate)
	fmt.Println("Rating 	:", temp.Rating)
}

// ✅
func PrintAllMovie() {
	temp := database.DBMovie.Next
	if temp == nil {
		fmt.Println("DATA KOSONG")
	} else {
		for temp != nil {
			fmt.Println("ID 		:", temp.Data.Id)
			fmt.Println("Title 		:", temp.Data.Title)
			fmt.Println("Release 	:", temp.Data.Release)
			fmt.Println("Categoty 	:", temp.Data.Category)
			fmt.Println("Studio\t\t:", temp.Data.Studio)
			fmt.Println("Genre 		:", temp.Data.Genre)
			fmt.Println("Age Rate 	:", temp.Data.Agerate)
			fmt.Println("Rating\t\t:", temp.Data.Rating)
			fmt.Println("-----------------------")
			temp = temp.Next
		}
	}
}
