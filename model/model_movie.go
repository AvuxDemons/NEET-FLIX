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

// ██    ██ ██ ███████ ██     ██
// ██    ██ ██ ██      ██     ██
// ██    ██ ██ █████   ██  █  ██
//  ██  ██  ██ ██      ██ ███ ██
//   ████   ██ ███████  ███ ███

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

// ✅ Check Data
func CheckData(params int) bool {
	temp := database.DBMovie.Next
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

// ✅ Raw Data View By
func ViewMovieBy(params string) []entity.Movie {
	temp := database.DBMovie.Next
	movies := []entity.Movie{}
	if temp == nil {
		return movies
	}

	x, _ := strconv.Atoi(params)

	for temp != nil {
		if temp.Data.Id == x || strings.Contains(strings.ToLower(temp.Data.Title), strings.ToLower(params)) || strings.Contains(strings.Join(temp.Data.Genre, ","), params) || strings.Contains(strings.Join(temp.Data.Category, ","), params) {
			movies = append(movies, temp.Data)
		}
		temp = temp.Next
	}
	return movies
}

// ✅ Raw Data View By
// func ViewMovieBy(params string) entity.Movie {
// 	temp := database.DBMovie.Next
// 	if temp == nil {
// 		return entity.Movie{}
// 	}

// 	x, _ := strconv.Atoi(params)

// 	for temp != nil {
// 		if temp.Data.Id == x || strings.Contains(strings.Join(temp.Data.Genre, ","), params) || strings.Contains(strings.ToLower(temp.Data.Title), strings.ToLower(params)) {
// 			return temp.Data
// 		}
// 		temp = temp.Next
// 	}
// 	return entity.Movie{}
// }

// func UpdateMovie(Id int) {
// 	temp := &database.DBMovie
// 	for temp.Next != nil {
// 		if temp.Next.Data.Id == Id {
// 			if temp.Data.Title != "-" {
// 				temp.Next.Data.Title = temp.Data.Title
// 			}
// 			if temp.Data.Release != 0 {
// 				temp.Next.Data.Release = temp.Data.Release
// 			}
// 			if len(temp.Data.Genre) != 0 {
// 				temp.Next.Data.Genre = temp.Data.Genre
// 			}
// 			if len(temp.Data.Category) != 0 {
// 				temp.Next.Data.Category = temp.Data.Category
// 			}
// 			if temp.Data.Studio != "-" {
// 				temp.Next.Data.Studio = temp.Data.Studio
// 			}
// 			if temp.Data.Rating != 0 {
// 				temp.Next.Data.Rating = temp.Data.Rating
// 			}
// 			if temp.Data.Agerate != "-" {
// 				temp.Next.Data.Agerate = temp.Data.Agerate
// 			}
// 			break
// 		}
// 		temp = temp.Next
// 	}
// }

// ✅ Delete Data
func DeleteMovie(params int) {
	temp := &database.DBMovie
	// if temp == nil || temp.Next == nil {
	// 	fmt.Println("DATA MOVIE KOSONG")
	// 	return
	// }

	prev := temp
	current := temp.Next

	for current != nil {
		if current.Data.Id == params {
			prev.Next = current.Next
			fmt.Println("MOVIE BERHASIL DIHAPUS")
			return
		}
		prev = current
		current = current.Next
	}
	fmt.Println("MOVIE TIDAK DITEMUKAN")
}
