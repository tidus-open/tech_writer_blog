package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServeTLS(":9090", "key/cert.pem", "key/key.pem", router))
}
