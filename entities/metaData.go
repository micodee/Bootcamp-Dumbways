package entities

type MetaData struct {
	Title     string
	IsLogin   bool
	UserName  string
	FlashData string
}

var Data = map[string]interface{}{
	"Title" : "Home | Marcel",
	"IsLogin" : true,
	"Id" : 1,
	"UserName" : "Tommy",
}