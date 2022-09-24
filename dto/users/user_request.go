package usersdto

type CreateUserRequest struct {
  FullName     	string `json:"fullname" form:"name" validate:"required"`
  Email    		string `json:"email" form:"email" validate:"required"`
  Password 		string `json:"password" form:"password" validate:"required"`
  Gender 		string `json:"gender" form:"gender" validate:"required"`
  Phone 		string `json:"phone" form:"phone" validate:"required"`
  Address 		string `json:"address" form:"address" validate:"required"`
//   Subscribe 	string `json:"subscribe" form:"subscribe" validate:"required"`
}

type UpdateUserRequest struct {
  FullName     	string `json:"fullname" form:"name"`
  Email    		string `json:"email" form:"email"`
  Password		string `json:"password" form:"password"`
  Gender 		string `json:"gender" form:"gender" validate:"required"`
  Phone 		string `json:"phone" form:"phone" validate:"required"`
  Address 		string `json:"address" form:"address" validate:"required"`
//   Subscribe 	string `json:"subscribe" form:"subscribe" validate:"required"`
}