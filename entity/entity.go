package entity

type Movie struct {
	Release, Id            int
	Title, Studio, Agerate string
	Genre, Category        []string
	Rating                 float32
}

type DatabaseMovie struct {
	Data Movie
	Next *DatabaseMovie
}
