package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World from webapp!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Bytes written: ", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
