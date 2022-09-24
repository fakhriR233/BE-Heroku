package handlers

import (
	categorydto "dumbflix_be/dto/category"
	dto "dumbflix_be/dto/result"
	"dumbflix_be/models"
	"dumbflix_be/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
  }

func HandlerCategory(CategoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}

func (h *handlerCategory) FindCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	category, err := h.CategoryRepository.FindCategories()
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK, 
		Data: category,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	var category models.Category
	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(category)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerCategory) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(categorydto.CategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	category := models.Category{
		Name:    	request.Name,
	}
  
	category, err = h.CategoryRepository.CreateCategory(category)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	category, _ = h.CategoryRepository.GetCategory(category.ID)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	request := new(categorydto.CategoryUpdateRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	categoryDataOld, _ := h.CategoryRepository.GetCategory(id)
  
	category := models.Category{}
  
	if request.Name != "" {
		category.Name = request.Name
	}else {
		category.Name = categoryDataOld.Name
	}
	
	data, err := h.CategoryRepository.UpdateCategory(category,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategoryUpdate(data)}
	json.NewEncoder(w).Encode(response)
  }

func (h *handlerCategory) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.CategoryRepository.DeleteCategory(category,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteCategory(data)}
	json.NewEncoder(w).Encode(response)
  }

func convertResponseCategory(u models.Category) models.Category {
	return models.Category{
		ID:			u.ID,
	  Name:    		u.Name,
	}
}

func convertResponseCategoryUpdate(u models.Category) categorydto.CategoryUpdateResponse {
	return categorydto.CategoryUpdateResponse{
		ID:			u.ID,
	  Name:    		u.Name,
	}
}

func convertResponseDeleteCategory(u models.Category) categorydto.CategoryDeleteResponse {
	return categorydto.CategoryDeleteResponse{
	  ID:    u.ID,
	}
}