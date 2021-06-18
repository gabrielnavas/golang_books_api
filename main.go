package main

import (
	"books_api/database"
	factorycontroller "books_api/factory"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := database.MakePostgresSQLDatabase()

	category_controller := factorycontroller.MakeCategoryController(db)
	r.HandleFunc("/category", category_controller.Create).Methods("POST")
	r.HandleFunc("/category", category_controller.GetAll).Methods("GET", "OPTIONS")
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
