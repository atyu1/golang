package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templates/indexSlice.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	data := []string{"Mike", "Rachel", "Gerry", "Joey", "Peter"}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
