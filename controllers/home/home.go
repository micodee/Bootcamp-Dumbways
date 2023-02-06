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
		Sdate: "05 February 2023",
		Edate: "12 February 2023",
		Duration: "1 Weeks",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error, aut voluptas officiis odio tempora reprehenderit quos voluptates mollitia explicabo. Dolores, tempore expedita?",
		Tnode: true,
		Treact: true,
		Tjs: true,
		Thtml: true,
	},
	{
		Title : "bai bfdsl ksssu ldjkacau",
		Sdate: "16 February 2023",
		Edate: "22 February 2023",
		Duration: "6 Days",
		Content : "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Dolor provident culpa magni eum ab voluptatum, iste error",
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

	formatLayout := "2006-01-02"
	sDate, _ := time.Parse(formatLayout, SD)
	sDateFormat := sDate.Format("02 January 2006")
	eDate, _ := time.Parse(formatLayout, ED)
	eDateFormat := eDate.Format("02 January 2006")

	
	durasi := eDate.Sub(sDate)
	// days := int(duration.Hours() / 24)
	// months := int(days / 30)
	// years := int(months / 12)
	var distance string
	if durasi.Hours()/24 < 7 {
		distance = strconv.FormatFloat(durasi.Hours()/24, 'f', 0, 64) + " Days"
	} else if durasi.Hours()/24/7 < 4 {
		distance = strconv.FormatFloat(durasi.Hours()/24/7, 'f', 0, 64) + " Weeks"
	} else if durasi.Hours()/24/30 < 12 {
		distance = strconv.FormatFloat(durasi.Hours()/24/30, 'f', 0, 64) + " Months"
	} else {
		distance = strconv.FormatFloat(durasi.Hours()/24/30/12, 'f', 0, 64) + " Years"
	}

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
		Sdate: sDateFormat,
		Edate: eDateFormat,
		Duration: distance,
		Content: content,
		Tnode: node,
		Treact: react,
		Tjs: js,
		Thtml: hateml,
	}

	project = append(project, newAdd)

	// fmt.Println("Project Name : " + r.PostForm.Get("pname"))
	// fmt.Println("Description : " + r.PostForm.Get("desc"))
	// fmt.Println("Start Date : " + r.PostForm.Get("sdate"))
	// fmt.Println("End Date : " + r.PostForm.Get("edate"))
	// fmt.Println("Duration : " + selisih)
	// fmt.Println("Node : " + r.PostForm.Get("nodejs"))
	// fmt.Println("React : " + r.PostForm.Get("reactjs"))
	// fmt.Println("Javascript : " + r.PostForm.Get("js"))
	// fmt.Println("HTML : " + r.PostForm.Get("html"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/projectUpdate.html")
	// error handling
	if err != nil {
		panic(err)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	projectGetUpdate := entities.Project{}
	for i, item := range project {
		ConvertSdate, _ := time.Parse("02 January 2006", item.Sdate)
		ConvertEdate, _ := time.Parse("02 January 2006", item.Edate)
		if i == id {
			projectGetUpdate = entities.Project{
				Title: item.Title,
				Content: item.Content,
				Sdate: ConvertSdate.Format("2006-01-02"),
				Edate: ConvertEdate.Format("2006-01-02"),
				Duration: item.Duration,
				Treact: item.Treact,
				Tnode: item.Tnode,
				Tjs: item.Tjs,
				Thtml: item.Thtml,
			}
		}
	}

	var data = map[string]interface{}{
		"title" : "Edit Project",
		"isLogin" : true,
	}

	resp := map[string]interface{}{
		"Data" : data,
		"GetUpdate" : projectGetUpdate,
		"ID":    id,
	}

	// fmt.Println(projectGetUpdate.Sdate)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}


func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

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

	formatLayout := "2006-01-02"
	sDate, _ := time.Parse(formatLayout, SD)
	sDateFormat := sDate.Format("02 January 2006")
	eDate, _ := time.Parse(formatLayout, ED)
	eDateFormat := eDate.Format("02 January 2006")

	
	durasi := eDate.Sub(sDate)
	// days := int(duration.Hours() / 24)
	// months := int(days / 30)
	// years := int(months / 12)
	var distance string
	if durasi.Hours()/24 < 7 {
		distance = strconv.FormatFloat(durasi.Hours()/24, 'f', 0, 64) + " Days"
	} else if durasi.Hours()/24/7 < 4 {
		distance = strconv.FormatFloat(durasi.Hours()/24/7, 'f', 0, 64) + " Weeks"
	} else if durasi.Hours()/24/30 < 12 {
		distance = strconv.FormatFloat(durasi.Hours()/24/30, 'f', 0, 64) + " Months"
	} else {
		distance = strconv.FormatFloat(durasi.Hours()/24/30/12, 'f', 0, 64) + " Years"
	}

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

	for i := range project {
		update := &project[id]
		if i == id {
			(*update).Title = title
			(*update).Sdate = sDateFormat
			(*update).Edate = eDateFormat
			(*update).Duration = distance
			(*update).Content = content
			(*update).Tnode = node
			(*update).Treact = react
			(*update).Tjs = js
			(*update).Thtml = hateml
		}
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Delete(w http.ResponseWriter, r *http.Request) {
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


	projectDetail := entities.Project{}
	for i, item := range project {
		if i == id {
			projectDetail = entities.Project{
				Title: item.Title,
				Content: item.Content,
				Sdate: item.Sdate,
				Edate: item.Edate,
				Duration: item.Duration,
				Treact: item.Treact,
				Tnode: item.Tnode,
				Tjs: item.Tjs,
				Thtml: item.Thtml,
			}
		}
	}

	var data = map[string]interface{}{
		"title" : "Detail Project",
		"isLogin" : true,
	}

	resp := map[string]interface{}{
		"Data" : data,
		"Projects" : projectDetail,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}