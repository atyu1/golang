//The best solution is to use HandleFunc so we don't need to create own type
//Easier to write and cleaner code
package main

import (
  "io"
  "net/http"
)

func d(req http.ResponseWriter, req *http.Request) {
  io.WriteString(res, "Doggy dog")
}

func c(req http.ResponseWriter, req *http.Request) {
  io.WriteString(res, "Cat miau")
}

func main() {
  // Trailing slash will catch every URL starting with /dog/...
  // Using new HandleFunc
  http.HandleFunc("/dog/", d)
  http.HandleFunc("/cat", c)

  http.ListenAndServe(":8080", nil)  //Default using nil

}
