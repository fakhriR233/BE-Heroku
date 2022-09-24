package routes

import (
	"dumbflix_be/handlers"
	"dumbflix_be/pkg/middleware"
	"dumbflix_be/pkg/mysql"
	"dumbflix_be/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
  transactionRepository := repositories.RepositoryTransaction(mysql.DB)
  h := handlers.HandlerTransaction(transactionRepository)

  r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
  r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
  r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
  r.HandleFunc("/transaction/{id}", h.UpdatesTransaction).Methods("PATCH")
  r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
  r.HandleFunc("/notification", h.Notification).Methods("POST")
}