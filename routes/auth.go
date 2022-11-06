package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")

}
