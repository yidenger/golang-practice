package main

import (
	"fmt"
	"net/http"
	"encoding/json"
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

func write(w http.ResponseWriter, r *http.Request) {
	str := `hello golang`

	w.Header().Set("Location", "http://google.com")
	w.Write([]byte(str))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Post struct {
		name string
		age int
	}
	post := &Post{
		name: "xiaoming",
		age: 22,
	}

	jsonData, err := json.Marshal(post)

	fmt.Println(err);

	w.Write(jsonData)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Node.js",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:3080",
	}

	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", world)

	http.HandleFunc("/body", body)

	http.HandleFunc("/process", process)

	http.HandleFunc("/write", write)

	http.HandleFunc("/redirect", redirect)

	http.HandleFunc("/json", jsonExample)

	http.HandleFunc("/set-cookie", setCookie)

	server.ListenAndServe()
}