package entities

type Project struct {
	Id      int
	Title  string
	Sdate  string
	Edate  string
	Duration string
	Content string
	Technologies  []bool `json:"Technologies"`
}
