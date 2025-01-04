package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Message represents the structure of WebSocket messages
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	From    string      `json:"from"`
	To      string      `json:"to,omitempty"`
}

// Client represents a connected WebSocket client
type Client struct {
	ID       string
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
	mu       sync.Mutex
	IsAlive  bool
	UserData map[string]interface{} // Store custom user data
	LastPing time.Time
}

// Hub manages all connected clients
type Hub struct {
	LastID     int
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	mu         sync.RWMutex
	// Config
	Config struct {
		PingInterval   time.Duration
		PongWait       time.Duration
		MaxMessageSize int64
		WriteWait      time.Duration
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
		LastID:     0,
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte, 256),
	}

	hub.Config.PingInterval = 30 * time.Second
	hub.Config.PongWait = 60 * time.Second
	hub.Config.MaxMessageSize = 512 * 1024 // 512KB
	hub.Config.WriteWait = 10 * time.Second

	return hub
}

func (h *Hub) Run() {
	go h.periodicCleanup()

	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.ID] = client
			h.mu.Unlock()

			msg := Message{
				Type:    "system",
				Payload: fmt.Sprintf("User %s joined", client.ID),
			}
			if msgBytes, err := json.Marshal(msg); err == nil {
				h.Broadcast <- msgBytes
			}

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)

				msg := Message{
					Type:    "system",
					Payload: fmt.Sprintf("User %s left", client.ID),
				}
				if msgBytes, err := json.Marshal(msg); err == nil {
					h.Broadcast <- msgBytes
				}
			}
			h.mu.Unlock()

		case message := <-h.Broadcast:
			h.mu.RLock()
			for _, client := range h.Clients {
				if !client.IsAlive {
					continue
				}
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					client.IsAlive = false
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) periodicCleanup() {
	ticker := time.NewTicker(h.Config.PingInterval)
	defer ticker.Stop()

	for range ticker.C {
		h.mu.Lock()
		for id, client := range h.Clients {
			if !client.IsAlive || time.Since(client.LastPing) > h.Config.PongWait {
				log.Printf("Cleaning up inactive client: %s", id)
				h.Unregister <- client
			}
		}
		h.mu.Unlock()
	}
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(c.Hub.Config.MaxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(c.Hub.Config.PongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.mu.Lock()
		c.LastPing = time.Now()
		c.mu.Unlock()
		c.Conn.SetReadDeadline(time.Now().Add(c.Hub.Config.PongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("error parsing message: %v", err)
			continue
		}

		switch msg.Type {
		case "private":
			if targetClient, ok := c.Hub.Clients[msg.To]; ok {
				msg.From = c.ID
				if msgBytes, err := json.Marshal(msg); err == nil {
					targetClient.Send <- msgBytes
				}
			}
		default:
			msg.From = c.ID
			if msgBytes, err := json.Marshal(msg); err == nil {
				c.Hub.Broadcast <- msgBytes
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(c.Hub.Config.PingInterval)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(c.Hub.Config.WriteWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(c.Hub.Config.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	hub.mu.Lock()
	client := &Client{
		ID:       fmt.Sprintf("%d", hub.LastID),
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
		IsAlive:  true,
		UserData: make(map[string]interface{}),
		LastPing: time.Now(),
	}
	hub.LastID++
	hub.mu.Unlock()

	client.Hub.Register <- client

	go client.readPump()
	go client.writePump()
}
