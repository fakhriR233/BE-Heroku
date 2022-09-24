package routes

import (
	"dumbflix_be/handlers"
	"dumbflix_be/pkg/mysql"
	"dumbflix_be/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
  categoryRepository := repositories.RepositoryCategory(mysql.DB)
  h := handlers.HandlerCategory(categoryRepository)

  r.HandleFunc("/categories", h.FindCategories).Methods("GET")
  r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
  r.HandleFunc("/category", h.CreateCategory).Methods("POST")
  r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")
  r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE")
}