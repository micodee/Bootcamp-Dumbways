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

type postproject struct {
	Title string
	postdate string
	Durasi string
	Content string
}
var Projects = []postproject{
	{
		Title : "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		postdate : "12 Jul 2021 22:30 WIB",
		Durasi : "1 Bulan",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "Disini masih ada",
		postdate : "12 Jul 2021 22:30 WIB",
		Durasi : "2 Minggu",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "lalalaaa",
		postdate : "12 Jul 2021 22:30 WIB",
		Durasi : "6 Hari",
		Content : "App that used for dumbways student, it was deployed and can downloaded on playstore. Happy download",
	},
	{
		Title : "Hallo World",
		postdate : "12 Jul 2021 22:30 WIB",
		Durasi : "3 Hari",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "bai bfdsl ksssu ldjkacau",
		postdate : "12 Jul 2021 22:30 WIB",
		Durasi : "5 Bulan",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	}


func main() {
	// menyiapkan routingan
	router := mux.NewRouter()

	// create static folder for public
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// routing
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/project", Project).Methods("GET")
	router.HandleFunc("/detail/{id}", ProjectDetail).Methods("GET")
	router.HandleFunc("/addProject", AddProject).Methods("POST")
	router.HandleFunc("/contact", Contact).Methods("GET")

	fmt.Println("Server Running on port 5000")
	http.ListenAndServe("localhost:5000", router)
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

	resp := map[string]interface{}{
		"Data" : data,
		"Projects" : Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// function handling project
func Project (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/projectAdd.html")
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
	var tmpl, err = template.ParseFiles("./views/projectDetail.html")
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

	title := req.PostForm.Get("pname")
	content := req.PostForm.Get(("desc"))

	var newProject = postproject{
		Title: title,
		Content: content,
	}

	Projects = append(Projects, newProject)

	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}