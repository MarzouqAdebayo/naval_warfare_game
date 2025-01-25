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

	dbClient, err := db.InitializeDBClient()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	// if err := db.Migrate(dbClient); err != nil {
	// 	log.Fatalf("Error migrating the database: %v", err)
	// }

	apiHandler(ctx, dbClient)

	port := ":5000"
	log.Printf("Server starting on %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func apiHandler(ctx context.Context, dbClient *db.Database) {
	hub := ws.NewHub()
	go hub.Run()

	http.Handle("/", http.FileServer(http.Dir("../../views/")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r, dbClient)
	})

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w)
	})
}
