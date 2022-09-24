package handlers

import (
	filmdto "dumbflix_be/dto/film"
	dto "dumbflix_be/dto/result"
	"dumbflix_be/models"
	"dumbflix_be/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
  }

// var path_file = "http://localhost:5000/uploads/"

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fmt.Println(os.Getenv("PATH_FILE"))

	film, err := h.FilmRepository.FindFilms()
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	for i, p := range film {
		film[i].ThumbnailFilm = os.Getenv("PATH_FILE") + p.ThumbnailFilm
	}

	// for i, p := range film {
	// 	film[i].ThumbnailFilm = path_file + p.ThumbnailFilm
	// }
  
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK, 
		Data: film,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	var film models.Film
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	film.ThumbnailFilm = os.Getenv("PATH_FILE") + film.ThumbnailFilm
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64)) 

	dataContex := r.Context().Value("dataFile") // add this code
  	filename := dataContex.(string)

	year, _ := strconv.Atoi(r.FormValue("Year"))
	category_id, _ := strconv.Atoi(r.FormValue("CategoryID"))
  
	request := filmdto.FilmRequest{
		Title:       		r.FormValue("Title"),
		Description:       	r.FormValue("Description"),
		Year:      			year,
		CategoryID: 		category_id,
	}
	
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }
  
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	film := models.Film{
		Title:    			request.Title,
		ThumbnailFilm:    	filename,
		Year:    			request.Year,
		CategoryID:    		request.CategoryID,
		Description:    	request.Description,
	}
  
	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	film, _ = h.FilmRepository.GetFilm(film.ID)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	// request := new(filmdto.FilmUpdateRequest) //take pattern data submission
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile") // add this code
  	filename := dataContex.(string)

	// fmt.Println(dataContex)

	// return

	year, _ := strconv.Atoi(r.FormValue("Year"))
	category_id, _ := strconv.Atoi(r.FormValue("CategoryID"))
  
	request := filmdto.FilmUpdateRequest{
		Title:       		r.FormValue("Title"),
		Description:       	r.FormValue("Description"),
		Year:      			year,
		CategoryID: 		category_id,
	}
  
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filmDataOld, _ := h.FilmRepository.GetFilm(id)
  
	film := models.Film{}
  
	if request.Title != "" {
		film.Title = request.Title
	}else {
		film.Title = filmDataOld.Title
	}

	// fmt.Println(request.ThumbnailFilm)

	if filename != "false" {
		film.ThumbnailFilm = filename
		// fmt.Println(filename)
	}else {
		film.ThumbnailFilm = filmDataOld.ThumbnailFilm
	}

	if request.Year != 0 {
		film.Year = request.Year
	}else {
		film.Year = filmDataOld.Year
	}

	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}else {
		film.CategoryID = filmDataOld.CategoryID
		film.Category = filmDataOld.Category
	}
	if request.Description != "" {
		film.Description = request.Description
	}else {
		film.Description = filmDataOld.Description
	}

	// if filename != "false" {
	// 	film.ThumbnailFilm = filename
	// } else {
	// 	film.ThumbnailFilm = filmDataOld.ThumbnailFilm
	// }

	data, err := h.FilmRepository.UpdateFilm(film,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	// fmt.Println(data)

	 data, _ = h.FilmRepository.GetFilm(id)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilmUpdate(data)}
	json.NewEncoder(w).Encode(response)
  }

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.FilmRepository.DeleteFilm(film,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteFilm(data)}
	json.NewEncoder(w).Encode(response)
  }

func convertResponseFilm(u models.Film) models.Film {
	return models.Film{
		ID:						u.ID,
	  Title:    				u.Title,
	  ThumbnailFilm:    		u.ThumbnailFilm,
	  Year:    					u.Year,
	  CategoryID:    			u.CategoryID,
	  Category:    				u.Category,
	  Description:    			u.Description,
	}
}

func convertResponseFilmUpdate(u models.Film) models.Film {
	return models.Film{
		ID:						u.ID,
	  Title:    				u.Title,
	  ThumbnailFilm:    		u.ThumbnailFilm,
	  Year:    					u.Year,
	  CategoryID:				u.CategoryID,
	  Category:    				u.Category,
	  Description:    			u.Description,
	}
}

func convertResponseDeleteFilm(u models.Film) filmdto.FilmDeleteResponse {
	return filmdto.FilmDeleteResponse{
	  ID:    u.ID,
	}
}