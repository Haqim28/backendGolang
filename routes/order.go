package routes

import (
	"backend/handlers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/order", h.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", h.GetOrderByCart).Methods("GET")
	r.HandleFunc("/order/{id}", h.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/order/{id}", h.UpdateOrder).Methods("PATCH")

}
