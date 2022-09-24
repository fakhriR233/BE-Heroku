package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
  UserRoutes(r)
  AuthRoutes(r)
  TransactionRoutes(r)
  CategoryRoutes(r)
  FilmRoutes(r)
  EpisodeRoutes(r)
}