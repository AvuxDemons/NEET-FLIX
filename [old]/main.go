package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Age struct {
	date, month, year int
}

type User struct {
	username, email, membership string
	born                        Age
}

type Movie struct {
	id, title, category, studio, agerate string
	release                              int
	genre                                []string
	rating                               float64
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

var userMembership = []string{"Bronze", "Silver", "Gold", "Platinum", "Diamond"}
var genreList = []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Historical", "Horror", "Mistery", "Romance", "Sci-fi", "Sport", "Supranatural", "Thriller"}
var categoryList = []string{"Movie", "Anime", "Series", "Cartoon", "Documenter"}
var ageRate = []string{"KIDS | 6-11", "TEEN | 12-18", "ADULT | 18+"}

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

func reInput() {
	fmt.Print("\033[F")
	fmt.Print("\033[K")
}

func InsertMovie(listUser *structUser, listMovie *structMovie, user User, movie Movie, selector int) {
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
				if user.email != "-" {
					tempUser.next.data.email = user.email
				}
				if user.membership != "-" {
					tempUser.next.data.membership = user.membership
				}
				if user.born.date != 0 {
					tempUser.next.data.born.date = user.born.date
				}
				if user.born.month != 0 {
					tempUser.next.data.born.month = user.born.month
				}
				if user.born.year != 0 {
					tempUser.next.data.born.year = user.born.year
				}
				break
			}
			tempUser = tempUser.next
		}
	} else {
		for tempMovie.next != nil {
			if tempMovie.next.data.id == movie.id {
				if movie.title != "-" {
					tempMovie.next.data.title = movie.title
				}
				if movie.release != 0 {
					tempMovie.next.data.release = movie.release
				}
				if len(movie.genre) != 0 {
					tempMovie.next.data.genre = movie.genre
				}
				if movie.category != "-" {
					tempMovie.next.data.category = movie.category
				}
				if movie.studio != "-" {
					tempMovie.next.data.studio = movie.studio
				}
				if movie.rating != 0 {
					tempMovie.next.data.rating = movie.rating
				}
				if movie.agerate != "-" {
					tempMovie.next.data.agerate = movie.agerate
				}
				break
			}
			tempMovie = tempMovie.next
		}
	}
}

func PrintData(listUser *structUser, listMovie *structMovie, selector int) {
	tempUser := listUser
	tempMovie := listMovie

	if selector == 1 {
		fmt.Println("| Username       :", tempUser.data.username)
		fmt.Println("| Email          :", tempUser.data.email)
		fmt.Println("| Membership     :", tempUser.data.membership)
		fmt.Printf("| Tanggal Lahir  : %d/%d/%d \n", tempUser.data.born.date, tempUser.data.born.month, tempUser.data.born.year)
	} else {
		fmt.Println("| MovieID      :", tempMovie.data.id)
		fmt.Println("| Judul        :", tempMovie.data.title)
		fmt.Println("| Tahun Rilis  :", tempMovie.data.release)
		fmt.Printf("| Genre        : %v\n", strings.Join(tempMovie.data.genre, " "))
		fmt.Println("| Kategori     :", tempMovie.data.category)
		fmt.Println("| Studio       :", tempMovie.data.studio)
		fmt.Println("| Rating       :", tempMovie.data.rating)
		fmt.Println("| Rate Umur    :", tempMovie.data.agerate)
	}
}

func ReadData(listUser *structUser, listMovie *structMovie, selector int, params string) {
	tempUser := listUser.next
	tempMovie := listMovie.next

	if params == "all" {
		// View All
		if selector == 1 {
			if tempUser == nil {
				fmt.Println("========================================")
				println("|              DATA KOSONG              ")
				return
			}
			for tempUser != nil {
				fmt.Println("========================================")
				PrintData(tempUser, tempMovie, 1)
				tempUser = tempUser.next
			}
		} else {
			if tempMovie == nil {
				fmt.Println("========================================")
				println("|              DATA KOSONG              ")
				return
			}
			for tempMovie != nil {
				fmt.Println("========================================")
				PrintData(tempUser, tempMovie, 2)
				tempMovie = tempMovie.next
			}
		}
	} else {
		// View By
		if selector == 1 {
			if tempUser == nil {
				fmt.Println("========================================")
				println("|              DATA KOSONG              ")
				return
			}
			for tempUser != nil {
				if tempUser.data.username == params {
					fmt.Println("========================================")
					PrintData(tempUser, tempMovie, 1)
				}
				tempUser = tempUser.next
			}
		} else {
			if tempMovie == nil {
				fmt.Println("========================================")
				println("|              DATA KOSONG              ")
				return
			}
			for tempMovie != nil {
				if tempMovie.data.id == params || tempMovie.data.category == params || strings.Contains(strings.Join(tempMovie.data.genre, ","), params) || strings.Contains(strings.ToLower(tempMovie.data.title), strings.ToLower(params)) {
					fmt.Println("========================================")
					PrintData(tempUser, tempMovie, 2)
				}
				tempMovie = tempMovie.next
			}
		}
	}
}

