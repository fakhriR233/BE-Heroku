package categorydto

type CategoryRequest struct {
	ID					int							`json:"id"`
	Name				string						`json:"name" form:"name" gorm:"type: varchar(255)"`
}

type CategoryUpdateRequest struct {
	ID					int							`json:"id"`
	Name				string						`json:"name" form:"name" gorm:"type: varchar(255)"`
}