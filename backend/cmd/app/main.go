package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	Admin    = "admin"
	Password = "admin"
)

func auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Header.Get("Authorization") == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Server"`)
		http.Error(w, fmt.Sprintf("Error trying to read Basic auth."), http.StatusUnauthorized)
		return
	}

	user, pw, ok := r.BasicAuth()

	if !ok {
		http.Error(w, fmt.Sprintf("Error trying to read Basic auth."), http.StatusBadRequest)
		return
	}

	if (user == Admin) && (pw == Password) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Authorized")
		return
	} else {
		http.Error(w, fmt.Sprintf("Bad credentials"), http.StatusUnauthorized)
		return
	}
}

func main() {
	fs := http.FileServer(http.Dir("./frontend/build"))

	http.Handle("/", fs)

	http.HandleFunc("/login", auth)

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
