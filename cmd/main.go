package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "KCD")
}

func main() {
	http.HandleFunc("/", index)

	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
