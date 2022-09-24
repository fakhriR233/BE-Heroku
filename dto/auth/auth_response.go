package authdto

type AuthResponse struct {
  ID				  int			`json:"id"`
  FullName     string `gorm:"type: varchar(255)" json:"FullName"`
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
  Status    	string `gorm:"type: varchar(255)" json:"status"`
}

type LoginResponse struct {
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
  Status    	string `gorm:"type: varchar(255)" json:"status"`
}

type RegisterResponse struct {
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
  ID				  int			`json:"id"`
  FullName     string `gorm:"type: varchar(255)" json:"FullName"`
  Email       string `gorm:"type: varchar(255)" json:"email"`
  Token    	string `gorm:"type: varchar(255)" json:"token"`
  Status    	string `gorm:"type: varchar(255)" json:"status"`
  Gender 	  	string		`json:"gender" gorm:"type: varchar(255)"`
  Phone 	 	string		`json:"phone" gorm:"type: varchar(255)"`
  Address 	  	string		`json:"address" gorm:"type: varchar(255)"`
  Subscribe 	string		`json:"subscribe" gorm:"type: varchar(255)"`
}