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
func InsertData(Id int, Title string, Release int, Category []string, Studio string, Genre []string, AgeRate string, Rating float32) {
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
	model.InsertData(data)
}

// func UpdateMovie(movie entity.Movie) {
// 	temp := &database.DBMovie
// 	for temp.Next != nil {
// 		if temp.Next.Data.Id == movie.Id {
// 			if movie.Title != "-" {
// 				temp.Next.Data.Title = movie.Title
// 			}
// 			if movie.Release != 0 {
// 				temp.Next.Data.Release = movie.Release
// 			}
// 			if len(movie.Genre) != 0 {
// 				temp.Next.Data.Genre = movie.Genre
// 			}
// 			// if movie.Category != "-" {
// 			// 	temp.Next.Data.Category = movie.Category
// 			// }
// 			if movie.Studio != "-" {
// 				temp.Next.Data.Studio = movie.Studio
// 			}
// 			if movie.Rating != 0 {
// 				temp.Next.Data.Rating = movie.Rating
// 			}
// 			if movie.Agerate != "-" {
// 				temp.Next.Data.Agerate = movie.Agerate
// 			}
// 			break
// 		}
// 		temp = temp.Next
// 	}
// }

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
