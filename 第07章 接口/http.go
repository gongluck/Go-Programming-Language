package main

import (
	"fmt"
	"log"
	"net/http"
)

type ServeHandle int

func (h ServeHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path)
}

func main() {
	var h ServeHandle
	log.Fatal(http.ListenAndServe("localhost:8000", h))
}
