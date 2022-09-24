package handlers

import (
	episodedto "dumbflix_be/dto/episode"
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

type handlerEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
  }

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}

func (h *handlerEpisode) FindEpisodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	episodes, err := h.EpisodeRepository.FindEpisodes()
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	for i, p := range episodes {
		episodes[i].ThumbnailFilm = os.Getenv("PATH_FILE") + p.ThumbnailFilm
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK, 
		Data: episodes,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpisode) GetEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	var episode models.Episode
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}

	episode.ThumbnailFilm = os.Getenv("PATH_FILE") + episode.ThumbnailFilm
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(episode)}
	json.NewEncoder(w).Encode(response)
  }

  func (h *handlerEpisode) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	dataContex := r.Context().Value("dataFile")
  	filename := dataContex.(string)

	film_id, _ := strconv.Atoi(r.FormValue("FilmID"))
  
	request := episodedto.EpisodeRequest{
		Title:       		r.FormValue("Title"),
		LinkFilm:       	r.FormValue("LinkFilm"),
		FilmID: 			film_id,
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	episode := models.Episode{
		Title:    			request.Title,
		LinkFilm:    		request.LinkFilm,
		ThumbnailFilm:    	filename,
		FilmID:    			request.FilmID,
	}
  
	episode, err = h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	episode, _ = h.EpisodeRepository.GetEpisode(episode.ID)
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: episode}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpisode) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	// request := new(episodedto.EpisodeUpdateRequest) //take pattern data submission
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//   w.WriteHeader(http.StatusBadRequest)
	//   response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	//   json.NewEncoder(w).Encode(response)
	//   return
	// }
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	dataContex := r.Context().Value("dataFile") // add this code
  	filename := dataContex.(string)

	film_id, _ := strconv.Atoi(r.FormValue("FilmID"))

	request := episodedto.EpisodeUpdateRequest{
		Title:       		r.FormValue("Title"),
		LinkFilm:       	r.FormValue("LinkFilm"),
		FilmID: 			film_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	episodeDataOld, _ := h.EpisodeRepository.GetEpisode(id)
  
	episode := models.Episode{}
  
	if request.Title != "" {
		episode.Title = request.Title
	}else {
		episode.Title = episodeDataOld.Title
	}

	if filename != "false" {
		episode.ThumbnailFilm = filename
		// fmt.Println(filename)
	}else {
		episode.ThumbnailFilm = episodeDataOld.ThumbnailFilm
	}

	if request.LinkFilm != "" {
		episode.LinkFilm = request.LinkFilm
	}else {
		episode.LinkFilm = episodeDataOld.LinkFilm
	}

	if request.FilmID != 0 {
		episode.FilmID = request.FilmID
		episodeDataNew, _ := h.EpisodeRepository.GetEpisode(episode.FilmID)
		episode.Film = episodeDataNew.Film
	}else {
		episode.FilmID = episodeDataOld.FilmID
		episode.Film = episodeDataOld.Film
	}
	
	data, err := h.EpisodeRepository.UpdateEpisode(episode,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisodeUpdate(data)}
	json.NewEncoder(w).Encode(response)
  }

func (h *handlerEpisode) DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	data, err := h.EpisodeRepository.DeleteEpisode(episode,id)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteEpisode(data)}
	json.NewEncoder(w).Encode(response)
  }

func convertResponseEpisode(u models.Episode) episodedto.EpisodeResponse {
	return episodedto.EpisodeResponse{
		ID:						u.ID,
	  Title:    				u.Title,
	  ThumbnailFilm:    		u.ThumbnailFilm,
	  LinkFilm:    				u.LinkFilm,
	  FilmID:    				u.FilmID,
	  Film:    					u.Film,
	}
}

func convertResponseEpisodeUpdate(u models.Episode) episodedto.EpisodeUpdateResponse {
	return episodedto.EpisodeUpdateResponse{
		ID:						u.ID,
	  Title:    				u.Title,
	  ThumbnailFilm:    		u.ThumbnailFilm,
	  LinkFilm:    				u.LinkFilm,
	  FilmID:    				u.FilmID,
	  Film:    					u.Film,
	}
}

func convertResponseDeleteEpisode(u models.Episode) episodedto.EpisodeDeleteResponse {
	return episodedto.EpisodeDeleteResponse{
	  ID:    u.ID,
	}
}