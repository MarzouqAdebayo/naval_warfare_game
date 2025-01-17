package ws

import (
	"encoding/json"
	"fmt"

	g "server/internal/game"
)

// EventType represent the type of events
type EventType string

// EventType definitions
const (
	// Incoming Events
	EventSetUserData EventType = "set_user_data"
	EventAttack      EventType = "attack"
	EventFindGame    EventType = "find_game"
	EventQuitGame    EventType = "quit_game"
	EventPlaceShip   EventType = "place_ships"

	// Outgoing Events
	EventFindGameWaiting    EventType = "find_game_waiting"
	EventFindGameStart      EventType = "find_game_start"
	EventShipRandomized     EventType = "randomized_place_ship_response"
	EventBroadcastAttack    EventType = "broadcast_attack"
	EventPong               EventType = "pong"
	EventOpponentQuit       EventType = "opponent_quit"
	EventClientDisconnected EventType = "client_disconnected"
)

// BaseEvent represents a common structure for all other events
type BaseEvent struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// SetUserDataPayload respresents the payload for user data event
type SetUserDataPayload struct {
	Username string `json:"name"`
}

// FindGamePayload represents the payload for find_game event
type FindGameStartPayload struct {
	RoomID      string        `json:"roomID"`
	Index       int           `json:"index"`
	Players     [2]PlayerData `json:"players"`
	Message     string        `json:"message"`
	CurrentTurn int           `json:"currentTurn"`
	GameOver    bool          `json:"gameOver"`
	Mode        g.GameMode    `json:"mode"`
	Winner      int           `json:"winner"`
	Status      g.GameStatus  `json:"status"`
}

// FindGamePayload represents the payload for find_game event
type FindGameWaitingPayload struct {
	RoomID string `json:"roomID"`
}

// AttackPayload respresents the payload for attack event
type AttackPayload struct {
	RoomID         string     `json:"roomID"`
	AttackerIndex  int        `json:"attackerIndex"`
	AttackPosition g.Position `json:"attackPosition"`
}

// PlaceShipPayload represents the payload for placing_ship event
type PlaceShipPayload struct {
	Instruction g.PlaceInstruction `json:"instruction"`
	RoomID      string             `json:"roomID"`
	PlayerIndex int                `json:"playerIndex"`
}

// ShipRandomizedPayload respresents the payload for placing ships randomly
type ShipRandomizedPayload struct {
	Message    string     `json:"message"`
	PlayerData PlayerData `json:"playerData"`
}

// AttackBroadcastPayload represents the payload for an outgoing attack event
type BroadcastAttackPayload struct {
	RoomID      string        `json:"roomID"`
	Index       int           `json:"index"`
	Players     [2]PlayerData `json:"players"`
	Message     string        `json:"message"`
	CurrentTurn int           `json:"currentTurn"`
	GameOver    bool          `json:"gameOver"`
	Mode        g.GameMode    `json:"mode"`
	Winner      int           `json:"winner"`
}

// EventPongPayload
type EventPongPayload struct {
	NoOfClients int `json:"noOfClients"`
	NoOfRooms   int `json:"noOfRooms"`
}

// Event interface
type Event interface {
	GetType() EventType
	GetPayload() interface{}
}

// Concrete event types
type SetUserDataEvent struct {
	Type    EventType          `json:"type"`
	Payload SetUserDataPayload `json:"payload"`
}

type FindGameEvent struct {
	Type EventType `json:"type"`
}

type FindGameWaitingEvent struct {
	Type    EventType              `json:"type"`
	Payload FindGameWaitingPayload `json:"payload"`
}

type FindGameStartEvent struct {
	Type    EventType            `json:"type"`
	Payload FindGameStartPayload `json:"payload"`
}

type QuitGameEvent struct {
	Type    EventType `json:"type"`
	Payload string    `json:"payload"`
}

type PlaceShipEvent struct {
	Type    EventType        `json:"type"`
	Payload PlaceShipPayload `json:"payload"`
}

type ShipRandomizedEvent struct {
	Type    EventType             `json:"type"`
	Payload ShipRandomizedPayload `json:"payload"`
}

type AttackEvent struct {
	Type    EventType     `json:"type"`
	Payload AttackPayload `json:"payload"`
}

type BroadcastAttackEvent struct {
	Type    EventType              `json:"type"`
	Payload BroadcastAttackPayload `json:"payload"`
}

type PongEvent struct {
	Type    EventType        `json:"type"`
	Payload EventPongPayload `json:"payload"`
}

type ClientDisconnectedEvent struct {
	Type    EventType `json:"type"`
	Payload string    `json:"payload"`
}

