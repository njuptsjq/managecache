package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello!")
	})
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))
	http.Handle("/cachepage", &templateHandler{filename: "cachepage.html"})
	http.HandleFunc("/getparam", getparam)

	log.Fatal(http.ListenAndServe(":9502", nil))
}
