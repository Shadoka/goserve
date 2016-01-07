package main

import (
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Path[1:]
	
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}