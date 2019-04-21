package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request")
	fmt.Fprintf(w, "<h1>Hello, I am app2!</h1> Developed using <b>Golang</b>")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

