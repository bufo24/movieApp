package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", first)
	http.ListenAndServe(":8090", nil)
}

func first(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "first\n")
}