package home

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projek/entities"
	"strconv"

	"github.com/gorilla/mux"
)

var project = []entities.Project{
	{
		Title : "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
		Durasi : "1 Bulan",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "Disini masih ada",
		Durasi : "2 Minggu",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "lalalaaa",
		Durasi : "6 Hari",
		Content : "App that used for dumbways student, it was deployed and can downloaded on playstore. Happy download",
	},
	{
		Title : "Hallo World",
		Durasi : "3 Hari",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
	},
	{
		Title : "bai bfdsl ksssu ldjkacau",
		Durasi : "5 Bulan",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
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

	fmt.Println("Project Name : " + r.PostForm.Get("pname"))
	fmt.Println("Description : " + r.PostForm.Get("desc"))

	var newAdd = entities.Project{
		Title : title,
		Content: content,
	}

	project = append(project, newAdd)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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