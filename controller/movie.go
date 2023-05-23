package controller

import (
	"NeetFlix/entity"
	"NeetFlix/model"
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

// ✅ Insert Data
func InsertMovie(Id int, Title string, Release int, Category []string, Studio string, Genre []string, AgeRate string, Rating float32) {
	data := entity.Movie{
		Id:       Id,
		Title:    Title,
		Release:  Release,
		Category: Category,
		Studio:   Studio,
		Genre:    Genre,
		Agerate:  AgeRate,
		Rating:   Rating,
	}
	model.InsertMovie(data)
}

func UpdateMovie(Id int, Title string, Release int, Category []string, Studio string, Genre []string, AgeRate string, Rating float32) {
	data := entity.Movie{
		Id:       Id,
		Title:    Title,
		Release:  Release,
		Category: Category,
		Studio:   Studio,
		Genre:    Genre,
		Agerate:  AgeRate,
		Rating:   Rating,
	}
	model.InsertMovie(data)
}

// func ViewBy(params string) entity.Movie {
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
