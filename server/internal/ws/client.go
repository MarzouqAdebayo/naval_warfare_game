package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	db "server/internal/db"
)

// Client represents a connected WebSocket client
type Client struct {
	id       string
	hub      *Hub
	gameroom *GameRoom
	conn     *websocket.Conn
	send     chan []byte
	mu       sync.RWMutex
	isAlive  bool
	userData map[string]interface{} // Store custom user data
	lastPing time.Time
	db	 *db.Database
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
		c.isAlive = true
		c.lastPing = time.Now()
		c.mu.Unlock()
		c.conn.SetReadDeadline(time.Now().Add(c.hub.config.pongWait))
		return nil
	})

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		evt, err := ParseEvent(payload)
		log.Printf("%#+v", evt)
		switch evt.GetType() {
		case EventSetUserData:
			SetUserDataEventHandler(evt, c)
		case EventFindGame:
			FindGameEventHandler(c)
		case EventPlaceShip:
			PlaceShipEventHandler(evt, c)
		case EventShipReady:
			ShipReadyEventHandler(evt, c)
		case EventQuitGame:
			QuitGameEventHandler(evt, c)
		case EventAttack:
			AttackEventHandler(evt, c)
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
