package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Ping")
	})

	fmt.Printf("Listening on 8001")

	log.Fatal(http.ListenAndServe(":8001", nil))
}
