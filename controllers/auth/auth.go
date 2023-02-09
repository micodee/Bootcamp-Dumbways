package auth

import (
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset-utf-8")

	temp, err := template.ParseFiles("view/authRegister.html")

	
}