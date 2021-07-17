package main

import (
	"fmt"
	"net/http"
	"backend/Database"
)

func main() {
	http.HandleFunc("/api", first)
	http.ListenAndServe(":8090", nil)
}

func first(w http.ResponseWriter, r *http.Request) {
	info := db.Query();
	fmt.Fprintf(w, info)
}