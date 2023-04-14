package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Age struct {
	date  int
	month int
	year  int
}

type User struct {
	username   string
	email      string
	membership string
	born       Age
}

type Movie struct {
	id       string
	title    string
	release  int
	genre    []string
	category int
	studio   string
	rating   float32
	agerate  string
}

type structUser struct {
	data User
	next *structUser
}

type structMovie struct {
	data Movie
	next *structMovie
}

var movieID int = 0
var user structUser
var movie structMovie

var genreList = []string{"Action", "Commedy", "Horror", "Romance", "Sci-fi", "Adventure", "Supranatural", "Magic", "Fantasy", "Thriler"}
var kategoriList = []string{"Movie", "Anime", "Series", "Cartoon", "Documenter"}

func Cls() {
	fmt.Print("\033[H\033[2J")
}

func InsertData(listUser *structUser, listMovie *structMovie, user User, movie Movie, selector int) {
	var tempUser structUser
	var tempMovie structMovie
	if selector == 1 {
		tempUser.data = user
		temp := listUser

		if temp.next == nil {
			temp.next = &tempUser
		} else {
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = &tempUser
		}
	} else {
		tempMovie.data = movie
		temp := listMovie

		if temp.next == nil {
			temp.next = &tempMovie
		} else {
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = &tempMovie
		}
	}
}

func DeleteData(listUser *structUser, listMovie *structMovie, selector int, params string) {
	tempUser := listUser
	tempMovie := listMovie

	var counter int = 0
	if selector == 1 {
		for tempUser.next != nil {
			counter++
			if tempUser.next.data.username == params {
				if counter == 1 {
					tempUser.next = tempUser.next.next
					break
				} else if tempUser.next.next != nil {
					tempUser.next = tempUser.next.next
					break
				} else {
					tempUser.next = nil
					break
				}
			}
			tempUser = tempUser.next
		}
	} else {
		for tempMovie.next != nil {
			counter++
			if tempMovie.next.data.id == params {
				if counter == 1 {
					tempMovie.next = tempMovie.next.next
					break
				} else if tempMovie.next.next != nil {
					tempMovie.next = tempMovie.next.next
					break
				} else {
					tempMovie.next = nil
					break
				}
			}
			tempMovie = tempMovie.next
		}
	}
}

func UpdateData(listUser *structUser, listMovie *structMovie, user User, movie Movie, selector int) {
	tempUser := listUser
	tempMovie := listMovie

	if selector == 1 {
		for tempUser.next != nil {
			if tempUser.next.data.username == user.username {
				tempUser.next.data = user
				break
			}
			tempUser = tempUser.next
		}
	} else {
		for tempMovie.next != nil {
			if tempMovie.next.data.id == movie.id {
				tempMovie.next.data = movie
				break
			}
			tempMovie = tempMovie.next
		}
	}
}

func ReadData(listUser *structUser, listMovie *structMovie, selector int, params string) {
	tempUser := listUser.next
	tempMovie := listMovie.next

	if tempUser == nil && tempMovie == nil {
		println("DATA MASIH KOSONG")
	} else {

		if params == "all" {
			// View All
			if selector == 1 {
				for tempUser != nil {
					fmt.Println("=================================")
					fmt.Println("|-Username       :", tempUser.data.username)
					fmt.Println("|-Email          :", tempUser.data.email)
					fmt.Println("|-Membership     :", tempUser.data.membership)
					fmt.Println("|-Tanggal Lahir  :", tempUser.data.born.date, "-", tempUser.data.born.month, "-", tempUser.data.born.year)
					tempUser = tempUser.next
				}
			} else {
				for tempMovie != nil {
					fmt.Println("=================================")
					fmt.Println("|-MovieID      :", tempMovie.data.id)
					fmt.Println("|-Judul        :", tempMovie.data.title)
					fmt.Println("|-Tahun Rilis  :", tempMovie.data.release)
					fmt.Print("|-Genre        : ")
					for i := 0; i < len(tempMovie.data.genre); i++ {
						fmt.Printf("%v ", tempMovie.data.genre[i])
					}
					fmt.Println("")
					fmt.Println("|-Kategori     : ", tempMovie.data.category)
					fmt.Println("|-Studio       : ", tempMovie.data.studio)
					fmt.Println("|-Rating       : ", tempMovie.data.rating)
					fmt.Println("|-Rate Umur    : ", tempMovie.data.agerate)
					tempMovie = tempMovie.next
				}
			}
		} else {
			// View By
			if selector == 1 {
				for tempUser != nil {
					if tempUser.data.username == params { // Param : username
						fmt.Println("=================================")
						fmt.Println("|-Username       : ", tempUser.data.username)
						fmt.Println("|-Email          : ", tempUser.data.email)
						fmt.Println("|-Membership     : ", tempUser.data.membership)
						fmt.Println("|-Tanggal Lahir  : ", tempUser.data.born.date, "-", tempUser.data.born.month, "-", tempUser.data.born.year)
					}
					tempUser = tempUser.next
				}
			} else {
				for tempMovie != nil { // tempMovie.data.category == params || tempMovie.data.genre == params ||
					if tempMovie.data.agerate == params {
						fmt.Println("=================================")
						fmt.Println("|-Movie ID     : ", tempMovie.data.id)
						fmt.Println("|-Judul        : ", tempMovie.data.title)
						fmt.Println("|-Tahun Rilis  : ", tempMovie.data.release)
						fmt.Println("|-Genre        : ", tempMovie.data.genre)
						fmt.Println("|-Kategori     : ", tempMovie.data.category)
						fmt.Println("|-Studio       : ", tempMovie.data.studio)
						fmt.Println("|-Rating       : ", tempMovie.data.rating)
						fmt.Println("|-Rate Umur    : ", tempMovie.data.agerate)
					}
					tempMovie = tempMovie.next
				}
			}
		}
	}
}

