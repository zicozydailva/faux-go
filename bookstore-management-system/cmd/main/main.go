package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/zicozydailva/faux-go/bookstore-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port 8083...")
	if err := http.ListenAndServe(":8083", r); err != nil {
		log.Fatal(err)
	}
}