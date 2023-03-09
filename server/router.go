package server

import (
	"tsi/films_website/resources/categories"
	"tsi/films_website/resources/films"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"}, AllowCredentials: false, MaxAge: 300}))

	router.Mount("/films", films.Routes())
	router.Mount("/categories", categories.Routes())

	return router
}
