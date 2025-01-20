package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	db "server/internal/db"
	ws "server/internal/ws"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	database, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	db.Migrate(database)

	apiHandler(ctx)

	port := ":5000"
	log.Printf("Server starting on %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func apiHandler(ctx context.Context) {
	hub := ws.NewHub()
	go hub.Run()

	http.Handle("/", http.FileServer(http.Dir("../../views/")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w)
	})
}
