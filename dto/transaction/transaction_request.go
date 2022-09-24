package transactiondto

import "dumbflix_be/models"

type TransactionRequest struct {
	ID				int							`json:"id" validate:"required"`
	StartDate		string						`json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate			string						`json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	User   			models.UserResponse    		`json:"userId"`
	UserID 			int                			`json:"user_id" form:"user_id" gorm:"-"`
	Attache			string						`json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status			string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}

type TransactionUpdateRequest struct {
	// ID				int							`json:"id" gorm:"primary_key:auto_increment"`
	StartDate		string						`json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate			string						`json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	UserID 			int                			`json:"user_id" form:"user_id" gorm:"-"`
	User   			models.UserResponse    		`json:"userId"`
	Attache			string						`json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status			string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}

type TransactionUpdateResponse struct {
	ID				int							`json:"id"`
	StartDate		string						`json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	DueDate			string						`json:"dueDate" form:"dueDate" gorm:"type:varchar(255)"`
	UserID 			int                			`json:"user_id" form:"user_id" gorm:"-"`
	User   			models.UserResponse    		`json:"userId"`
	Attache			string						`json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status			string						`json:"status" form:"status" gorm:"type: varchar(255)"`
}

type TransactionDeleteResponse struct {
	ID				int							`json:"id"`
}