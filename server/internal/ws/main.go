package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	hub.mu.Lock()
	client := &Client{
		id:       fmt.Sprintf("%d", hub.lastClientID),
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
		isAlive:  true,
		userData: make(map[string]interface{}),
		lastPing: time.Now(),
	}
	hub.lastClientID++
	hub.mu.Unlock()

	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}
