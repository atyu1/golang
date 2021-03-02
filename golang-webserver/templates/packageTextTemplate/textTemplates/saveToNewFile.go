package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.txt")
	if err != nil {
		log.Fatalln(err)
	}

  // ParseFiles can be added here if we want to join multiple files to 1 template
  // After we can use ExecuteTeamplate(tplVar, <template file name>, variables)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(nf, nil) // Save template + variables updated(now nil) to index.DOCTYPEhtml
	if err != nil {
		log.Fatalln(err)
	}
}
