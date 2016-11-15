package main

import (
	"io"
	"net/http"
)

//Handling function takes resposewriter & request
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

//sample Server using GO
func main() {
	http.HandleFunc("/", hello) //http.HandleFunc takes URL and Handling function
	http.ListenAndServe(":8000", nil)
}
