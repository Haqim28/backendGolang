package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetProductByPartner).Methods("GET")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.UploadFile(h.UpdateProduct)).Methods("PATCH")
	r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}
