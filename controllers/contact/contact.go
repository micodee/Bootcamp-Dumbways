package contact

import (
	"html/template"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	var tmpl, err = template.ParseFiles("./views/contact-me.html")
	// error handling
	if err != nil {
		panic(err)
	}

	var data = map[string]interface{}{
		"title" : "Contact Me",
		"isLogin" : true,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}