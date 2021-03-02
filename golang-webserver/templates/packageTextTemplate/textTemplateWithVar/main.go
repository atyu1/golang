package main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  tpl, err := template.ParseFiles("templates/indexVar.tmpl")
  if err != nil {
    log.Fatalln(err)
  }

  err = tpl.Execute(os.Stdout, "Wow, this text has been just magically added!")
  if err != nil {
    log.Fatalln(err)
  }
}
