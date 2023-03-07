package main

import (
	"tsi/films_website/database"
	"tsi/films_website/resources/categories"
	filmcategory "tsi/films_website/resources/film_category"
	"tsi/films_website/resources/films"
	"tsi/films_website/server"
)

func main() {
	database.Init()

	database.DB.AutoMigrate(&films.Film{})
	database.DB.AutoMigrate(&categories.Category{})
	database.DB.AutoMigrate(&filmcategory.FilmCategory{})

	server.Init()

}
