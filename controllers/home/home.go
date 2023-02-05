package home

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projek/entities"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var project = []entities.Project{
	{
		Title : "Hallo World",
		Sdate: time.Now().Format("2023-02-05"),
		Edate: time.Now().Format("2023-02-09"),
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
		Tnode: true,
		Treact: true,
		Tjs: true,
		Thtml: true,
	},
	{
		Title : "bai bfdsl ksssu ldjkacau",
		Sdate: time.Now().Format("2023-01-02"),
		Edate: time.Now().Format("2023-01-06"),
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
		Tnode: true,
		Treact: true,
		Tjs: false,
		Thtml: false,
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

		// parsing template html
	var tmpl, err = template.ParseFiles("./views/index.html")
	if err != nil {
		panic(err)
	}
	
var data = map[string]interface{}{
	"title" : "Home",
	"isLogin" : true,
}

	resp := map[string]interface{}{
		"Data" : data,
		"Projects" : project,
	}



	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

func Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		log.Fatal(err)
	}

	title := r.PostForm.Get("pname")
	content := r.PostForm.Get(("desc"))
	SD := r.PostForm.Get("sdate")
	ED := r.PostForm.Get("edate")
	node := false
	react := false
	js := false
	hateml := false

	// if checked
	if r.FormValue("nodejs") != "" {
		node = true
	}
	if r.FormValue("reactjs") != "" {
		react = true
	}
	if r.FormValue("js") != "" {
		js = true
	}
	if r.FormValue("html") != "" {
		hateml = true
	}

	var newAdd = entities.Project{
		Title : title,
		Sdate: SD,
		Edate: ED,
		Content: content,
		Tnode: node,
		Treact: react,
		Tjs: js,
		Thtml: hateml,
	}

	project = append(project, newAdd)

	fmt.Println("Project Name : " + r.PostForm.Get("pname"))
	fmt.Println("Description : " + r.PostForm.Get("desc"))
	fmt.Println("Start Date : " + r.PostForm.Get("sdate"))
	fmt.Println("End Date : " + r.PostForm.Get("edate"))
	fmt.Println("Node : " + r.PostForm.Get("nodejs"))
	fmt.Println("React : " + r.PostForm.Get("reactjs"))
	fmt.Println("Javascript : " + r.PostForm.Get("js"))
	fmt.Println("HTML : " + r.PostForm.Get("html"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
}


func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	
	project = append(project[:id], project[id+1:]...)

	fmt.Println(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}



func Detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// manangkap id (id, _ (tanda _ tidak ingin menampilkan eror))
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// parsing template html
	var tmpl, err = template.ParseFiles("./views/projectDetail.html")
	// error handling
	if err != nil {
		panic(err)
	}

	var data = map[string]interface{}{
		"title" : "Detail Project",
		"isLogin" : true,
	}

	resp := map[string]interface{}{
		"Data" : data,
		"ID" : id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}