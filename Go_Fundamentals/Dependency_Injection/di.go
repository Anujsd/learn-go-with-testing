package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(buffer io.Writer, name string) {
	fmt.Fprintf(buffer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Anuj")
}

func main() {
	fmt.Printf("Hello, Anuj")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
