//Basic Mutex, still not the best option
package main

import (
  "io"
  "net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(req http.ResponseWriter, req *http.Request) {
  io.WriteString(res, "Doggy dog")
}

type hotcat int

func (d hotcat) ServeHTTP(req http.ResponseWriter, req *http.Request) {
  io.WriteString(res, "Cat miau")
}

func main() {
  var d hotdog
  var c hotcat

  mux := http.NewServeMux()
  // Trailing slash will catch every URL starting with /dog/...
  mux.Handle("/dog/", d)
  mux.Handle("/cat", c)

  http.ListenAndServe(":8080", mux)

}
