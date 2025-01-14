package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Hub manages all connected clients
type Hub struct {
	lastClientID int
	lastRoomID   int
	clients      map[string]*Client
	rooms        map[string]*GameRoom
	register     chan *Client
	unregister   chan *Client
	broadcast    chan []byte
	mu           sync.RWMutex
	// config
	config struct {
		pingInterval   time.Duration
		pongWait       time.Duration
		maxMessageSize int64
		writeWait      time.Duration
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// TODO In production, implement proper origin checking
		return true
	},
}

func NewHub() *Hub {
	hub := &Hub{
		lastClientID: 1,
		lastRoomID:   1,
		clients:      make(map[string]*Client),
		rooms:        make(map[string]*GameRoom),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		broadcast:    make(chan []byte, 256),
	}

	hub.config.pingInterval = 30 * time.Second
	hub.config.pongWait = 60 * time.Second
	hub.config.maxMessageSize = 512 * 1024 // 512KB
	hub.config.writeWait = 10 * time.Second

	return hub
}

func (g *GameRoom) Broadcast(message []byte) {
	for _, client := range g.clients {
		if !client.isAlive {
			continue
		}
		select {
		case client.send <- message:
		default:
			close(client.send)
			client.isAlive = false
		}
	}
}

func (h *Hub) Run() {
	go h.periodicCleanup()

	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.id] = client
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			h.clientDisconnected(client)
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.RLock()
			for _, client := range h.clients {
				if !client.isAlive {
					continue
				}
				select {
				case client.send <- message:
				default:
					close(client.send)
					client.isAlive = false
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) periodicCleanup() {
	ticker := time.NewTicker(h.config.pingInterval)
	defer ticker.Stop()

	for range ticker.C {
		h.mu.Lock()
		for id, client := range h.clients {
			client.mu.RLock()
			if !client.isAlive || time.Since(client.lastPing) > h.config.pongWait {
				log.Printf("Cleaning up inactive client: %s", id)
				h.unregister <- client
			}
			client.mu.RUnlock()
		}
		h.mu.Unlock()
	}
}

func (h *Hub) clientDisconnected(client *Client) {
	if _, ok := h.clients[client.id]; ok {
		for id, room := range h.rooms {
			if !(slices.Contains(room.clients, client)) {
				continue
			}
			for _, c := range room.clients {
				msg := Message{
					Type:    "user_left",
					Payload: fmt.Sprintf("%s left", client.userData["name"]),
				}
				if msgBytes, err := json.Marshal(msg); err == nil {
					c.send <- msgBytes
				}
			}
			delete(h.rooms, id)
		}
		delete(h.clients, client.id)
		close(client.send)

		// FIXME Delete later
		fmt.Println()
		fmt.Printf("%s left, number of clients = %d, number of rooms = %d, ", client.userData["name"], len(h.clients), len(h.rooms))
		for _, room := range h.rooms {
			fmt.Printf("%v has %d clients", room.id, len(room.clients))
			for _, client := range room.clients {
				fmt.Printf(" -- client %s", client.userData["name"])
			}
			fmt.Println()
		}
		// FIXME Delete later
	}
}
