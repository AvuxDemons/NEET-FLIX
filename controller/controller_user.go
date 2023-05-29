package controller

import (
	"NeetFlix/database"
	"NeetFlix/entity"
	"NeetFlix/model"
	"fmt"
)

// func ValidateMovie(params int) bool {
// 	temp := &database.DBMovie
// 	for temp != nil {
// 		if temp.Data.Id == params {
// 			return true
// 		}
// 		temp = temp.Next
// 	}
// 	return false
// }

func IsUserEmpty() bool {
	temp := &database.DBUser
	if temp.Next == nil {
		fmt.Println("========================================")
		fmt.Println("| DATA KOSONG")
		return true
	}
	return false
}

// ██ ███    ██ ███████ ███████ ██████  ████████
// ██ ████   ██ ██      ██      ██   ██    ██
// ██ ██ ██  ██ ███████ █████   ██████     ██
// ██ ██  ██ ██      ██ ██      ██   ██    ██
// ██ ██   ████ ███████ ███████ ██   ██    ██

// ✅ Insert Data
func InsertUser(Id int, Username string, Email string, Membership string, Date int, Month int, Year int) {
	calendar := entity.Age{
		Date:  Date,
		Month: Month,
		Year:  Year,
	}

	data := entity.User{
		Id:         Id,
		Username:   Username,
		Email:      Email,
		Membership: Membership,
		Born:       calendar,
	}
	model.InsertUser(data)
}

// ██    ██ ██████  ██████   █████  ████████ ███████
// ██    ██ ██   ██ ██   ██ ██   ██    ██    ██
// ██    ██ ██████  ██   ██ ███████    ██    █████
// ██    ██ ██      ██   ██ ██   ██    ██    ██
//	██████  ██      ██████  ██   ██    ██    ███████

func UpdateUser(Id int, Username string, Email string, Membership string, Date int, Month int, Year int) {
	temp := &database.DBUser
	for temp != nil {
		if temp.Data.Id == Id {
			if Username != "-" {
				temp.Next.Data.Username = Username
			}
			if Email != "-" {
				temp.Next.Data.Email = Email
			}
			if Membership != "-" {
				temp.Next.Data.Membership = Membership
			}
			if Date != 0 {
				temp.Next.Data.Born.Date = Date
			}
			if Month != 0 {
				temp.Next.Data.Born.Month = Month
			}
			if Year != 0 {
				temp.Next.Data.Born.Year = Year
			}
		}
		temp = temp.Next
	}
}

// func ViewMovieBy(params string) entity.Movie {
// 	temp := &database.DBMovie
// 	if temp == nil {
// 		return entity.Movie{}
// 	}
// 	for temp != nil {
// 		if strings.Contains(strings.Join(temp.Data.Genre, ","), params) || strings.Contains(strings.ToLower(temp.Data.Title), strings.ToLower(params)) {
// 			return temp.Data
// 		}
// 		temp = temp.Next
// 	}
// 	return entity.Movie{}
// }
