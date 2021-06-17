package main

import (
	"books_api/controller"
	"books_api/database"
	"books_api/repository"
	"books_api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := database.MakePostgresSQLDatabase()

	category_repository := repository.NewCategoryRepository(db)
	category_service := service.NewCategoryService(category_repository)
	category_controller := controller.NewCategoryRepository(category_service)
	r.HandleFunc("/category", category_controller.Create).Methods("POST")

	// r.HandleFunc("/category").Methods("GET", "OPTIONS")
	// r.HandleFunc("/category/{category_id}").Methods("GET", "OPTIONS")
	// r.HandleFunc("/category").Methods("PUT", "OPTIONS")
	// r.HandleFunc("/category").Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", r))
}
