package films

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)                      // get the list of films
	router.Get("/{id}", FilmByID)                   // get one film by id
	router.Post("/", CreateFilm)                    // create new film
	router.Delete("/{id}", DeleteById)              // delete film by id
	router.Patch("/{id}", UpdateById)               // update film by id
	router.Get("/search/{keyword}", FilmsByKeyword) // get films by keyword

	return router
}
