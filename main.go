package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	router.HandleFunc("/", func (w http.ResponseWriter, req *http.Request)  {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}).Methods("GET")

	fmt.Println("Server Running in port 5000")
	http.ListenAndServe("localhost:5000", router)
}