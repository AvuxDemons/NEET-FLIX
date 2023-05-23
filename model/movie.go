package model

import (
	"NeetFlix/database"
	"NeetFlix/entity"
)

// ✅ Insert Data To Database
func InsertMovie(movie entity.Movie) {
	tempMovie := entity.DatabaseMovie{}
	tempMovie.Data = movie
	temp := &database.DBMovie

	if temp.Next == nil {
		temp.Next = &tempMovie
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &tempMovie
	}
}

// ✅ Raw Data All Movie
func ViewAllMovie() []entity.Movie {
	temp := database.DBMovie.Next
	movies := []entity.Movie{}
	for temp != nil {
		movies = append(movies, temp.Data)
		temp = temp.Next
	}
	return movies
}

func UpdateMovie(movie entity.Movie) {
	temp := &database.DBMovie
	for temp.Next != nil {
		if temp.Next.Data.Id == movie.Id {
			if movie.Title != "-" {
				temp.Next.Data.Title = movie.Title
			}
			if movie.Release != 0 {
				temp.Next.Data.Release = movie.Release
			}
			if len(movie.Genre) != 0 {
				temp.Next.Data.Genre = movie.Genre
			}
			if len(movie.Category) != 0 {
				temp.Next.Data.Category = movie.Category
			}
			if movie.Studio != "-" {
				temp.Next.Data.Studio = movie.Studio
			}
			if movie.Rating != 0 {
				temp.Next.Data.Rating = movie.Rating
			}
			if movie.Agerate != "-" {
				temp.Next.Data.Agerate = movie.Agerate
			}
			break
		}
		temp = temp.Next
	}
}

// ✅ Delete Data
func DeleteData(params int) {
	temp := &database.DBMovie
	if temp == nil || temp.Next == nil {
		// Data Kosong
		return
	}

	prev := temp
	current := temp.Next

	for current != nil {
		if current.Data.Id == params {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}
