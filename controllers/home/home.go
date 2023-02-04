package home

import (
	"html/template"
	"net/http"
	"projek/entities"
)

func Home(w http.ResponseWriter, r *http.Request) {
		// parsing template html
	var tmpl, err = template.ParseFiles("./views/index.html")
	if err != nil {
		panic(err)
	}

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
	
var data = map[string]interface{}{
	"title" : "Personal Web",
	"isLogin" : true,
}

	resp := map[string]interface{}{
		"Data" : data,
		"Projects" : project,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}