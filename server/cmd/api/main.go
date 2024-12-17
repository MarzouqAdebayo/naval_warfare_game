package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
