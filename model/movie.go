package model

import (
	"NeetFlix/database"
	"NeetFlix/entity"
)

// ✅ Insert Data To Database
func InsertData(movie entity.Movie) {
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
