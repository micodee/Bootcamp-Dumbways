package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")

	fmt.Println("Server Running on port 5000")
	http.ListenAndServe("localhost:5000", router)
}

func hello (w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

// function handling index.html
func home (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/index.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}