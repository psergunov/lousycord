package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/build"))

	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server on port", port)
	err := http.ListenAndServeTLS(":"+port, "certs/cert.pem", "certs/key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
