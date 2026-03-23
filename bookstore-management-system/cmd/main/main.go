package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server at port 8083...")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}