func ValidateData(listUser *structUser, listMovie *structMovie, selector int, params string) bool {
	tempUser := listUser.next
	tempMovie := listMovie.next

	if selector == 1 {
		for tempUser != nil {
			if tempUser.data.username == params {
				return true
			}
			tempUser = tempUser.next
		}

	} else {
		for tempMovie != nil {
			if tempMovie.data.id == params {
				return true
			}
			tempMovie = tempMovie.next
		}
	}

	return false
}

func InputUser(selector int, model int, params string) {
	scanner := bufio.NewScanner(os.Stdin)

	var x string
	currentTime := time.Now()
	yearNow := currentTime.Year()

	// User
	var username, email, membership string
	var date, month, year int

	//Movie
	var id, title, category, studio, agerate string
	var release, cat, age, ship int
	var rating float64
	var genre []string

	masterData := "Movie"
	if selector == 1 {
		masterData = "User"
	}

	if model == 1 {
		fmt.Println("========================================")
		fmt.Println("|\t Masukkan Data", masterData)
		fmt.Println("========================================")
		for {
			if selector == 1 {
				scanner.Scan()
				for {
					fmt.Print("| Username   : ")
					scanner.Scan()
					username = scanner.Text()
					if strings.Contains(username, " ") {
						reInput()
						fmt.Println("| ! Username harus 1 kata !")
					} else {
						if ValidateData(&user, &movie, selector, username) {
							reInput()
							fmt.Println("| ! Username Telah Digunakan !")
						} else {
							break
						}
					}
				}

				fmt.Print("| Email      : ")
				fmt.Scan(&email)

				fmt.Println("| Membership : ")
				for k := 0; k < len(userMembership); k++ {
					fmt.Printf("|   %d. %v \n", k+1, userMembership[k])
				}

				fmt.Print("|  Pilih : ")
				fmt.Scan(&ship)
				reInput()
				fmt.Println("|   -", userMembership[ship-1], "Telah Dipilih")
				membership = userMembership[ship-1]

				fmt.Println("| Tanggal Lahir")
				for {
					fmt.Print("|   Tanggal : ")
					fmt.Scan(&date)

					if date > 31 || 1 > date {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						reInput()
						break
					}
				}

				for {
					fmt.Print("|   Bulan   : ")
					fmt.Scan(&month)
					if month > 12 || 1 > month {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						reInput()
						break
					}
				}

				for {
					fmt.Print("|   Tahun   : ")
					fmt.Scan(&year)
					if year > (yearNow-5) || (yearNow-90) > year {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						reInput()
						break
					}
				}

				reInput()
				fmt.Printf("| Tanggal Lahir   : %d/%d/%d \n", date, month, year)

				scanner.Scan()
			} else {
				genre = []string{}
				movieID++
				id = fmt.Sprintf("movie-%d", movieID)

				scanner.Scan()
				fmt.Print("| Judul       : ")
				scanner.Scan()
				title = scanner.Text()

				fmt.Print("| Tahun Rilis : ")
				fmt.Scan(&release)

				fmt.Println("| Genre	")
				for i := 0; i < len(genreList); i++ {
					fmt.Printf("|   %d. %v \n", i+1, genreList[i])
				}
				fmt.Printf("|   %d. Exit \n", len(genreList)+1)

				var genrex int
				for {
					fmt.Print("|  Pilih : ")
					fmt.Scan(&genrex)
					reInput()
					if genrex > len(genreList)+1 || genrex < 1 {
						fmt.Println("|   ! Invalid Input !")
					} else if genrex == len(genreList)+1 {
						if len(genre) != 0 {
							break
						} else {
							fmt.Println("|   ! Tambahkan Minimal 1 Genre !")
						}
					} else {
						if strings.Contains(strings.Join(genre, ""), genreList[genrex-1]) {
							fmt.Println("|   ! Genre Sudah Dipilih , Silahkan Tambahkan Genre Lain !")
						} else {
							fmt.Println("|   -", genreList[genrex-1], "Telah Ditambahkan")
							genre = append(genre, genreList[genrex-1])
						}
					}
				}

				fmt.Println("| Kategori")
				for k := 0; k < len(categoryList); k++ {
					fmt.Printf("|   %d. %v \n", k+1, categoryList[k])
				}
				for {
					fmt.Print("|  Pilih : ")
					fmt.Scan(&cat)
					reInput()
					if cat > len(categoryList) || cat < 1 {
						fmt.Println("|  ! Invalid Input !")
					} else {
						fmt.Println("|  -", categoryList[cat-1], "Telah Dipilih")
						category = categoryList[cat-1]
						break
					}
				}

				scanner.Scan()
				fmt.Print("| Studio      : ")
				scanner.Scan()
				studio = scanner.Text()

				for {
					fmt.Print("| Rating      : ")
					fmt.Scan(&rating)
					if rating > 100 || rating < 1 {
						reInput()
						fmt.Println("| ! Invalid Input !")
					} else {
						break
					}
				}

				fmt.Println("| Batasan Umur")
				for k := 0; k < len(ageRate); k++ {
					fmt.Printf("|   %d. %v \n", k+1, ageRate[k])
				}
				for {
					fmt.Print("|  Pilih : ")
					fmt.Scan(&age)
					reInput()
					if age > len(ageRate) || age < 1 {
						fmt.Println("|   ! Invalid Input !")
					} else {
						fmt.Println("|   -", ageRate[age-1], "Telah Dipilih")
						agerate = ageRate[age-1]
						break
					}
				}
			}

			// Data Skeleton
			dataUser := User{username: username, email: email, membership: membership, born: Age{date, month, year}}
			dataMovie := Movie{id: id, title: title, release: release, genre: genre, category: category, studio: studio, rating: rating, agerate: agerate}

			// Insert Data
			InsertMovie(&user, &movie, dataUser, dataMovie, selector)

			fmt.Println("|---------------------------------------")
			fmt.Print("| Menambah Data Lain ? (y/t) => ")
			fmt.Scan(&x)
			if x == "t" {
				clearScreen()
				break
			} else {
				fmt.Println("|---------------------------------------")
			}
		}
	} else {
		fmt.Println("|\t Masukkan Data Baru", masterData)
		fmt.Println("|		( Untuk tidak merubah data ) ", masterData)
		fmt.Println("========================================")
		if selector == 1 {

			fmt.Print("| Email      : ")
			fmt.Scan(&email)

			fmt.Println("| Membership : ")
			for k := 0; k < len(userMembership); k++ {
				fmt.Printf("|   %d. %v \n", k+1, userMembership[k])
			}
			fmt.Print("| Pilih : ")
			fmt.Scan(&ship)
			if ship != 0 {
				membership = userMembership[ship-1]
			} else {
				membership = "-"
			}

			fmt.Println("| Tanggal Lahir")
			for {
				fmt.Print("|   Tanggal : ")
				var dates string
				fmt.Scan(&dates)
				if dates != "-" {
					reInput()
					change, _ := strconv.Atoi(dates)
					if change > 31 || 1 > change {
						fmt.Println("|   ! Invalid Input !")
					} else {
						date = change
						break
					}
				} else {
					date = 0
					break
				}
			}

			for {
				fmt.Print("|   Bulan   : ")
				var months string
				fmt.Scan(&months)
				if months != "-" {
					change, _ := strconv.Atoi(months)
					if change > 12 || 1 > change {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						month = change
						reInput()
						break
					}
				} else {
					month = 0
					break
				}
			}

			for {
				fmt.Print("|   Tahun   : ")
				var years string
				fmt.Scan(&years)
				if years != "-" {
					reInput()
					change, _ := strconv.Atoi(years)
					if change > (yearNow-5) || (yearNow-90) > change {
						fmt.Println("|   ! Invalid Input !")
					} else {
						year = change
						break
					}
				} else {
					year = 0
					break
				}
			}

			reInput()
			fmt.Printf("| Tanggal Lahir   : %d/%d/%d \n", date, month, year)

			scanner.Scan()
		} else {
			scanner.Scan()
			fmt.Print("| Judul       : ")
			scanner.Scan()
			title = scanner.Text()

			fmt.Print("| Tahun Rilis : ")
			var releases string
			fmt.Scan(&releases)
			if releases != "-" {
				change, _ := strconv.Atoi(releases)
				release = change
			} else {
				release = 0
			}

			fmt.Println("| Genre	")
			for i := 0; i < len(genreList); i++ {
				fmt.Printf("|   %d. %v \n", i+1, genreList[i])
			}
			fmt.Printf("|   %d. Exit \n", len(genreList)+1)

			var genres string
			for {
				fmt.Print("|  Pilih : ")
				fmt.Scan(&genres)
				reInput()
				change, _ := strconv.Atoi(genres)
				if change == len(genreList)+1 || genres == "-" {
					break
				} else if change > len(genreList)+1 {
					fmt.Println("|   ! Invalid Input !")
				} else {
					if strings.Contains(strings.Join(genre, ""), genreList[change-1]) {
						fmt.Println("|   ! Genre Sudah Dipilih , Silahkan Tambahkan Genre Lain !")
					} else {
						fmt.Println("|   -", genreList[change-1], "Telah Ditambahkan")
						genre = append(genre, genreList[change-1])
					}
				}
			}

			fmt.Println("| Kategori")
			for k := 0; k < len(categoryList); k++ {
				fmt.Printf("|   %d. %v \n", k+1, categoryList[k])
			}
			for {
				fmt.Print("|  Pilih : ")
				var cats string
				fmt.Scan(&cats)
				if cats != "-" {
					reInput()
					change, _ := strconv.Atoi(cats)
					if change > len(categoryList) || change < 1 {
						fmt.Println("|   ! Invalid Input !")
					} else {
						fmt.Println("|   -", categoryList[change-1], "Telah Dipilih")
						category = categoryList[change-1]
						break
					}
				} else {
					category = "-"
					break
				}
			}

			scanner.Scan()
			fmt.Print("| Studio      : ")
			scanner.Scan()
			studio = scanner.Text()

			for {
				fmt.Print("| Rating      : ")
				var ratings string
				fmt.Scan(&ratings)
				if ratings != "-" {
					change, _ := strconv.ParseFloat(ratings, 64)
					if change > 100 || change < 1 {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						rating = change
						break
					}
				} else {
					rating = 0
					break
				}
			}

			fmt.Println("| Batasan Umur")
			for k := 0; k < len(ageRate); k++ {
				fmt.Printf("|   %d. %v \n", k+1, ageRate[k])
			}
			for {
				fmt.Print("|  Pilih : ")
				var ages string
				fmt.Scan(&ages)
				if ages != "-" {
					reInput()
					change, _ := strconv.Atoi(ages)
					if change > len(ageRate) || change < 1 {
						fmt.Println("|   ! Invalid Input !")
					} else {
						fmt.Println("|   -", ageRate[change-1], "Telah Dipilih")
						agerate = ageRate[change-1]
						break
					}
				} else {
					agerate = "-"
					break
				}
			}
		}

		// Data Skeleton
		dataUser := User{username: params, email: email, membership: membership, born: Age{date, month, year}}
		dataMovie := Movie{id: params, title: title, release: release, genre: genre, category: category, studio: studio, rating: rating, agerate: agerate}

		// Update Data
		UpdateData(&user, &movie, dataUser, dataMovie, selector)
	}
}

