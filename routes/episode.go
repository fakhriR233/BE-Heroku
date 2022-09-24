package routes

import (
	"dumbflix_be/handlers"
	"dumbflix_be/pkg/middleware"
	"dumbflix_be/pkg/mysql"
	"dumbflix_be/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
  episodeRepository := repositories.RepositoryEpisode(mysql.DB)
  h := handlers.HandlerEpisode(episodeRepository)

  r.HandleFunc("/episodes", h.FindEpisodes).Methods("GET")
  r.HandleFunc("/episode/{id}", h.GetEpisode).Methods("GET")
  r.HandleFunc("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode))).Methods("POST")
  r.HandleFunc("/episode/{id}", middleware.Auth(middleware.UploadFile(h.UpdateEpisode))).Methods("PATCH")
  r.HandleFunc("/episode/{id}", h.DeleteEpisode).Methods("DELETE")
}