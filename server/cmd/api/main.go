package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	ws "server/internal/ws-impl"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	apiHandler(ctx)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "Hello World\n")
}

func apiHandler(ctx context.Context) {
	server := ws.NewGameServer()

	http.Handle("/", http.FileServer(http.Dir("../../views/")))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ws", server.HandleWebSocket)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w)
	})
}
