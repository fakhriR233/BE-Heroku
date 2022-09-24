package categorydto

type CategoryResponse struct {
	ID					int							`json:"id"`
	Name				string						`json:"name" form:"name" gorm:"type: varchar(255)"`
}

type CategoryDeleteResponse struct {
	ID				int						`json:"id"`
}

type CategoryUpdateResponse struct {
	ID				int						`json:"id"`
	Name			string					`json:"name" form:"name" gorm:"type: varchar(255)"`
}