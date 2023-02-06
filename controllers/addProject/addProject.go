package addproject

import (
	"html/template"
	"net/http"
)

func AddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/projectAdd.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var data = map[string]interface{}{
		"title" : "Add Project | Marcel",
		"isLogin" : true,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}