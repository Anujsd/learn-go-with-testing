package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	err := http.ListenAndServe(":5000", server)
	log.Fatal(err)
}
