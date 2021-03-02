//Not the best solution see another projects
package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "<h1>WUH WUH</h1>")
	case "/cat":
		io.WriteString(res, "<h1>MIAU MIAU</h1>")
	default:
		io.WriteString(res, "Please specify Cat or Dog in URL")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
