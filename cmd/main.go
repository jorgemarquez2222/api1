package main

import (
	"log"
	"net/http"

	"github.com/jorgemarquez2222/api1/handler"
	"github.com/jorgemarquez2222/api1/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)
	println("runing on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Printf("error: %v", err)
	}
}