func InputUser(selector int, model int) {
	scanner := bufio.NewScanner(os.Stdin)
	var a, b, c, d, e, x, id string
	var release, date, month, year, sasa int
	var rating float32
	var genremov []string

	scanner.Scan()

	if model == 1 {

		for x != "t" {
			fmt.Print("MASUKKAN DATA ")
			if selector == 1 {

				fmt.Println("USER")
				fmt.Print("Username   : ")
				fmt.Scan(&a)

				scanner.Scan()

				fmt.Print("Email      : ")
				scanner.Scan()
				b = scanner.Text()
				fmt.Print("Membership : ")
				scanner.Scan()
				c = scanner.Text()

				fmt.Println("TANGGAL LAHIR")
				fmt.Print("Tanggal :")
				fmt.Scan(&date)
				fmt.Print("Bulan   :")
				fmt.Scan(&month)
				fmt.Print("Tahun   :")
				fmt.Scan(&year)
				scanner.Scan()
			} else {
				movieID++
				id = fmt.Sprintf("movie-%d", movieID)

				fmt.Println("MOVIE")
				fmt.Print("Judul       : ")
				scanner.Scan()
				a = scanner.Text()

				fmt.Print("Tahun Rilis : ")
				fmt.Scan(&release)

				for i := 0; i < len(genreList); i++ {
					fmt.Printf("%d. %v \n", i+1, genreList[i])
				}
				fmt.Printf("x. Exit")

				var genrex int
				fmt.Print("Genre       : ")
				for genrex != 99 {
					fmt.Scan(&genrex)
					if genrex != 99 {
						fmt.Print(",")
						genremov = append(genremov, genreList[genrex-1])
					}
				}

				for k := 1; k < len(kategoriList); k++ {
					fmt.Printf("%d. %v \n", k, kategoriList[k])
				}
				fmt.Printf("x. Exit")

				fmt.Print("Kategori    : ")
				fmt.Scan(&sasa)

				scanner.Scan()
				fmt.Print("Studio      : ")
				scanner.Scan()
				d = scanner.Text()

				fmt.Print("Rating      : ")
				fmt.Scan(&rating)
				// ASULAH
				scanner.Scan()
				fmt.Print("Rating Umur : ")
				scanner.Scan()
				e = scanner.Text()
			}

			// Data Skeleton
			dataUser := User{username: a, email: b, membership: c, born: Age{date, month, year}}
			dataMovie := Movie{id: id, title: a, release: release, genre: genremov, category: sasa, studio: d, rating: rating, agerate: e}

			// Insert Data
			InsertData(&user, &movie, dataUser, dataMovie, selector)

			// Break Looping
			// scanner.Scan() /* Reset fmt.Scan */
			fmt.Println("TAMBAH DATA ? [ y / t ]")
			scanner.Scan()
			x = strings.TrimSpace(scanner.Text())
		}
	} else {
		fmt.Print("EDIT DATA")
		if selector == 1 {
			fmt.Println("USER")
			fmt.Print("Username   : ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Email      : ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Membership : ")
			scanner.Scan()
			c = scanner.Text()

			fmt.Println("TANGGAL LAHIR")
			fmt.Print("Tanggal : ")
			fmt.Scan(&date)
			fmt.Print("Bulan   : ")
			fmt.Scan(&month)
			fmt.Print("Tahun   : ")
			fmt.Scan(&month)
			scanner.Scan()
		} else {
			fmt.Println("MOVIE")
			fmt.Print("Judul    : ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Tahun Rilis : ")
			scanner.Scan()
			fmt.Scan(&release)

			for i := 1; i < len(genreList); i++ {
				fmt.Printf("%d. %v", i, genreList[i])
			}

			fmt.Print("Genre    : ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Kategori : ")
			scanner.Scan()
			c = scanner.Text()
			fmt.Print("Studio   : ")
			scanner.Scan()
			d = scanner.Text()
			fmt.Print("Rating   : ")
			scanner.Scan()
			fmt.Scan(&rating)
			fmt.Print("Rating Umur : ")
			scanner.Scan()
			e = scanner.Text()
		}

		// Data Skeleton
		dataUser := User{username: a, email: b, membership: c, born: Age{date, month, year}}
		dataMovie := Movie{title: a, release: release, genre: genremov, category: sasa, studio: d, rating: rating, agerate: e}

		// Update Data
		UpdateData(&user, &movie, dataUser, dataMovie, selector)
	}
}

