package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	g "server/internal/game"
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
	id       string
	hub      *Hub
	gameroom *GameRoom
	conn     *websocket.Conn
	send     chan []byte
	mu       sync.Mutex
	isAlive  bool
	userData map[string]interface{} // Store custom user data
	lastPing time.Time
}

// GameRoom represents the BattleshipGame with extra info
type GameRoom struct {
	clients []*Client
	g.BattleshipGame
}

// Hub manages all connected clients
type Hub struct {
	lastID     int
	clients    map[string]*Client
	rooms      []*GameRoom
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.RWMutex
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
		lastID:     0,
		clients:    make(map[string]*Client),
		rooms:      make([]*GameRoom, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 256),
	}

	hub.config.pingInterval = 30 * time.Second
	hub.config.pongWait = 60 * time.Second
	hub.config.maxMessageSize = 512 * 1024 // 512KB
	hub.config.writeWait = 10 * time.Second

	return hub
}

func (g *GameRoom) Broadcast(message []byte) {
	switch "" {
	case "start_game":
		return
	case "attack":
		return
	case "game_over":
		return
	}
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

func (h *Hub) None() {}

func (h *Hub) Run() {
	go h.periodicCleanup()

	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.id] = client
			h.mu.Unlock()

			msg := Message{
				Type:    "user_deets",
				Payload: fmt.Sprintf("User-%s", client.id),
			}
			if msgBytes, err := json.Marshal(msg); err == nil {
				client.conn.WriteMessage(websocket.TextMessage, msgBytes)
			}

			msg = Message{
				Type:    "system",
				Payload: fmt.Sprintf("User %s joined", client.id),
			}
			if msgBytes, err := json.Marshal(msg); err == nil {
				h.broadcast <- msgBytes
			}

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
				close(client.send)

				msg := Message{
					Type:    "system",
					Payload: fmt.Sprintf("User %s left", client.id),
				}
				if msgBytes, err := json.Marshal(msg); err == nil {
					h.broadcast <- msgBytes
				}
			}
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
			if !client.isAlive || time.Since(client.lastPing) > h.config.pongWait {
				log.Printf("Cleaning up inactive client: %s", id)
				h.unregister <- client
			}
		}
		h.mu.Unlock()
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(c.hub.config.maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(c.hub.config.pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.mu.Lock()
		c.lastPing = time.Now()
		c.mu.Unlock()
		c.conn.SetReadDeadline(time.Now().Add(c.hub.config.pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
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
			if targetClient, ok := c.hub.clients[msg.To]; ok {
				msg.From = c.id
				if msgBytes, err := json.Marshal(msg); err == nil {
					targetClient.send <- msgBytes
				}
			}
		case "find_game":
			for _, room := range c.hub.rooms {
				switch len(room.clients) {
				case 0:
					room.clients = append(room.clients, c)
					break
				case 1:
					room.clients = append(room.clients, c)
					room.BattleshipGame = *g.NewBattleshipGame(10)
					room.Broadcast(message)
					break
				default:
					newRoom := &GameRoom{
						clients: make([]*Client, 0),
					}
					c.hub.rooms = append(c.hub.rooms, newRoom)
				}
			}
		default:
			msg.From = c.id
			if msgBytes, err := json.Marshal(msg); err == nil {
				c.hub.broadcast <- msgBytes
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(c.hub.config.pingInterval)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(c.hub.config.writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(c.hub.config.writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
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
		id:       fmt.Sprintf("%d", hub.lastID),
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
		isAlive:  true,
		userData: make(map[string]interface{}),
		lastPing: time.Now(),
	}
	hub.lastID++
	hub.mu.Unlock()

	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}
