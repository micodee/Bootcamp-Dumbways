package entities

type Project struct {
	Id      int
	Title  string
	Sdate  string
	Edate  string
	Duration string
	Content string
	Technologies []string
	Tnode   bool
	Treact  bool
	Tjs     bool
	Thtml   bool
}
