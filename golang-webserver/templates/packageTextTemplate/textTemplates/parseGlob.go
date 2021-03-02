package main

import (
	"os"
  "log"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "foo.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "bar.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
