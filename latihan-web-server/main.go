package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called ....")
	indexPage := "Ini adalah halaman index"
	w.Write([]byte(indexPage)) // conv dari str ke slice of byte
	}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About called ....")
	aboutPage := "Ini adalah halaman about"
	w.Write([]byte(aboutPage)) // conv dari str ke slice of byte
	}

	// Routing
	func handleRoute()  {
		http.HandleFunc("/", index)
		http.HandleFunc("/about", about)
	}

	// Buat server
	func startServer(port string)  {
		handleRoute()
		fmt.Println("Server running at port ", port)
		http.ListenAndServe(port, nil)
	}


func main() {
	port := ":9999"
	startServer(port)
}