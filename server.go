package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8000"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)

	log.Println("Listening on port", addr, "...")
	http.ListenAndServe(addr, nil)
}
