package view

import (
	"NeetFlix/controller"
	"NeetFlix/database"
	"NeetFlix/model"
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"strconv"
	"strings"
	"time"
)

// ✅ Fungsi Input Data Movie
func InsertUser() {
	scanner := bufio.NewScanner(os.Stdin)

	currentTime := time.Now()
	yearNow := currentTime.Year()

	var Username, Email, Membership string
	var Id, Date, Month, Year int

	// Id
	database.IDUser += 1
	Id = database.IDUser

	for {
		fmt.Print("| Username   : ")
		scanner.Scan()
		Username = scanner.Text()
		if strings.Contains(Username, " ") {
			reInput()
			fmt.Println("| ! Username harus minimal 1 kata !")
		} else {
			// MATIKAN VALIDASI SEMENTARA
			// if ValidateData(&user, &movie, selector, username) {
			// 	reInput()
			// 	fmt.Println("| ! Username Telah Digunakan !")
			// } else {
			break
			// }
		}
	}

	fmt.Print("| Email      : ")
	fmt.Scan(&Email)

	fmt.Println("| Membership : ")
	for k := 0; k < len(database.UserMembership); k++ {
		fmt.Printf("|   %d. %v \n", k+1, database.UserMembership[k])
	}

	var x int
	fmt.Print("|  Pilih : ")
	fmt.Scan(&x)
	reInput()
	fmt.Println("|   -", database.UserMembership[x-1], "Telah Dipilih")
	Membership = database.UserMembership[x-1]

	fmt.Println("| Tanggal Lahir")
	for {
		fmt.Print("|   Tanggal : ")
		fmt.Scan(&Date)
		reInput()
		if Date > 31 || 1 > Date {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	for {
		fmt.Print("|   Bulan   : ")
		fmt.Scan(&Month)
		reInput()
		if Month > 12 || 1 > Month {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	for {
		fmt.Print("|   Tahun   : ")
		fmt.Scan(&Year)
		reInput()
		if Year > (yearNow-5) || (yearNow-90) > Year {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	reInput()
	fmt.Printf("| Tanggal Lahir   : %d/%d/%d \n", Date, Month, Year)

	controller.InsertUser(Id, Username, Email, Membership, Date, Month, Year)
}

func UpdateUser3() {
	scanner := bufio.NewScanner(os.Stdin)

	currentTime := time.Now()
	yearNow := currentTime.Year()

	var Username, Email, Membership string
	var Id, Date, Month, Year int

	// Id
	fmt.Print("| Masukkan ID : ")
	fmt.Scan(&Id)
	scanner.Scan()

	for {
		fmt.Print("| Username   : ")
		scanner.Scan()
		Username = scanner.Text()
		if strings.Contains(Username, " ") {
			reInput()
			fmt.Println("| ! Username harus minimal 1 kata !")
		} else {
			// MATIKAN VALIDASI SEMENTARA
			// if ValidateData(&user, &movie, selector, username) {
			// 	reInput()
			// 	fmt.Println("| ! Username Telah Digunakan !")
			// } else {
			break
			// }
		}
	}

	fmt.Print("| Email      : ")
	fmt.Scan(&Email)

	fmt.Println("| Membership : ")
	for k := 0; k < len(database.UserMembership); k++ {
		fmt.Printf("|   %d. %v \n", k+1, database.UserMembership[k])
	}

	var x int
	fmt.Print("|  Pilih : ")
	fmt.Scan(&x)
	reInput()
	fmt.Println("|   -", database.UserMembership[x-1], "Telah Dipilih")
	Membership = database.UserMembership[x-1]

	fmt.Println("| Tanggal Lahir")
	for {
		fmt.Print("|   Tanggal : ")
		fmt.Scan(&Date)
		reInput()
		if Date > 31 || 1 > Date {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	for {
		fmt.Print("|   Bulan   : ")
		fmt.Scan(&Month)
		reInput()
		if Month > 12 || 1 > Month {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	for {
		fmt.Print("|   Tahun   : ")
		fmt.Scan(&Year)
		reInput()
		if Year > (yearNow-5) || (yearNow-90) > Year {
			fmt.Println("|   ! Invalid Input !")
		} else {
			break
		}
	}

	reInput()
	fmt.Printf("| Tanggal Lahir   : %d/%d/%d \n", Date, Month, Year)
	fmt.Println("========================================")

	controller.UpdateUser(Id, Username, Email, Membership, Date, Month, Year)
}

func MenuViewUserBy() {
	fmt.Println("| PILIH BERDASARKAN")
	fmt.Println("| 1. USERNAME ")
	fmt.Println("| 2. EMAIL ")

	var x int
	fmt.Println("========================================")
	fmt.Print("| PILIH : ")
	fmt.Scan(&x)
	ViewUserBy(x)
}

func ViewUserBy(selector int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("========================================")
	fmt.Print("| Search Query :  ")
	scanner.Scan()
	scanner.Scan()
	Params := scanner.Text()

	PrintUser(Params)
}

func PrintUser(params string) {
	data := model.ViewUserBy(params)
	for _, temp := range data {
		fmt.Println("| ID\t\t:", temp.Id)
		fmt.Println("| Username\t:", temp.Username)
		fmt.Println("| Email\t\t:", temp.Email)
		fmt.Println("| Membership\t:", temp.Membership)
		fmt.Println("| Tanggal\t:", temp.Born.Date)
		fmt.Println("| Bulan\t\t:", temp.Born.Month)
		fmt.Println("| Tahun\t\t:", temp.Born.Year)
	}
}

func PrintAllUser() {
	data := model.ViewAllUser()
	length := len(data)
	for x, temp := range data {
		fmt.Println("| ID           :", temp.Id)
		fmt.Println("| Username     :", temp.Username)
		fmt.Println("| Email        :", temp.Email)
		fmt.Println("| Membership   :", temp.Membership)
		fmt.Println("| Tanggal      :", temp.Born.Date)
		fmt.Println("| Bulan        :", temp.Born.Month)
		fmt.Println("| Tahun        :", temp.Born.Year)
		if x+1 != length {
			fmt.Println("-----------------------")
		}
	}
}

// ERROR DI BAGIAN PEMILIHAN ID
func UpdateUser() {
	scanner := bufio.NewScanner(os.Stdin)

	currentTime := time.Now()
	yearNow := currentTime.Year()

	var Temp int

	fmt.Println("========================================")
	fmt.Print("| Pilih ID : ")
	fmt.Scan(&Temp)
	scanner.Scan()

	var Username, Email, Membership string
	var Id, Date, Month, Year int

	if model.CheckDataUser(Temp) {

		Params := strconv.Itoa(Temp)
		PrintUser(Params)

		// EDIT DATA

		// Username
		fmt.Println("=================================")
		fmt.Print("| Username       : ")
		scanner.Scan()
		Username = scanner.Text()
		if Username == "-" {
			fmt.Println("| Data Tidak Berubah")
		}

		// Email
		fmt.Println("=================================")
		fmt.Print("| Email : ")
		fmt.Scan(&Email)
		if Email != "-" {
			fmt.Println("| Data Tidak Berubah")
		}

		// Membership
		// fmt.Println("=================================")
		// fmt.Print("| Membership      : ")
		// scanner.Scan()
		// Membership = scanner.Text()
		// if Membership == "-" {
		// 	fmt.Println("| Data Tidak Berubah")
		// }

		// Membership
		fmt.Println("=================================")
		fmt.Println("| Membership : ")
		fmt.Println("|--------------------------------")
		for k, Membership := range database.UserMembership {
			fmt.Printf("| %d. %v \n", k+1, Membership)
		}
		fmt.Println("|--------------------------------")
		for {
			var member string

			fmt.Print("| Pilih : ")
			fmt.Scan(&member)

			reInput()

			if member != "-" {
				x, _ := strconv.Atoi(member)
				if x < 1 || x > len(database.UserMembership) {
					fmt.Println("| ! Invalid Input !")
					continue
				} else {
					Membership = database.UserMembership[x-1]
					break
				}
			} else {
				Membership = "-"
				fmt.Println("| Data Tidak Berubah")
				break
			}
		}

		fmt.Println("| Tanggal Lahir")
		for {
			fmt.Print("|   Tanggal : ")
			var dates string
			fmt.Scan(&dates)
			if dates != "-" {
				reInput()
				x, _ := strconv.Atoi(dates)
				if x > 31 || 1 > x {
					fmt.Println("|   ! Invalid Input !")
				} else {
					Date = x
					break
				}
			} else {
				Date = 0
				break
			}
		}

		for {
			fmt.Print("|   Bulan   : ")
			var months string
			fmt.Scan(&months)
			if months != "-" {
				x, _ := strconv.Atoi(months)
				if x > 12 || 1 > x {
					reInput()
					fmt.Println("|   ! Invalid Input !")
				} else {
					Month = x
					reInput()
					break
				}
			} else {
				Month = 0
				break
			}
		}

		for {
			fmt.Print("|   Tahun   : ")
			var years string
			fmt.Scan(&years)
			if years != "-" {
				reInput()
				x, _ := strconv.Atoi(years)
				if x > (yearNow-5) || (yearNow-90) > x {
					fmt.Println("|   ! Invalid Input !")
				} else {
					Year = x
					break
				}
			} else {
				Year = 0
				break
			}
		}

		reInput()
		fmt.Printf("| Tanggal Lahir   : %d/%d/%d \n", Date, Month, Year)

		controller.UpdateUser(Id, Username, Email, Membership, Date, Month, Year)

	} else {
		fmt.Println("| Data Masih Kosong")
	}
}

func MenuUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var input int
	masterData := "User"
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
			fmt.Println("|\t\tData User")
			fmt.Println("========================================")
			var loop string
			for {
				if loop != "t" {
					InsertUser()
					fmt.Print("| Ulang ? : ")
					fmt.Scan(&loop)
					scanner.Scan()
				} else {
					break
				}
			}
		case 2:
			var del int
			if !controller.IsUserEmpty() {
				fmt.Println("========================================")
				fmt.Println("|  Masukkan UserID yang ingin dihapus ")
				fmt.Print("| Pilih ID : ")
				fmt.Scan(&del)
				fmt.Println("========================================")
				model.DeleteUser(del)
			}
		case 3:
			// var temp entity.User
			if !controller.IsUserEmpty() {
				fmt.Println("========================================")
				fmt.Println("|  Masukkan User yang ingin diupdate ")
				UpdateUser()
				fmt.Println("========================================")
			}
		case 4:
			if !controller.IsUserEmpty() {
				PrintAllUser()
			}
		case 5:
			if !controller.IsUserEmpty() {
				MenuViewUserBy()
			}
		}
		if input > exitcase {
			fmt.Println(" INVALID INPUT!")
		}
	}

}
