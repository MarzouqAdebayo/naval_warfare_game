package wsimpl

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type GameRoom struct {
	players    map[*Player]bool
	broadcast  chan []byte
	register   chan *Player
	unregister chan *Player
	mutex      sync.RWMutex
	maxPlayers int
}

type Player struct {
	conn     *websocket.Conn
	room     *GameRoom
	send     chan []byte
	username string
}

type GameServer struct {
	rooms    map[string]*GameRoom
	mutex    sync.RWMutex
	upgrader websocket.Upgrader
}

func NewGameServer() *GameServer {
	return &GameServer{
		rooms: make(map[string]*GameRoom),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// TODO In production, implement proper origin checking
				return true
			},
		},
	}
}

func newGameRoom() *GameRoom {
	return &GameRoom{
		players:    make(map[*Player]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		maxPlayers: 2,
	}
}

func (room *GameRoom) run() {
	for {
		select {
		case player := <-room.register:
			room.mutex.Lock()
			if len(room.players) < room.maxPlayers {
				room.players[player] = true
			}
			room.mutex.Unlock()

		case player := <-room.unregister:
			room.mutex.Lock()
			if _, ok := room.players[player]; ok {
				delete(room.players, player)
				close(player.send)
			}
			room.mutex.Unlock()

		case message := <-room.broadcast:
			room.mutex.RLock()
			for player := range room.players {
				select {
				case player.send <- message:
				default:
					close(player.send)
					delete(room.players, player)
				}
			}
			room.mutex.RUnlock()
		}
	}
}

func (player *Player) readPump() {
	defer func() {
		player.room.unregister <- player
		player.conn.Close()
	}()

	for {
		_, message, err := player.conn.ReadMessage()
		if err != nil {
			break
		}
		player.room.broadcast <- message
	}
}

func (player *Player) writePump() {
	defer player.conn.Close()

	for {
		select {
		case message, ok := <-player.send:
			if !ok {
				player.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := player.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}

func (server *GameServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	roomID := r.URL.Query().Get("room")

	server.mutex.Lock()
	room, exists := server.rooms[roomID]
	if !exists {
		room = newGameRoom()
		server.rooms[roomID] = room
		go room.run()
	}
	server.mutex.Unlock()

	player := &Player{
		conn: conn,
		room: room,
		send: make(chan []byte, 256),
	}

	room.register <- player

	go player.writePump()
	go player.readPump()
}
