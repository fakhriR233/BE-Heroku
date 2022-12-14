package usersdto

type UserResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"FullName" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
	Subscribe string `json:"subscribe" form:"subscribe" validate:"required"`
}

type UserResponseDelete struct {
	ID int `json:"id"`
}