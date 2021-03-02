package main

import (
  "net/http"
  "io"
  "strings"
)

func root(resp http.ResponseWriter, req *http.Request) {
  io.WriteString(resp, "Home is home")
}

func dog(resp http.ResponseWriter, req *http.Request) {
  io.WriteString(resp, "Waff")
}

func me(resp http.ResponseWriter, req *http.Request) {
  io.WriteString(resp, "This is me:")
  url := req.URL.RequestURI()
  name := strings.Split(url, "/")[-1]
  io.WriteString(resp, name)
}

func main() {
  http.HandleFunc("/", root)
  http.HandleFunc("/dog/", dog)
  http.HandleFunc("/me/", me)

  http.ListenAndServe(":8080", nil)

}
