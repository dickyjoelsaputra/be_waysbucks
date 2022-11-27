package routes

import (
	"be_waysbucks/handlers"
	"be_waysbucks/pkg/middleware"
	"be_waysbucks/pkg/mysql"
	"be_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/orders", middleware.Auth(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/orders/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/orders/{id}", middleware.Auth(h.GetOrder)).Methods("GET")
	r.HandleFunc("/mytransaction", middleware.Auth(h.MyTransaction)).Methods("GET")
}