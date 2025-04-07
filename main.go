package main

import (
	"fmt"
	"log"
	"net/http"
)

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func main() {
	http.HandleFunc("/", aboutPage)
	fmt.Println("Server Started on port: 8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
