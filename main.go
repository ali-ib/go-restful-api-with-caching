package main

import (
	"fmt"
	"go-restful-api-with-caching/configs"
	"go-restful-api-with-caching/controllers"
	"log"
	"net/http"
	"strconv"

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

	port := strconv.Itoa(configs.Configs.PORTNUMBER)
	fmt.Printf("Starting server on port %s...\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
