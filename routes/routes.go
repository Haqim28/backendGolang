package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
	OrderRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)
}
