package categories

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListCategories)
	return router
}
