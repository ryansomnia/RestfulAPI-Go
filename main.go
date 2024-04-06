package main

import (
	"RestfulAPI/config"
	"RestfulAPI/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	err := config.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}

	r := mux.NewRouter()

	// Routes

	r.HandleFunc("/get/user", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/post/user", handler.CreateUser).Methods("POST")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