type PlayerData struct {
	Board [][]g.CellState `json:"board"`
	Fleet g.PlayerFleet   `json:"fleet"`
}

type GameStartPayload struct {
	RoomID      string        `json:"roomID"`
	Index       int           `json:"index"`
	Players     [2]PlayerData `json:"players"`
	Message     string        `json:"message"`
	CurrentTurn int           `json:"currentTurn"`
	GameOver    bool          `json:"gameOver"`
	Mode        g.GameMode    `json:"mode"`
	Winner      int           `json:"winner"`
}

// Helper methods for each event type
func (e *SetUserDataEvent) GetType() EventType {
	return EventSetUserData
}

func (e *SetUserDataEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *FindGameEvent) GetType() EventType {
	return EventFindGame
}

func (e *FindGameEvent) GetPayload() interface{} {
	return struct{}{}
}

func (e *QuitGameEvent) GetType() EventType {
	return EventQuitGame
}

func (e *QuitGameEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *PlaceShipEvent) GetType() EventType {
	return EventPlaceShip
}

func (e *PlaceShipEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *AttackEvent) GetType() EventType {
	return EventAttack
}

func (e *AttackEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *BroadcastAttackEvent) GetType() EventType {
	return EventBroadcastAttack
}

func (e *BroadcastAttackEvent) GetPayload() interface{} {
	return e.Payload
}

// ParseEvent parses a raw event into its specific type
func ParseEvent(data []byte) (Event, error) {
	var base BaseEvent
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, fmt.Errorf("failed to parse base event: %w", err)
	}

	switch base.Type {
	case EventSetUserData:
		var event SetUserDataEvent
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v", err)
			return nil, fmt.Errorf("failed to parse SetUserData event: %w", err)
		}
		return &event, nil
	case EventFindGame:
		var event FindGameEvent
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v", err)
			return nil, fmt.Errorf("failed to parse FindGameEvent event: %w", err)
		}
		return &event, nil
	case EventPlaceShip:
		var event PlaceShipEvent
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v", err)
			return nil, fmt.Errorf("failed to parse ShipRandomizedEvent event: %w", err)
		}
		return &event, nil
	case EventQuitGame:
		var event QuitGameEvent
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v", err)
			return nil, fmt.Errorf("failed to parse QuitGameEvent event: %w", err)
		}
		return &event, nil
	case EventAttack:
		var event AttackEvent
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("%v", err)
			return nil, fmt.Errorf("failed to parse Attack event: %w", err)
		}
		return &event, nil
	}
	return nil, fmt.Errorf("unknown event type: %s", base.Type)
}

func SetUserDataEventHandler(e Event, c *Client) {
	c.hub.mu.Lock()
	defer c.hub.mu.Unlock()
	if userEvent, ok := e.(*SetUserDataEvent); ok {
		c.userData["name"] = userEvent.Payload.Username
	}
}

func PlaceShipEventHandler(e Event, c *Client) {
	placeEvent, ok := e.(*PlaceShipEvent)
	if !ok {
		return
	}
	roomID := placeEvent.Payload.RoomID
	room, ok := c.hub.rooms[roomID]
	if !ok {
		fmt.Println("room not found")
		return
	}
	playerIndex := placeEvent.Payload.PlayerIndex
	if playerIndex+1 > len(room.Players) {
		fmt.Println("index invalid")
		return
	}
	playerClient := room.clients[playerIndex]
	if playerClient != c {
		fmt.Println("client mismatch")
		return
	}
	player := room.Players[playerIndex]
	instr := placeEvent.Payload.Instruction
	if instr == g.Randomize {
		player.GenerateAndPlaceShips()

		message := "Hi Captain, our fleet is ready"
		outgoingEvent := ShipRandomizedEvent{
			Type: EventShipRandomized,
			Payload: ShipRandomizedPayload{
				Message: message,
				PlayerData: PlayerData{
					Board: player.Board.PlainBoard(),
					Fleet: player.GetPlainFleetInfo(),
				},
			},
		}

		b, err := json.Marshal(outgoingEvent)
		if err != nil {
			fmt.Printf("failed to marshal FindGameWaitingEvent: %v", err)
			return
		}
		c.send <- b
	}
}

func QuitGameEventHandler(e Event, c *Client) {
	roomEvent, ok := e.(*QuitGameEvent)
	if !ok {
		return
	}
	roomID := roomEvent.Payload
	evt := ClientDisconnectedEvent{
		Type:    EventClientDisconnected,
		Payload: fmt.Sprintf("%s left", c.userData["name"]),
	}
	if b, err := json.Marshal(evt); err == nil {
		c.send <- b
	}
	delete(c.hub.rooms, roomID)
}

