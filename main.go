package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World! This is an update"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Missing env variable 'PORT'")
	}

	log.Printf("Listening on port %s", port)
	http.ListenAndServe(":"+port, nil)
}
