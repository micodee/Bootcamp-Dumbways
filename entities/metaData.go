package entities

type MetaData struct {
	Title     string
	IsLogin   bool
	Username  string
	FlashData string
}

var Data = MetaData{
	Title: "Personal Web",
}