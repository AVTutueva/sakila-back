package categories

import (
	"time"
)

type Category struct {
	CategoryId int       `gorm:"column:category_id;primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(25)"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (Category) TableName() string {
	return "category"
}
