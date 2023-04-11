package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	username    string
	email       string
	membership  string
	tanggalLahir int
}

type Movie struct {
	judul      string
	tahunRilis int
	genre      string
	kategori   string
	studio     string
	rating     float32
	rateUmur   string
}

type structUser struct {
	data User
	next *structUser
}

type structMovie struct {
	data Movie
	next *structMovie
}

var id int = 0
var user structUser
var movie structMovie

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

func DeleteData(listUser *structUser, listMovie *structMovie, selector int, userId string) {
	tempUser := listUser
	tempMovie := listMovie

	var counter int = 0
	if selector == 1 {
		for tempUser.next != nil {
			counter++
			if tempUser.next.data.username == userId {
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
			if tempMovie.next.data.judul == userId {
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
			if tempMovie.next.data.judul == movie.judul {
				tempMovie.next.data = movie
				break
			}
			tempMovie = tempMovie.next
		}
	}
}

func ReadData(listUser *structUser, listMovie *structMovie, selector int, sort string) {
	tempUser := listUser.next
	tempMovie := listMovie.next

	fmt.Print("View Data ")
	if selector == 1 {
		fmt.Println("Dosen")
	} else {
		fmt.Println("Mahasiswa")
	}

	if sort == "nil" {
		// View All
		if selector == 1 {
			for tempUser != nil {
				fmt.Println("-Username\t : ", tempUser.data.username)
				fmt.Println("-Email\t : ", tempUser.data.email)
				fmt.Println("-Membership\t : ", tempUser.data.membership)
				fmt.Println("-Tanggal Lahir\t : ", tempUser.data.tanggalLahir)
				tempUser = tempUser.next
			}
		} else {
			for tempMovie != nil {
				fmt.Println("-Judul\t : ", tempMovie.data.judul)
				fmt.Println("-Tahun Rilis\t : ", tempMovie.data.tahunRilis)
				fmt.Println("-Genre\t : ", tempMovie.data.genre)
				fmt.Println("-Kategori\t : ", tempMovie.data.kategori)
				fmt.Println("-Studio\t : ", tempMovie.data.studio)
				fmt.Println("-Rating\t : ", tempMovie.data.rating)
				fmt.Println("-Rate Umur\t : ", tempMovie.data.rateUmur)
				tempMovie = tempMovie.next
			}
		}
	} else {
		// View By Username or Judul
		if selector == 1 {
			for tempUser != nil {
				if tempUser.data.username == sort {
					fmt.Println("-Username\t : ", tempUser.data.username)
					fmt.Println("-Email\t : ", tempUser.data.email)
					fmt.Println("-Membership\t : ", tempUser.data.membership)
					fmt.Println("-Tanggal Lahir\t : ", tempUser.data.tanggalLahir)
				}
				tempUser = tempUser.next
			}
		} else {
			for tempMovie != nil {
				if tempMovie.data.judul == sort {
					fmt.Println("-Judul\t : ", tempMovie.data.judul)
					fmt.Println("-Tahun Rilis\t : ", tempMovie.data.tahunRilis)
					fmt.Println("-Genre\t : ", tempMovie.data.genre)
					fmt.Println("-Kategori\t : ", tempMovie.data.kategori)
					fmt.Println("-Studio\t : ", tempMovie.data.studio)
					fmt.Println("-Rating\t : ", tempMovie.data.rating)
					fmt.Println("-Rate Umur\t : ", tempMovie.data.rateUmur)
				}
				tempMovie = tempMovie.next
			}
		}
	}
}

func inputUser(selector int, model int) {
	scanner := bufio.NewScanner(os.Stdin)
	var x string
	var a, b, c, d, e string //agar lebih hemat variabel, saya pake di input user dan movienya
	var tanggalLahir, tahunRilis int
	var rating float32

	scanner.Scan()

	if model == 1 {

		for x != "t" {
			fmt.Print("Masukkan Data ")
			if selector == 1 {
				fmt.Println("User")
				fmt.Print("Username\t: ")
				scanner.Scan()
				a = scanner.Text()
				fmt.Print("Email\t: ")
				scanner.Scan()
				b = scanner.Text()
				fmt.Print("Membership\t: ")
				scanner.Scan()
				c = scanner.Text()

				fmt.Print("Tanggal Lahir\t: ")
				fmt.Scan(&tanggalLahir)
			} else {
				fmt.Println("Movie")
				fmt.Print("Judul\t: ")
				scanner.Scan()
				a = scanner.Text()
				
				fmt.Print("Tahun Rilis\t: ")
				fmt.Scan(&tahunRilis)

				scanner.Scan()
				fmt.Print("Genre\t: ")
				scanner.Scan()
				b = scanner.Text()
				fmt.Print("Kategori\t: ")
				scanner.Scan()
				c = scanner.Text()
				fmt.Print("Studio\t: ")
				scanner.Scan()
				d = scanner.Text()
				
				fmt.Print("Rating\t: ")
				fmt.Scan(&rating)
				
				scanner.Scan()
				fmt.Print("Rate Umur\t: ")
				scanner.Scan()
				e = scanner.Text()
			}

			// Data Skeleton
			dataUser := User{username: a, email: b, membership: c, tanggalLahir: tanggalLahir}
			dataMovie := Movie{judul: a, tahunRilis: tahunRilis, genre: b, kategori: c, studio: d, rating: rating, rateUmur: e}
			
			// Insert Data
			InsertData(&user, &movie, dataUser, dataMovie, selector)
			
			// Break Looping
			scanner.Scan() // Reset fmt.Scan

			fmt.Println("TAMBAH DATA ? [ y / t ]")
			scanner.Scan()
			x = strings.TrimSpace(scanner.Text())
		}
	} else {
		fmt.Print("Masukkan Data Baru ")
		if selector == 1 {
			fmt.Println("User")
			fmt.Print("Username\t: ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Email\t: ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Membership\t: ")
			scanner.Scan()
			c = scanner.Text()
			fmt.Print("Tanggal Lahir\t: ")
			scanner.Scan()
			fmt.Scan(&tanggalLahir)
		} else {
			fmt.Println("Movie")
			fmt.Print("Judul\t: ")
			scanner.Scan()
			a = scanner.Text()
			fmt.Print("Tahun Rilis\t: ")
			scanner.Scan()
			fmt.Scan(&tahunRilis)
			fmt.Print("Genre\t: ")
			scanner.Scan()
			b = scanner.Text()
			fmt.Print("Kategori\t: ")
			scanner.Scan()
			c = scanner.Text()
			fmt.Print("Studio\t: ")
			scanner.Scan()
			d = scanner.Text()
			fmt.Print("Rating\t: ")
			scanner.Scan()
			fmt.Scan(&rating)
			fmt.Print("Rate Umur\t: ")
			scanner.Scan()
			e = scanner.Text()
		}

		// Data Skeleton
		dataUser := User{username: a, email: b, membership: c, tanggalLahir: tanggalLahir}
		dataMovie := Movie{judul: a, tahunRilis: tahunRilis, genre: b, kategori: c, studio: d, rating: rating, rateUmur: e}

		// Update Data
		UpdateData(&user, &movie, dataUser, dataMovie, selector)
	}
}

func mainMenu() {
	var input int
	for input != 3 {
		fmt.Println("N.E.E.T - FLIX")
		fmt.Println("1. USER")
		fmt.Println("2. MOVIE")
		fmt.Println("3. EXIT")
		fmt.Scan(&input)
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
	// var query string
	var pilih string

	if selector == 1 {
		pilih = "User"
	} else {
		pilih = "Movie"
	}

	for input != 6 {
		fmt.Print("PENGOLAHAN DATA ")
		if selector == 1 {
			fmt.Println("USER")
		} else {
			fmt.Println("MOVIE")
		}

		fmt.Println("1. Tambah", pilih)
		fmt.Println("2. Hapus", pilih)
		fmt.Println("3. Update", pilih)
		fmt.Println("4. Lihat Semua", pilih)
		fmt.Println("5. Lihat KONTOL", pilih)
		fmt.Println("6. Kembali")

		// Select Menu
		fmt.Scan(&input)

		switch input {
		case 1:
			if selector == 1 {
				inputUser(1, 1)
			} else {
				inputUser(2, 1)
			}
		case 2:
			
		case 3:
			
		case 4:
			
		case 5:
			
		}
	}
}

func main() {
	mainMenu()
}