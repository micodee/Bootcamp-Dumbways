package home

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

		// parsing template html
	var tmpl, err = template.ParseFiles("./views/index.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
	
					rows, _ := config.ConnDB.Query(context.Background(), "SELECT id, project_name, description, start_date, end_date, duration FROM public.tb_projects;")
					
				var result []entities.Project

				for rows.Next() {
					var each = entities.Project{}

					var err = rows.Scan(&each.Id, &each.Title, &each.Content, &each.Sdate, &each.Edate, &each.Duration)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					
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

	addID := "INSERT INTO tb_projects(project_name, start_date, end_date, duration, description) VALUES ($1, $2, $3, $4, $5)"
	config.ConnDB.Exec(context.Background(), addID, title, sDateFormat, eDateFormat, distance, content)

	// // fmt.Println("Project Name : " + r.PostForm.Get("pname"))
	// // fmt.Println("Description : " + r.PostForm.Get("desc"))
	// // fmt.Println("Start Date : " + r.PostForm.Get("sdate"))
	// // fmt.Println("End Date : " + r.PostForm.Get("edate"))
	// // fmt.Println("Duration : " + selisih)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/projectUpdate.html")
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
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

	selectDB := "SELECT id, project_name, start_date, end_date, duration, description FROM tb_projects WHERE id=$1"
	projectDetail := entities.Project{}
	err = config.ConnDB.QueryRow(context.Background(), selectDB, id).Scan(&projectDetail.Id, &projectDetail.Title, &projectDetail.Sdate, &projectDetail.Edate, &projectDetail.Duration, &projectDetail.Content)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
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