package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called ....")
	hello := "Hello mom!"
	w.Write([]byte(hello)) // conv dari str ke slice of byte
	}

	// Routing
	func handleRoute()  {
		http.HandleFunc("/", index)
	}

	// Buat server
	func startServer(port string)  {
		handleRoute()
		fmt.Println("Server running at port ", port)
		http.ListenAndServe(port, nil)
	}


func main() {
	port := ":8000"
	startServer(port)
}