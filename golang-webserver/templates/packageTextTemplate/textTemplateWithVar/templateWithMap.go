package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templates/indexMap.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]string{
		"TZ": "TechZone",
		"BZ": "BackZone",
		"LT": "LessThen",
		"GT": "GreaterThen",
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
