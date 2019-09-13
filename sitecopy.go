package main

import (
	"fmt"
	"net/http"
	
)
func redirect(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "http://www.google.com",101)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Online"))
}


func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/google", redirect)
	fmt.Println("listening to port ")
	err := (http.ListenAndServe(":1234", nil))
	if err != nil {
		panic(err)
	}
}


