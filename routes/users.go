package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/role/{role}", h.GetRole).Methods("GET")
	r.HandleFunc("/user/role/{role}/{id}", h.GetRoleId).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.UploadFile(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
