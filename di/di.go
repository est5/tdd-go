package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(buffer io.Writer, word string) {
	fmt.Fprintf(buffer, "Hello, %s", word)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
