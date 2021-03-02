//Basic Mutex, still not the best option
//Much better to use defaultMutex
//If we dont create NewServerMux then we use default
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

  // Trailing slash will catch every URL starting with /dog/...
  http.Handle("/dog/", d)
  http.Handle("/cat", c)

  http.ListenAndServe(":8080", nil)  //Default using nil

}
