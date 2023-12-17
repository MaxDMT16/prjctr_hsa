package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
	})

	log.Println("Server started on :80")

	log.Fatal(http.ListenAndServe(":80", http.DefaultServeMux))

	log.Println("Exiting...")
}