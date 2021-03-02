package main

import (
	"log"
	"os"
	"text/template"
)

type region string

const (
	Southern region = "Southern"
	Central  region = "Central"
	Northern region = "Northern"
)

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		Hotel{
			Name:    "Casadila",
			Address: "Soth Beach 3",
			City:    "San Francisco",
			Zip:     "30-7201",
			Region:  Northern,
		},
		Hotel{
			Name:    "InkaBas",
			Address: "Main Blvd 93",
			City:    "San Jose",
			Zip:     "9932-701",
			Region:  Northern,
		},
		Hotel{
			Name:    "Margarita",
			Address: "Main square 1",
			City:    "New York",
			Zip:     "31-723301",
			Region:  Central,
		},
		Hotel{
			Name:    "Margarita",
			Address: "New Hames 201",
			City:    "New Mexico",
			Zip:     "3224-2421",
			Region:  Southern,
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
