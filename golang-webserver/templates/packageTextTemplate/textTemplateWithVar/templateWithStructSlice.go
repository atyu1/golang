package main

import (
	"log"
	"os"
	"text/template"
)

type Gender int

const (
	male = iota
	female
)

type Person struct {
	Name string
	Age  int
	Sex  Gender
}

func main() {
	tpl, err := template.ParseFiles("templates/indexStructSlice.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	p1 := Person{
		Name: "Mike Cornel",
		Age:  31,
		Sex:  male,
	}

	p2 := Person{
		Name: "Peter Golang",
		Age:  35,
		Sex:  male,
	}

	p3 := Person{
		Name: "Rachel Malinksy",
		Age:  22,
		Sex:  female,
	}

	data := []Person{p1, p2, p3}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
