package project

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projek/config"
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
	var tmpl, err = template.ParseFiles("./views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	
	readDT := "SELECT id, project_name, description, start_date, end_date, technologies FROM public.tb_projects;"
	rows, _ := config.ConnDB.Query(context.Background(), readDT)
					
	var result []entities.Project
 for rows.Next() {
	var each = entities.Project{}
 var err = rows.Scan(&each.Id, &each.Title, &each.Content, &each.Sdate, &each.Edate, &each.Technologies)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// duration
	ConvertSdate, _ := time.Parse("02 January 2006", each.Sdate)
	ConvertEdate, _ := time.Parse("02 January 2006", each.Edate)
	duration := ConvertEdate.Sub(ConvertSdate)
	var distance string
	if duration.Hours()/24 < 7 {
		distance = strconv.FormatFloat(duration.Hours()/24, 'f', 0, 64) + " Days"
	} else if duration.Hours()/24/7 < 4 {
		distance = strconv.FormatFloat(duration.Hours()/24/7, 'f', 0, 64) + " Weeks"
	} else if duration.Hours()/24/30 < 12 {
		distance = strconv.FormatFloat(duration.Hours()/24/30, 'f', 0, 64) + " Months"
	} else {
		distance = strconv.FormatFloat(duration.Hours()/24/30/12, 'f', 0, 64) + " Years"
	}
  each.Duration = distance

		result = append(result, each)
	}
		
		var data = map[string]interface{}{
			"title" : "Home | Marcel",
			"isLogin" : true,
		}
		
		resp := map[string]interface{}{
			"Data" : data,
			"Projects" : result,
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
	tech := r.Form["check"]

	// convert date ke string
	sDate, _ := time.Parse("2006-01-02", SD)
	sDateFormat := sDate.Format("02 January 2006")
	eDate, _ := time.Parse("2006-01-02", ED)
	eDateFormat := eDate.Format("02 January 2006")

	addID := "INSERT INTO tb_projects(project_name, start_date, end_date, description, technologies) VALUES ($1, $2, $3, $4, $5, $6)"
	config.ConnDB.Exec(context.Background(), addID, title, sDateFormat, eDateFormat, content, tech )

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/projectUpdate.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	selectID := "SELECT id, project_name, start_date, end_date, description, technologies FROM tb_projects WHERE id=$1"
	rows := config.ConnDB.QueryRow(context.Background(), selectID, id)
	var getID entities.Project
	err = rows.Scan(&getID.Id, &getID.Title, &getID.Sdate, &getID.Edate, &getID.Content, &getID.Technologies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// convert string ke date
	ConvertSdate, _ := time.Parse("02 January 2006", getID.Sdate)
	ConvertEdate, _ := time.Parse("02 January 2006", getID.Edate)
	SD := ConvertSdate.Format("2006-01-02")
	ED := ConvertEdate.Format("2006-01-02")

	result := []entities.Project{getID}

	var data = map[string]interface{}{
		"title" : "Edit Project",
		"isLogin" : true,
	}

	resp := map[string]interface{}{
		"Data" : data,
		"GetUpdate" : result,
		"Sdate" : SD,
		"Edate" : ED,
	}

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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	deleteID := "DELETE FROM tb_projects WHERE id=$1"
	_, err := config.ConnDB.Exec(context.Background(), deleteID, id) 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}



func Detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/projectDetail.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	
	// manangkap id (id, _ (tanda _ tidak ingin menampilkan eror))
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	selectID := "SELECT id, project_name, start_date, end_date, description, technologies FROM tb_projects WHERE id=$1"
	projectDetail := entities.Project{}
	rows := config.ConnDB.QueryRow(context.Background(), selectID, id)
	err = rows.Scan(&projectDetail.Id, &projectDetail.Title, &projectDetail.Sdate, &projectDetail.Edate, &projectDetail.Content, &projectDetail.Technologies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// duration
	ConvertSdate, _ := time.Parse("02 January 2006", projectDetail.Sdate)
	ConvertEdate, _ := time.Parse("02 January 2006", projectDetail.Edate)
	duration := ConvertEdate.Sub(ConvertSdate)
	var distance string
	if duration.Hours()/24 < 7 {
		distance = strconv.FormatFloat(duration.Hours()/24, 'f', 0, 64) + " Days"
	} else if duration.Hours()/24/7 < 4 {
		distance = strconv.FormatFloat(duration.Hours()/24/7, 'f', 0, 64) + " Weeks"
	} else if duration.Hours()/24/30 < 12 {
		distance = strconv.FormatFloat(duration.Hours()/24/30, 'f', 0, 64) + " Months"
	} else {
		distance = strconv.FormatFloat(duration.Hours()/24/30/12, 'f', 0, 64) + " Years"
	}

	var data = map[string]interface{}{
		"title" : "Detail Project",
		"isLogin" : true,
	}

	resp := map[string]interface{}{
		"Data" : data,
		"Projects" : projectDetail,
		"Duration" : distance,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}