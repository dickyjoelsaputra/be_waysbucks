package routes

import (
	"be_waysbucks/handlers"
	"be_waysbucks/pkg/middleware"
	"be_waysbucks/pkg/mysql"
	"be_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	topingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(topingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")	
	r.HandleFunc("/toping", middleware.Auth(middleware.UploadFile(h.CreateToping))).Methods("POST")
	r.HandleFunc("/toping/{id}", middleware.Auth(middleware.UploadFile(h.UpdateToping))).Methods("PUT")
	r.HandleFunc("/toping/{id}", middleware.Auth(h.DeleteToping)).Methods("DELETE")
}