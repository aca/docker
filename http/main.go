package main

import (
	"io"
	"log"
	"net/http"
)

func sleep(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
	log.Println(r.Header)
}

func main() {
	http.HandleFunc("/", sleep)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
