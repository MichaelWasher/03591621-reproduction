package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-correlation-id", "my values")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Hello, world!"}`)
}
