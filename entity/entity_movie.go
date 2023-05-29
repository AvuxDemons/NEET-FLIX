package entity

type Movie struct {
	Release, Id            int
	Title, Studio, Agerate string
	Genre, Category        []string
	Rating                 float64
}

type DatabaseMovie struct {
	Data Movie
	Next *DatabaseMovie
}

var GenreList = []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Historical", "Horror", "Mistery", "Romance", "Sci-fi", "Sport", "Supranatural", "Thriller"}
var CategoryList = []string{"Movie", "Anime", "Series", "Cartoon", "Documenter"}
var AgeRate = []string{"KIDS | 6-11", "TEEN | 12-18", "ADULT | 18+"}
