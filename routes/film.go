package routes

import (
	"dumbflix_be/handlers"
	"dumbflix_be/pkg/middleware"
	"dumbflix_be/pkg/mysql"
	"dumbflix_be/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
  filmRepository := repositories.RepositoryFilm(mysql.DB)
  h := handlers.HandlerFilm(filmRepository)

  r.HandleFunc("/films", h.FindFilms).Methods("GET")
  r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
  r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
  r.HandleFunc("/film/{id}", middleware.Auth(middleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
  r.HandleFunc("/film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE")
}