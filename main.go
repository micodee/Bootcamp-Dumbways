package main

import (
	"fmt"
	"net/http"
	"projek/config"
	"projek/controllers/addProject"
	"projek/controllers/contact"
	"projek/controllers/home"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// create connection database.go
	config.ConnectDB()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing
	router.HandleFunc("/", home.Home).Methods("GET")
	router.HandleFunc("/addProject", addproject.AddProject).Methods("GET")
	router.HandleFunc("/contact", contact.Contact).Methods("GET")

	router.HandleFunc("/add", home.Add).Methods("POST")
	router.HandleFunc("/update/{id}", home.Update).Methods("GET")
	router.HandleFunc("/upost/{id}", home.UpdatePost).Methods("POST")
	router.HandleFunc("/delete/{id}", home.Delete).Methods("GET")
	router.HandleFunc("/detail/{id}", home.Detail).Methods("GET")

	// create port
	port := "5000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, router)

	// fmt.Println("Server Running on port 5000")
	// http.ListenAndServe("localhost:5000", router)
}
