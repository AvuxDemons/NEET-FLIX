package controller

import (
	"NeetFlix/database"
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

// ██ ███    ██ ███████ ███████ ██████  ████████
// ██ ████   ██ ██      ██      ██   ██    ██
// ██ ██ ██  ██ ███████ █████   ██████     ██
// ██ ██  ██ ██      ██ ██      ██   ██    ██
// ██ ██   ████ ███████ ███████ ██   ██    ██

// ✅ Insert Data
func InsertMovie(Id int, Title string, Release int, Category []string, Studio string, Genre []string, AgeRate string, Rating float64) {
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

// ██    ██ ██████  ██████   █████  ████████ ███████
// ██    ██ ██   ██ ██   ██ ██   ██    ██    ██
// ██    ██ ██████  ██   ██ ███████    ██    █████
// ██    ██ ██      ██   ██ ██   ██    ██    ██
//  ██████  ██      ██████  ██   ██    ██    ███████

func UpdateMovie(Id int, Title string, Release int, Category []string, Studio string, Genre []string, AgeRate string, Rating float64) {
	temp := &database.DBMovie
	for temp != nil {
		if temp.Data.Id == Id {
			if Title != "-" {
				temp.Next.Data.Title = Title
			}
			if Release != 0 {
				temp.Next.Data.Release = Release
			}
			if len(Genre) != 0 {
				temp.Next.Data.Genre = Genre
			}
			if len(Category) != 0 {
				temp.Next.Data.Category = Category
			}
			if Studio != "-" {
				temp.Next.Data.Studio = Studio
			}
			if Rating != 0 {
				temp.Next.Data.Rating = Rating
			}
			if AgeRate != "-" {
				temp.Next.Data.Agerate = AgeRate
			}
			break
		}
		temp = temp.Next
	}
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
