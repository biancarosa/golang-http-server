package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Serving http...")
	http.ListenAndServe(":8000", nil)
}