func mainMenu() {
	var input int
	for input != 3 {

		fmt.Println("==========================")
		fmt.Println("|    <> NEET FLIX <> \t |")
		fmt.Println("==========================")
		fmt.Println("| 1. Manage User \t |")
		fmt.Println("| 2. Manage Movie \t |")
		fmt.Println("| 3. Exit \t \t |")
		fmt.Println("==========================")
		fmt.Print("| Input : ")
		fmt.Scan(&input)
		fmt.Println("==========================")

		// Cls()

		if input < 3 && input > 0 {
			chooseMenu(input)
		} else if input != 3 {
			fmt.Println("INVALID INPUT!")
		} else {
			fmt.Println("Exiting Program")
		}
	}
}

func chooseMenu(selector int) {
	var input, exitcase int
	var masterData, params string

	if selector == 1 {
		masterData = "User"
		exitcase = 6
	} else {
		masterData = "Movie"
		exitcase = 8
	}

	for input != exitcase {

		fmt.Printf("===========================%v\n", strings.Repeat("=", len(masterData)))
		fmt.Printf("|    Pengolahan Data %s%v |\n", masterData, strings.Repeat(" ", len(masterData)))
		fmt.Printf("===========================%v\n", strings.Repeat("=", len(masterData)))
		fmt.Printf("| 1. Add %s%v             |\n", masterData, strings.Repeat(" ", len(masterData)))
		fmt.Printf("| 2. Delete %s%v          |\n", masterData, strings.Repeat(" ", len(masterData)))
		fmt.Printf("| 3. Edit %s%v            |\n", masterData, strings.Repeat(" ", len(masterData)))
		fmt.Printf("| 4. View All %s%v        |\n", masterData, strings.Repeat(" ", len(masterData)))

		if selector == 1 {
			fmt.Printf("| 5. View %s By Username%v|\n", masterData, strings.Repeat(" ", len(masterData)))
			fmt.Printf("| 6. Back                 %v|\n", strings.Repeat(" ", len(masterData)))
		} else {
			fmt.Printf("| 5. View %s By Category%v|\n", masterData, strings.Repeat(" ", len(masterData)))
			fmt.Printf("| 6. View %s By Genre   %v|\n", masterData, strings.Repeat(" ", len(masterData)))
			fmt.Printf("| 7. View %s By Age Rate%v|\n", masterData, strings.Repeat(" ", len(masterData)))
			fmt.Printf("| 8. Back                  %v|\n", strings.Repeat(" ", len(masterData)))
		}

		fmt.Printf("===========================%v\n", strings.Repeat("=", len(masterData)))

		// Select Menu
		fmt.Scan(&input)

		if selector == 1 {
			switch input {
			case 1:
				InputUser(selector, 1)
			case 2:
				fmt.Println("Masukkan username yang ingin dihapus")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 3:
				InputUser(1, 1)
			case 4:
				ReadData(&user, &movie, selector, "all")
			case 5:
				fmt.Println("Masukkan username yang ingin dilihat")
				fmt.Scan(&params)

				ReadData(&user, &movie, selector, params)
			}
		} else {
			switch input {
			case 1:
				InputUser(2, 1)
			case 2:
				fmt.Println("Masukkan MovieID yang ingin dihapus")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 3:
				InputUser(selector, 2)
			case 4:
				ReadData(&user, &movie, selector, "all")
			case 5:
				fmt.Println("Tampilkan movie berdasarkan kategori")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 6:
				fmt.Println("Tampilkan movie berdasarkan genre")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 7:
				fmt.Println("Tampilkan movie berdasarkan rate umur")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			}
		}

		if input > exitcase {
			fmt.Println("INVALID INPUT!")
		}
	}
}

func main() {
	mainMenu()
}
