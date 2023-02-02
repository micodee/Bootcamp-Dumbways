package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// map string interface membuat string ke interface adalah tipe data utk menyimpan koleksi atau sejenis seperti object di javascript tanpa harus menggunakan banyak variable
var data = map[string]interface{}{
	"title" : "Personal Web",
	"isLogin" : true,
}

func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing
	router.HandleFunc("/hello", hello).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/project", Project).Methods("GET")
	router.HandleFunc("/contact", Contact).Methods("GET")
	router.HandleFunc("/detail/{id}", ProjectDetail).Methods("GET")
	router.HandleFunc("/addProject", AddProject).Methods("POST")

	fmt.Println("Server Running on port 5000")
	http.ListenAndServe("localhost:5000", router)
}

func hello (w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}




// function handling home
func Home (w http.ResponseWriter, req *http.Request) {
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
	tmpl.Execute(w, data)
}

// function handling project
func Project (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/my-project.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

// function handling contact
func Contact (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/contact-me.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

// function handling project detail
func ProjectDetail (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// manangkap id (id, _ (tanda _ tidak ingin menampilkan eror))
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/detail-project.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	resp := map[string]interface{}{
		"Data" : data,
		"ID" : id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// function handling add project
func AddProject (w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)		
	}

	fmt.Println("title : " + req.PostForm.Get("title"))
	fmt.Println("content : " + req.PostForm.Get("content"))

	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}