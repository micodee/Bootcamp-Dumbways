package main

import (
	"fmt"
	"net/http"
	"projek/config"
	"projek/controllers/pages"
	"projek/controllers/project"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// create connection database.go
	config.ConnectDB()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing pages
	router.HandleFunc("/", project.Home).Methods("GET")
	router.HandleFunc("/addProject", pages.AddProject).Methods("GET")
	router.HandleFunc("/contact", pages.Contact).Methods("GET")

// routing actions
	router.HandleFunc("/add", project.Add).Methods("POST")
	router.HandleFunc("/update/{id}", project.Update).Methods("GET")
	router.HandleFunc("/upost/{id}", project.UpdatePost).Methods("POST")
	router.HandleFunc("/delete/{id}", project.Delete).Methods("GET")
	router.HandleFunc("/detail/{id}", project.Detail).Methods("GET")

	// routing auth and session
	router.HandleFunc("/register", pages.Register).Methods("GET")
	router.HandleFunc("/login", pages.Login).Methods("GET")

	// create port
	port := "5000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, router)

	// fmt.Println("Server Running on port 5000")
	// http.ListenAndServe("localhost:5000", router)
}
