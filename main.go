package main

import (
	"books_api/controller"
	"books_api/database"
	"books_api/repository"
	servicecategory "books_api/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := database.MakePostgresSQLDatabase()

	category_repository := repository.NewCategoryRepository(db)
	category_service := servicecategory.NewCategoryService(category_repository)
	category_controller := controller.NewCategoryRepository(category_service)
	r.HandleFunc("/category", category_controller.Create).Methods("POST")

	// r.HandleFunc("/category").Methods("GET", "OPTIONS")
	// r.HandleFunc("/category/{category_id}").Methods("GET", "OPTIONS")
	// r.HandleFunc("/category").Methods("PUT", "OPTIONS")
	// r.HandleFunc("/category").Methods("DELETE", "OPTIONS")

	config_handlers := configCors()(r)
	log.Fatal(http.ListenAndServe(":8000", config_handlers))
}

func configCors() func(http.Handler) http.Handler {
	headersOptions := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOptions := handlers.AllowedOrigins([]string{"*"})
	methods := []string{}
	methods = append(methods, "POST")
	methodsOptions := handlers.AllowedMethods(methods)
	return handlers.CORS(originsOptions, headersOptions, methodsOptions)

}
