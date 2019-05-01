package main

import (
	"fmt"
	"gomongo/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize API routes
func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controllers.GetAllPeople).Methods("GET")
	router.HandleFunc("/api/{id}", controllers.GetPersonById).Methods("GET")
	router.HandleFunc("/api", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/api", controllers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/api", controllers.UpdatePerson).Methods("PUT")
	return router
}

func main() {
	// Starting caching process
	go controllers.SyncCache()

	router := initRouter()

	fmt.Println("Starting server on port 4500...")
	log.Fatal(http.ListenAndServe(":4500", router))
}
