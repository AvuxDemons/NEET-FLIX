package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Age struct {
	date 		 int
	year 		 int
	month 		 int
}

type User struct {
	username     string
	email        string
	membership   string
	born		 Age
}

type Movie struct {
	title      string
	release    int
	genre      string
	category   string
	studio     string
	rating     float32
	agerate    string
}

type structUser struct {
	data User
	next *structUser
}

type structMovie struct {
	data Movie
	next *structMovie
}

// var id int = 0
var user structUser
var movie structMovie

func Cls(){
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
			if tempMovie.next.data.title == params {
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
			if tempMovie.next.data.title == movie.title {
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

	if params == "nil" {
		// View All
		if selector == 1 {
			for tempUser != nil {
				fmt.Println("-Username  : ", tempUser.data.username)
				fmt.Println("-Email  : ", tempUser.data.email)
				fmt.Println("-Membership  : ", tempUser.data.membership)
				fmt.Println("-Tanggal Lahir  : ", tempUser.data.born)
				tempUser = tempUser.next
			}
		} else {
			for tempMovie != nil {
				fmt.Println("-Judul  : ", tempMovie.data.title)
				fmt.Println("-Tahun Rilis  : ", tempMovie.data.release)
				fmt.Println("-Genre  : ", tempMovie.data.genre)
				fmt.Println("-Kategori  : ", tempMovie.data.category)
				fmt.Println("-Studio  : ", tempMovie.data.studio)
				fmt.Println("-Rating  : ", tempMovie.data.rating)
				fmt.Println("-Rate Umur  : ", tempMovie.data.agerate)
				tempMovie = tempMovie.next
			}
		}
	} else {
		// View By
		if selector == 1 {
			for tempUser != nil {
				if tempUser.data.username == params {
					fmt.Println("-Username  : ", tempUser.data.username)
					fmt.Println("-Email  : ", tempUser.data.email)
					fmt.Println("-Membership  : ", tempUser.data.membership)
					fmt.Println("-Tanggal Lahir  : ", tempUser.data.born)
				}
				tempUser = tempUser.next
			}
		} else {
			for tempMovie != nil {
				if tempMovie.data.category == params || tempMovie.data.genre == params || tempMovie.data.agerate == params {
					fmt.Println("-Judul  : ", tempMovie.data.title)
					fmt.Println("-Tahun Rilis  : ", tempMovie.data.release)
					fmt.Println("-Genre  : ", tempMovie.data.genre)
					fmt.Println("-Kategori  : ", tempMovie.data.category)
					fmt.Println("-Studio  : ", tempMovie.data.studio)
					fmt.Println("-Rating  : ", tempMovie.data.rating)
					fmt.Println("-Rate Umur  : ", tempMovie.data.agerate)
				}
				tempMovie = tempMovie.next
			}
		}
	}
}

func InputUser(selector int, model int) {
	scanner := bufio.NewScanner(os.Stdin)
	var x string
	var a, b, c, d, e string //agar lebih hemat variabel, saya pake di input user dan movienya
	var release , date, month, year int
	var rating float32

	scanner.Scan()

	if model == 1 {

		for x != "t" {
			fmt.Print("Masukkan Data ")
			if selector == 1 {
				fmt.Println("User")
				fmt.Print("Username : ")
				scanner.Scan()
				a = scanner.Text()
				fmt.Print("Email : ")
				scanner.Scan()
				b = scanner.Text()
				fmt.Print("Membership : ")
				scanner.Scan()
				c = scanner.Text()

				fmt.Println("TITID LAHIR")
				fmt.Print("Tanggal :")
				fmt.Scan(&date)
				fmt.Print("Bulan :")
				fmt.Scan(&month)
				fmt.Print("Tahun :")
				fmt.Scan(&month)
			} else {
				fmt.Println("Movie")
				fmt.Print("Judul : ")
				scanner.Scan()
				a = scanner.Text()

				fmt.Print("Tahun Rilis : ")
				fmt.Scan(&release)

				scanner.Scan()
				fmt.Print("Genre : ")
				scanner.Scan()
				b = scanner.Text()
				fmt.Print("Kategori : ")
				scanner.Scan()
				c = scanner.Text()
				fmt.Print("Studio : ")
				scanner.Scan()
				d = scanner.Text()

				fmt.Print("Rating : ")
				fmt.Scan(&rating)

				scanner.Scan()
				fmt.Print("Rate Umur : ")
				scanner.Scan()
				e = scanner.Text()
			}

			// Data Skeleton
			dataUser := User{username: a, email: b, membership: c, born: Age{date,month,year}}
			dataMovie := Movie{title: a, release: release, genre: b, category: c, studio: d, rating: rating, agerate: e}

			// Insert Data
			InsertData(&user, &movie, dataUser, dataMovie, selector)

			// Break Looping
			scanner.Scan() /* Reset fmt.Scan */
			fmt.Println("TAMBAH DATA ? [ y / t ]")
			scanner.Scan()
			x = strings.TrimSpace(scanner.Text())
		}
	} else {
		fmt.Print("Masukkan Data Baru ")
		if selector == 1 {
			fmt.Println("User")
			fmt.Print("Username : ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Email : ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Membership : ")
			scanner.Scan()
			c = scanner.Text()

			fmt.Println("TITID LAHIR")
			fmt.Print("Tanggal : ")
			fmt.Scan(&date)
			fmt.Print("Bulan : ")
			fmt.Scan(&month)
			fmt.Print("Tahun : ")
			fmt.Scan(&month)
		} else {
			fmt.Println("Movie")
			fmt.Print("Judul : ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Tahun Rilis : ")
			scanner.Scan()
			fmt.Scan(&release)
			fmt.Print("Genre : ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Kategori : ")
			scanner.Scan()
			c = scanner.Text()
			fmt.Print("Studio : ")
			scanner.Scan()
			d = scanner.Text()
			fmt.Print("Rating : ")
			scanner.Scan()
			fmt.Scan(&rating)
			fmt.Print("Rate Umur : ")
			scanner.Scan()
			e = scanner.Text()
		}

		// Data Skeleton
		dataUser := User{username: a, email: b, membership: c, born: Age{date,month,year}}
		dataMovie := Movie{title: a, release: release, genre: b, category: c, studio: d, rating: rating, agerate: e}

		// Update Data
		UpdateData(&user, &movie, dataUser, dataMovie, selector)
	}
}

func mainMenu() {
	var input int
	for input != 3 {
		fmt.Println("----------------------")
		fmt.Println("|    N.E.E.T FLIX    |")
		fmt.Println("----------------------")
		fmt.Println("| 1. Manage User     |")
		fmt.Println("| 2. Manage Movie    |")
		fmt.Println("| 3. Exit            |")
		fmt.Println("----------------------")
		fmt.Print("| Input : ")
		fmt.Scan(&input)
		fmt.Println("----------------------")

		Cls()

		if input < 3 && input > 0 {
			chooseMenu(input)
		} else if input != 3 {
			fmt.Println("Invalid Input")
		} else {
			fmt.Println("Exiting Program")
		}
	}
}

func chooseMenu(selector int) {
	var input int
	var masterData, params string

	if selector == 1 {
		masterData = "User"
	} else {
		masterData = "Movie"
	}

	for input != 6 {

		fmt.Println("Pengolahan Data", masterData)

		fmt.Println("1. Add", masterData)
		fmt.Println("2. Delete", masterData)
		fmt.Println("3. Update", masterData)
		fmt.Println("4. View All", masterData)

		if selector == 1 {
			fmt.Println("5. View", masterData , "By Username")
			fmt.Println("6. Back")
		} else {
			fmt.Println("5. View", masterData , "By Category")
			fmt.Println("6. View", masterData , "By Genre")
			fmt.Println("7. View", masterData , "By Age Rate")
			fmt.Println("8. Back")
		}

		// Select Menu
		fmt.Scan(&input)

		if selector == 1 {
			switch input {
				case 1:
					InputUser(1, 1)
				case 2:
					fmt.Println("Masukkan NIP yang ingin dihapus")
					fmt.Scan(&params)
				
					DeleteData(&user, &movie, selector, params)
				case 3:
				
				case 4:
					ReadData(&user, &movie, 1, "nil")
				case 5:

			}
		} else {
			switch input {
				case 1:
					InputUser(2, 1)
				case 2:

				case 3:

				case 4:
					ReadData(&user, &movie, 2, "nil")
				case 5:

				case 6:

				case 7:

				case 8:

			}
		}
	}
}

func main() {
	mainMenu()
}