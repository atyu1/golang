package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("registering handlers...")

	// Handler for /ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	// Handler for /info
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			http.Error(w, "Unable to retrieve hostname", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, hostname)
	})

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
