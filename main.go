package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Server, Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
