package main

import (
	"be_waysbucks/database"
	"be_waysbucks/pkg/mysql"
	"be_waysbucks/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // ENV
	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }


	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	 //path file
	// r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))


	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}