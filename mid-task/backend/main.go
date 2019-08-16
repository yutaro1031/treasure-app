package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hgsoe")
	fmt.Fprintf(w, "Hello astax")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
