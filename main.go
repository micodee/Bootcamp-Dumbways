package main

import (
	"fmt"
	"net/http"
	"projek/config"
	"projek/controllers/auth"
	"projek/controllers/pages"
	"projek/controllers/project"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// connection database.go
	config.ConnectDB()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing pages
	router.HandleFunc("/", project.Home).Methods("GET")
	router.HandleFunc("/addProject", pages.AddProject).Methods("GET")
	router.HandleFunc("/contact", pages.Contact).Methods("GET")
	router.HandleFunc("/register", pages.FormRegister).Methods("GET")
	router.HandleFunc("/login", pages.FormLogin).Methods("GET")

// routing actions
	router.HandleFunc("/add", project.Add).Methods("POST")
	router.HandleFunc("/update/{id}", project.Update).Methods("GET")
	router.HandleFunc("/upost/{id}", project.UpdatePost).Methods("POST")
	router.HandleFunc("/delete/{id}", project.Delete).Methods("GET")
	router.HandleFunc("/detail/{id}", project.Detail).Methods("GET")

	// routing auth and session
	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	// create server port
	port := "5000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, router)
}
