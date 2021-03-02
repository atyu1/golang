package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/indexFunc.tmpl"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "indexFunc.tmpl",  time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
