package main

import "net/http"

func main() {
	Handler1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello this is our home page"))
	}
	http.HandleFunc("/home", Handler1)
	//start your server
	http.ListenAndServe(":8080", nil)

}
