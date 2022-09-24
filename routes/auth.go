package routes

import (
	"dumbflix_be/handlers"
	"dumbflix_be/pkg/middleware"
	"dumbflix_be/pkg/mysql"
	"dumbflix_be/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerAuth(userRepository)

  r.HandleFunc("/register", h.Register).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST")
  r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}