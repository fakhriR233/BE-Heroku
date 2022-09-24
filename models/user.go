package models

import "time"

// User model struct
type User struct {
  ID          	int			`json:"id"`
  FullName 		string		`json:"fullname" gorm:"type: varchar(255)"`
  Email		    string 		`json:"email" gorm:"type: varchar(255)"`
  Password 	  	string		`json:"password" gorm:"type: varchar(255)"`
  Gender 	  	string		`json:"gender" gorm:"type: varchar(255)"`
  Phone 	 	string		`json:"phone" gorm:"type: varchar(255)"`
  Address 	  	string		`json:"address" gorm:"type: varchar(255)"`
  Subscribe 	string		`json:"subscribe" gorm:"type: varchar(255)"`
  Status      string    `json:"status" gorm:"type: varchar(255)"`
  Token      string    `json:"token" gorm:"type: varchar(255)"`
  CreatedAt 	time.Time	`json:"created_at"`
  UpdatedAt 	time.Time	`json:"updated_at"`
}

type UserResponse struct {
  ID          	int			`json:"id"`
  FullName 		string		`json:"fullname" gorm:"type: varchar(255)"`
  Email		    string 		`json:"email" gorm:"type: varchar(255)"`
  Password 	  	string		`json:"password" gorm:"type: varchar(255)"`
  Gender 	  	string		`json:"gender" gorm:"type: varchar(255)"`
  Phone 	 	string		`json:"phone" gorm:"type: varchar(255)"`
  Address 	  	string		`json:"address" gorm:"type: varchar(255)"`
  Subscribe 	string		`json:"subscribe" gorm:"type: varchar(255)"`
  Status      string    `json:"status" gorm:"type: varchar(255)"`
  Token      string    `json:"token" gorm:"type: varchar(255)"`
  CreatedAt 	time.Time	`json:"created_at"`
  UpdatedAt 	time.Time	`json:"updated_at"`
}

func (UserResponse) TableName() string {
	return "users"
}