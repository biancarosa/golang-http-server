package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Message string `json:"msg"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	person := Greeting{Message: "Hello"}
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPerson)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Serving http...")
	http.ListenAndServe(":8000", nil)
}
