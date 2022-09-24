package filmdto

import "dumbflix_be/models"

type FilmRequest struct {
	ID					int							`json:"id" gorm:"primary_key:auto_increment"`
	Title				string						`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string						`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:text"`
	Year				int							`json:"year" form:"year" gorm:"type: int"`
	Category   			models.CategoryResponse    	`json:"category"`
	CategoryID 			int                			`json:"category_id" form:"category_id" gorm:"-"`
	Description			string						`json:"description" form:"description" gorm:"type:text"`
}

type FilmUpdateRequest struct {
	// ID					int							`json:"id" gorm:"primary_key:auto_increment"`
	Title				string						`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string						`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:text"`
	Year				int							`json:"year" form:"year" gorm:"type: int"`
	CategoryID 			int                			`json:"category_id" form:"category_id" gorm:"-"`
	Category   			models.CategoryResponse    	`json:"category"`
	Description			string						`json:"description" form:"description" gorm:"type:text"`
}