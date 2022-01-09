package main

import "net/http"

func main() {
	http.HandleFunc("/saludar", saludar)
	println("Listening in port http://localhost:8080/saludar ")
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}
