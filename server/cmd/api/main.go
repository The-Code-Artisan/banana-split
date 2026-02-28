package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /ping", ping)

	server := http.Server{
		Addr:    ":8001",
		Handler: router,
	}
	log.Println("Starting server on port :8001")
	server.ListenAndServe()
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!")
}
