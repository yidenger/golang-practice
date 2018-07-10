package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header["Accept-Encoding"])
	fmt.Fprintf(w, "World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler function called")
		h(w, r)
	}
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintf(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:3080",
	}

	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", world)

	http.HandleFunc("/body", body)

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}