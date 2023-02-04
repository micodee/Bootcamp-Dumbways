package main

import (
	"fmt"
	"projek/controllers/home"
	"projek/controllers/addProject"
	"projek/controllers/contact"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing
	// router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/", home.Home).Methods("GET")
	router.HandleFunc("/addProject", addproject.AddProject).Methods("GET")
	router.HandleFunc("/contact", contact.Contact).Methods("GET")

	router.HandleFunc("/detail/{id}", home.Detail).Methods("GET")
	router.HandleFunc("/add", home.Add).Methods("POST")
	router.HandleFunc("/delete/{id}", home.Delete).Methods("GET")
	

	fmt.Println("Server Running on port 5000")
	http.ListenAndServe("localhost:5000", router)
}
