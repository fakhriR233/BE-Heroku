package filmdto

import "dumbflix_be/models"

type FilmResponse struct {
	ID					int							`json:"id"`
	Title				string						`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string						`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:text"`
	Year				int							`json:"year" form:"year" gorm:"type: int"`
	CategoryID 			int                			`json:"-"`
	Category   			models.CategoryResponse    	`json:"category"`
	Description			string						`json:"description" form:"description" gorm:"type:text"`
}

type FilmUpdateResponse struct {
	ID					int									`json:"id"`
	Title				string								`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm		string								`json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:text"`
	Year				int									`json:"year" form:"year" gorm:"type: int"`
	CategoryID 			int                					`json:"category_id"`
	Category   			models.CategoryResponse   	`json:"category" gorm:"foreignKey:CategoryID"`
	Description			string								`json:"description" form:"description" gorm:"type:text"`
}

type FilmDeleteResponse struct {
	ID					int							`json:"id"`
}