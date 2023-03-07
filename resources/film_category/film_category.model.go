package filmcategory

import (
	"time"
)

type FilmCategory struct {
	FilmId             int       `gorm:"column:film_id;primaryKey"`
	CategoryId         int       `gorm:"column:categoryid;primaryKey"`
	LastUpdate         time.Time `gorm:"autoUpdateTime"`
	FilmFilmId         int       `gorm:"type:tinyint"`
	CategoryCategoryId int       `gorm:"type:tinyint"`
	FilmReferId        int       `gorm:"type:tinyint"`
	CategoryRefer      int       `gorm:"type:tinyint"`
}

func (FilmCategory) TableName() string {
	return "film_category"
}
