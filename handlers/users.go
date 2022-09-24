package handlers

import (
	dto "dumbflix_be/dto/result"
	usersdto "dumbflix_be/dto/users"
	"dumbflix_be/models"
	"dumbflix_be/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerUser struct {
  UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
  return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  users, err := h.UserRepository.FindUsers()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: users}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
  json.NewEncoder(w).Encode(response)
}

 // Write this code
 func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(usersdto.CreateUserRequest)
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
  
	// data form pattern submit to pattern entity db user
	user := models.User{
	  FullName:     request.FullName,
	  Email:    	request.Email,
	  Password: 	request.Password,
	  Gender: 		request.Gender,
	  Phone: 		request.Phone,
	  Address: 		request.Address,
	  Subscribe: 	"false",
	}
  
	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  json.NewEncoder(w).Encode(err.Error())
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(usersdto.UpdateUserRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userDataOld, _ := h.UserRepository.GetUser(id)
  
	user := models.User{}
  
	if request.FullName != "" {
		user.FullName = request.FullName
	  }else {
		  user.FullName = userDataOld.FullName
	  }
	
	  if request.Email != "" {
		user.Email = request.Email
	  }else {
		  user.Email = userDataOld.Email
	  }
	
	  if request.Password != "" {
		user.Password = request.Password
	  }else {
		  user.Password = userDataOld.Password
	  }

	  if request.Gender != "" {
		user.Gender = request.Gender
	  }else {
		  user.Gender = userDataOld.Gender
	  }

	  if request.Phone != "" {
		user.Phone = request.Phone
	  }else {
		  user.Phone = userDataOld.Phone
	  }

	  if request.Address != "" {
		user.Address = request.Address
	  }else {
		  user.Address = userDataOld.Address
	  }

	  user.Subscribe = userDataOld.Subscribe
  
	data, err := h.UserRepository.UpdateUser(user,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.UserRepository.DeleteUser(user,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDelete(data)}
	json.NewEncoder(w).Encode(response)
  }

func convertResponse(u models.User) usersdto.UserResponse {
  return usersdto.UserResponse{
    ID:       	u.ID,
    FullName:   u.FullName,
    Email:    	u.Email,
    Password: 	u.Password,
	Gender:		u.Gender,
	Phone:		u.Phone,
	Address:	u.Address,
	Subscribe: 	u.Subscribe,
  }
}

func convertResponseDelete(u models.User) usersdto.UserResponseDelete {
  return usersdto.UserResponseDelete{
    ID:       	u.ID,
  }
}