func mainMenu() {
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
			chooseMenu(input)
		} else if input != 3 {
			fmt.Println("|  INVALID INPUT! ")
		} else {
			fmt.Println("|  Exiting Program ")
		}
	}
}

func chooseMenu(selector int) {

	var input int
	var params string

	masterData := "Movie"
	exitcase := 8
	if selector == 1 {
		masterData, exitcase = "User", 6
	}

	for input != exitcase {

		fmt.Println("========================================")
		fmt.Printf("|\tPengolahan Data %s       \n", masterData)
		fmt.Println("========================================")
		fmt.Printf("| 1. Add %s                   \n", masterData)
		fmt.Printf("| 2. Delete %s                \n", masterData)
		fmt.Printf("| 3. Edit %s                  \n", masterData)
		fmt.Printf("| 4. View All %s              \n", masterData)
		if selector == 1 {
			fmt.Printf("| 5. View %s By Username      \n", masterData)
			fmt.Println("| 6. Back                     ")
		} else {
			fmt.Printf("| 5. View %s By Title      \n", masterData)
			fmt.Printf("| 6. View %s By Category         \n", masterData)
			fmt.Printf("| 7. View %s By Genre      \n", masterData)
			fmt.Println("| 8. Back                      ")
		}
		fmt.Println("========================================")
		fmt.Print("| Pilih Menu : ")
		fmt.Scan(&input)

		clearScreen()

		if selector == 1 {
			switch input {
			case 1:
				InputUser(selector, 1, "nil")
			case 2:
				fmt.Println("========================================")
				fmt.Println("|  Masukkan Username yang ingin dihapus ")
				fmt.Println("========================================")
				fmt.Print("| Masukkan Username : ")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 3:
				fmt.Println("========================================")
				fmt.Println("|  Masukkan Username yang ingin diupdate ")
				fmt.Println("========================================")
				fmt.Print("| Masukkan Username : ")
				fmt.Scan(&params)
				if ValidateData(&user, &movie, selector, params) {
					ReadData(&user, &movie, selector, params)
					fmt.Println("========================================")
					InputUser(selector, 2, params)
					fmt.Println("========================================")
					clearScreen()
				} else {
					reInput()
					fmt.Println("| User Tidak Ditemukan")
				}
			case 4:
				ReadData(&user, &movie, selector, "all")
			case 5:
				fmt.Println("========================================")
				fmt.Println("| Masukkan Username yang ingin dilihat ")
				fmt.Println("========================================")
				fmt.Print("| Masukkan Username : ")
				fmt.Scan(&params)

				ReadData(&user, &movie, selector, params)
			}
		} else {
			switch input {
			case 1:
				InputUser(selector, 1, "nil")
			case 2:
				fmt.Println("========================================")
				fmt.Println("|  Masukkan MovieID yang ingin dihapus ")
				fmt.Println("========================================")
				fmt.Scan(&params)

				DeleteData(&user, &movie, selector, params)
			case 3:
				fmt.Println("========================================")
				fmt.Println("|  Masukkan Movie yang ingin diupdate ")
				fmt.Println("========================================")
				fmt.Print("| Masukkan Movie ID : ")
				fmt.Scan(&params)
				if ValidateData(&user, &movie, selector, params) {
					ReadData(&user, &movie, selector, params)
					fmt.Println("========================================")
					InputUser(selector, 2, params)
					fmt.Println("========================================")
					clearScreen()
				} else {
					reInput()
					fmt.Println("|	Movie Tidak Ditemukan!")
				}
			case 4:
				ReadData(&user, &movie, selector, "all")
			case 5:
				fmt.Println("========================================")
				fmt.Println("|  Tampilkan movie berdasarkan judul ")
				fmt.Println("========================================")
				fmt.Printf("|Masukkan Judul : ")
				fmt.Scan(&params)

				ReadData(&user, &movie, selector, params)
			case 6:
				var x int
				fmt.Println("========================================")
				fmt.Println("|  Tampilkan movie berdasarkan kategori ")
				fmt.Println("========================================")
				for k := 0; k < len(categoryList); k++ {
					fmt.Printf("| %d. %v \n", k+1, categoryList[k])
				}

				fmt.Println("========================================")
				for {
					fmt.Print("| Kategori    : ")
					fmt.Scan(&x)
					if x < 1 || x > len(categoryList) {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						break
					}
				}
				params = categoryList[x-1]
				ReadData(&user, &movie, selector, params)
			case 7:
				var x int
				fmt.Println("========================================")
				fmt.Println("|  Tampilkan movie berdasarkan genre ")
				fmt.Println("========================================")
				for k := 0; k < len(genreList); k++ {
					fmt.Printf("| %d. %v \n", k+1, genreList[k])
				}

				fmt.Println("========================================")
				for {
					fmt.Print("| Genre    : ")
					fmt.Scan(&x)
					if x < 1 || x > len(genreList) {
						reInput()
						fmt.Println("|   ! Invalid Input !")
					} else {
						break
					}
				}
				params = genreList[x-1]
				ReadData(&user, &movie, selector, params)
			}
		}

		if input > exitcase {
			fmt.Println(" INVALID INPUT!")
		}
	}
}

func main() {
	mainMenu()
}
