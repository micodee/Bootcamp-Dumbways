package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("content.Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}).Methods("GET")

	port := "5000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, route)
}