func FindGameEventHandler(c *Client) {
	fmt.Println("New find game request")
	c.hub.mu.Lock()
	defer c.hub.mu.Unlock()
	room := findAvailableRoom(c.hub.rooms)

	if room != nil {
		room.clients = append(room.clients, c)
		room.BattleshipGame = *g.NewBattleshipGame(10)

		// for i := range room.clients {
		// 	player := room.BattleshipGame.Players[i]
		// 	player.GenerateAndPlaceShips()
		// }

		message := "Hi Captain, you've been given orders to eliminate the enemy, good luck sailor"
		for i, c := range room.clients {
			outgoingEvent := FindGameStartEvent{
				Type: EventFindGameStart,
				Payload: FindGameStartPayload{
					RoomID:      room.id,
					Index:       i,
					Players:     [2]PlayerData{},
					Message:     message,
					CurrentTurn: room.CurrentTurn,
					GameOver:    room.GameOver,
					Mode:        room.Mode,
					Winner:      room.CurrentTurn,
					Status:      g.Ready,
				},
			}

			// Send plain board for player
			player1 := room.Players[i]
			outgoingEvent.Payload.Players[i] = PlayerData{Board: player1.Board.PlainBoard(), Fleet: player1.GetPlainFleetInfo()}

			// Send masked board for opponent
			player2 := room.Players[1-i]
			outgoingEvent.Payload.Players[1-i] = PlayerData{Board: player2.Board.MaskBoard(), Fleet: player2.GetMaskedFleetInfo()}

			b, err := json.Marshal(outgoingEvent)
			if err != nil {
				fmt.Printf("failed to marshal FindGameWaitingEvent: %v", err)
				return
			}
			c.send <- b
		}

	} else {
		newRoomID := fmt.Sprintf("room-%d", c.hub.lastRoomID)
		newRoom := &GameRoom{
			id:      newRoomID,
			clients: make([]*Client, 0),
		}
		newRoom.clients = append(newRoom.clients, c)
		c.hub.rooms[newRoomID] = newRoom
		c.hub.lastRoomID++

		outgoingEvent := FindGameWaitingEvent{
			Type: EventFindGameWaiting,
			Payload: FindGameWaitingPayload{
				RoomID: newRoomID,
			},
		}
		b, err := json.Marshal(outgoingEvent)
		if err != nil {
			fmt.Printf("failed to marshal FindGameWaitingEvent: %v", err)
			return
		}
		c.send <- b
	}
}

func AttackEventHandler(e Event, c *Client) {
	// c.hub.mu.Lock()
	// defer c.hub.mu.Unlock()
	listRooms(c.hub)
	attackEvent, ok := e.(*AttackEvent)
	if !ok {
		return
	}
	room, ok := c.hub.rooms[attackEvent.Payload.RoomID]
	if !ok {
		return
	}
	hit, err := room.Attack(attackEvent.Payload.AttackerIndex, attackEvent.Payload.AttackPosition)
	if err != nil {
		return
	}
	if !hit {
	}
	message := "Hi Hi Captain"
	for i, c := range room.clients {
		outgoingEvent := BroadcastAttackEvent{
			Type: EventBroadcastAttack,
			Payload: BroadcastAttackPayload{
				RoomID:      room.id,
				Index:       i,
				Players:     [2]PlayerData{},
				Message:     message,
				CurrentTurn: room.CurrentTurn,
				GameOver:    room.GameOver,
				Mode:        room.Mode,
				Winner:      room.CurrentTurn,
			},
		}
		// Send plain board for player
		player1 := room.Players[i]
		outgoingEvent.Payload.Players[i] = PlayerData{Board: player1.Board.PlainBoard(), Fleet: player1.GetPlainFleetInfo()}
		// Send masked board for opponent
		player2 := room.Players[1-i]
		outgoingEvent.Payload.Players[1-i] = PlayerData{Board: player2.Board.MaskBoard(), Fleet: player2.GetMaskedFleetInfo()}
		b, err := json.Marshal(outgoingEvent)
		if err != nil {
			fmt.Printf("failed to marshal BroadcastAttackEvent: %v", err)
			return
		}
		c.send <- b
	}
}

// Finds a room that is not full
func findAvailableRoom(rooms map[string]*GameRoom) *GameRoom {
	for _, room := range rooms {
		if len(room.clients) == 2 {
			continue
		}
		return room
	}
	return nil
}
