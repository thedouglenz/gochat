package main

import (
	"fmt"
	"net/http"
)

type Message struct {
	Sender  string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
