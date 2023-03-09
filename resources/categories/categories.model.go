package categories

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Category struct {
	CategoryId int       `gorm:"column:category_id;primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(25)"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRequest struct {
	*Category
}

type CategoryResponse struct {
	*Category
}

func NewCategoryResponse(categories *Category) *CategoryResponse {
	return &CategoryResponse{categories}
}

func NewCategoryListResponse(categories []*Category) []render.Renderer {
	list := []render.Renderer{}
	for _, category := range categories {
		list = append(list, NewCategoryResponse(category))
	}
	return list
}

func (f *CategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
