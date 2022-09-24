package models

import "time"

type Film struct {
	ID					int					`json:"id" gorm:"primary_key:auto_increment"`
	Title				string				`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string				`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year				int					`json:"year" form:"year" gorm:"type: int"`
	CategoryID 			int                	`json:"category_id"`
	Category   			CategoryResponse   	`json:"category" form:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description			string				`json:"description" form:"description" gorm:"type: text"`
  	CreatedAt  			time.Time            `json:"-"`
  	UpdatedAt  			time.Time            `json:"-"`
}

type FilmResponse struct {
	ID					int							`json:"id"`
	Title				string						`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string						`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year				int							`json:"year" form:"year" gorm:"type: int"`
	CategoryID 			int                			`json:"category_id"`
	Category   			CategoryResponse    		`json:"category" form:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description			string						`json:"description" form:"description" gorm:"type: text"`
}

func (FilmResponse) TableName() string {
	return "films"
  }