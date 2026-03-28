package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var tmpl *template.Template

func main() {
	fmt.Println("Starting server at port 8084...")
	http.HandleFunc("/", handlerHello)
	http.HandleFunc("/hi", handleHi)
	
	http.ListenAndServe(":8084", nil)
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, World!")
}