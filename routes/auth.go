package routes

import (
	"be_waysbucks/handlers"
	"be_waysbucks/pkg/middleware"
	"be_waysbucks/pkg/mysql"
	"be_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerAuth(userRepository)

  r.HandleFunc("/register", middleware.UploadFile(h.Register)).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST")
}