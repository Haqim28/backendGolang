package routes

import (
	"backend/handlers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{id}", h.UpdateCart).Methods("PATCH")
	// r.HandleFunc("/cart/pending", h.FindPendingCart("GET"))
	// r.HandleFunc("/orders/{id}", h.GetOrderByCart).Methods("GET")
}
