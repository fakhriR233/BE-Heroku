package handlers

import (
	authdto "dumbflix_be/dto/auth"
	dto "dumbflix_be/dto/result"
	"dumbflix_be/models"
	"dumbflix_be/pkg/bcrypt"
	jwtToken "dumbflix_be/pkg/jwt"
	"dumbflix_be/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
  AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
  return &handlerAuth{AuthRepository}
}

// {"code":200,"data":{"email":"fakhriramadhan233@gmail.com","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjMyMTM3NzgsImZ1bGxuYW1lIjoiTW9oYW1hZCBGYWtocmkgUmFtYWRoYW4iLCJpZCI6MH0.TA2OnnJo5eajEBpOf85MNYS9KHY0mNIHmtGccbA5nMQ"}}
// {"code":200,"data":{"email":"jody@mail.com","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjMyMTgzMTUsImZ1bGxuYW1lIjoiam9keSBzIiwiaWQiOjB9.2ZzJJzWVkhFLfKguVeFQWlylcqlgZSLHIYXlprmsW-0"}}


func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(authdto.RegisterRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  password, err := bcrypt.HashingPassword(request.Password)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  user := models.User{
    FullName:     	request.FullName,
    Email:    		request.Email,
    Password: 		password,
    Gender: 		request.Gender,
    Phone: 			request.Phone,
    Address: 		request.Address,
	Subscribe: 		"false",
	Status: 		"user",
  }

  
  
  
  data, err := h.AuthRepository.Register(user)
  if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	}

claims := jwt.MapClaims{}
claims["id"] = data.ID
claims["fullname"] = data.FullName
claims["status"] = data.Status
claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

token, errGenerateToken := jwtToken.GenerateToken(&claims)
if errGenerateToken != nil {
	log.Println(errGenerateToken)
	fmt.Println("Unauthorize")
	return
}

  loginResponse := authdto.RegisterResponse{
	Email:    		data.Email,
	Token:			token,
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	user := models.User{
	  Email:    request.Email,
	  Password: request.Password,
	}
  
	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["fullname"] = user.FullName
	claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired
  
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
	  log.Println(errGenerateToken)
	  fmt.Println("Unauthorize")
	  return
	}
  
	loginResponse := authdto.LoginResponse{
	  Email:    	user.Email,
	  Status:    	user.Status,
	  Token:    	token,
	}
  
	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
  
  }

  func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	
	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       		user.ID,
		FullName:     	user.FullName,
		Email:    		user.Email,
		Status:   		user.Status,
		Gender:   		user.Gender,
		Phone:   		user.Phone,
		Subscribe:   	user.Subscribe,
		Address:   		user.Address,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}

// func convertResponseAuth(u models.User) authdto.AuthResponse {
// 	return authdto.AuthResponse{
// 	  Email:    	u.Email,
// 	  Password: 	password,
// 	}
// }

func convertResponseRegister(u models.User) authdto.RegisterResponse {
	return authdto.RegisterResponse{
	  Email:    	u.Email,
	}
}