package categories

import (
	"net/http"
	db "tsi/films_website/database"
	e "tsi/films_website/error"

	"github.com/go-chi/render"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
	var categories []*Category

	result := db.DB.Model(&Category{}).Find(&categories)

	if result.Error != nil {
		render.Render(w, r, e.ErrEmptyTable(result.Error))
		return
	}
	render.RenderList(w, r, NewCategoryListResponse(categories))

}
