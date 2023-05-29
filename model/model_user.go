package model

import (
	"NeetFlix/database"
	"NeetFlix/entity"
	"fmt"
	"strconv"
	"strings"
)

// ██ ███    ██ ███████ ███████ ██████  ████████
// ██ ████   ██ ██      ██      ██   ██    ██
// ██ ██ ██  ██ ███████ █████   ██████     ██
// ██ ██  ██ ██      ██ ██      ██   ██    ██
// ██ ██   ████ ███████ ███████ ██   ██    ██

// ✅ Insert Data To Database
func InsertUser(user entity.User) {
	tempUser := entity.DatabaseUser{}
	tempUser.Data = user
	temp := &database.DBUser

	if temp.Next == nil {
		temp.Next = &tempUser
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &tempUser
	}
}

// ✅ Raw Data All Movie
func ViewAllUser() []entity.User {
	temp := database.DBUser.Next
	Users := []entity.User{}

	for temp != nil {
		Users = append(Users, temp.Data)
		temp = temp.Next
	}
	return Users
}

// ✅ Check Data
func CheckDataUser(params int) bool {
	temp := database.DBUser.Next
	if temp == nil {
		return false
	}

	for temp != nil {
		if temp.Data.Id == params {
			return true
		}
		temp = temp.Next
	}
	return false
}

func ViewUserBy(params string) []entity.User {
	temp := database.DBUser.Next
	users := []entity.User{}
	if temp == nil {
		return users
	}

	x, _ := strconv.Atoi(params)

	for temp != nil {
		if temp.Data.Id == x || strings.Contains(strings.ToLower(temp.Data.Username), strings.ToLower(params)) || strings.Contains(strings.ToLower(temp.Data.Email), strings.ToLower(params)) {
			users = append(users, temp.Data)
		}
		temp = temp.Next
	}
	return users
}

// ✅ Delete Data
func DeleteUser(params int) {
	temp := &database.DBUser
	// if temp == nil || temp.Next == nil {
	// 	fmt.Println("DATA MOVIE KOSONG")
	// 	return
	// }

	prev := temp
	current := temp.Next

	for current != nil {
		if current.Data.Id == params {
			prev.Next = current.Next
			fmt.Println("USER BERHASIL DIHAPUS")
			return
		}
		prev = current
		current = current.Next
	}
	fmt.Println("USER TIDAK DITEMUKAN")
}
