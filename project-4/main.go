package main

import (
	"net/http"
	"project-4/handlers"
)

func main() {
	http.HandleFunc("/home", handlers.Home)
	//start your server
	http.ListenAndServe(":8081", nil)
}
