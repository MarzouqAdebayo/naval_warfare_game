package ws

import (
	"encoding/json"
	"fmt"
	"log"

	g "server/internal/game"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventSetUserData    = "set_user_data"
	EventFindGame       = "find_game"
	EventAttack         = "attack"
	EventValidPlacement = "valid_placement"
	EventOpponentQuit   = "opponent_quit"
)

type SetUserDataEvent struct {
	Username string `json:"name"`
}

type AttackEvent struct {
	RoomID         string     `json:"roomID"`
	AttackerIndex  int        `json:"attackerIndex"`
	AttackPosition g.Position `json:"attackPosition"`
}

type PlayerInfo struct {
	Board [][]g.CellState `json:"board"`
	Fleet g.PlayerFleet   `json:"fleet"`
}

type GameStartPayload struct {
	RoomID      string        `json:"roomID"`
	Index       int           `json:"index"`
	Players     [2]PlayerInfo `json:"players"`
	Message     string        `json:"message"`
	CurrentTurn int           `json:"currentTurn"`
	GameOver    bool          `json:"gameOver"`
	Mode        g.GameMode    `json:"mode"`
	Winner      int           `json:"winner"`
}

func (e *Event) Marshal() ([]byte, error) {
	d, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event: %v", err)
	}
	return d, nil
}

func (e Event) Unmarshal(data []byte) (*Event, error) {
	var v Event
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event: %v", err)
	}
	return &v, nil
}

// func SendMessageHandler(event Event, h *Hub, c *Client) error {
// 	var chatevent SendMessageEvent
// 	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
// 		return fmt.Errorf("bad payload in request: %v", err)
// 	}
//
// 	var broadMessage NewMessageEvent
//
// 	broadMessage.Sent = time.Now()
// 	broadMessage.Message = chatevent.Message
// 	broadMessage.From = chatevent.From
//
// 	data, err := json.Marshal(broadMessage)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal broadcast message: %v", err)
// 	}
//
// 	// Place payload into an Event
// 	var outgoingEvent Event
// 	outgoingEvent.Payload = data
// 	outgoingEvent.Type = EventBroadcastAttack
// 	// Broadcast to all other Clients
// 	for _, client := range c.hub.clients {
// 		// Only send to clients inside the same chatroom
// 		if client.gameroom == c.gameroom {
// 			msg, err := outgoingEvent.Marshal()
// 			if err != nil {
// 				return err
// 			}
// 			client.send <- msg
// 		}
// 	}
// 	return nil
// }

func AttackEventHandler(e AttackEvent, c *Client) {
	c.hub.mu.Lock()
	defer c.hub.mu.Unlock()
	for _, v := range c.hub.rooms {
		fmt.Printf("Room: '%s' ", v.id)
	}
	fmt.Println()
	if room, ok := c.hub.rooms[e.RoomID]; ok {
		r, err := room.Attack(e.AttackerIndex, e.AttackPosition)
		if err != nil {
			return
		}

		if r {
		}

		for i, c := range room.clients {
			outgoingEvent := &Event{Type: "game_update"}
			payload := &GameStartPayload{
				RoomID:      room.id,
				Index:       i,
				Players:     [2]PlayerInfo{},
				Message:     "Hi Hi Captain",
				CurrentTurn: room.CurrentTurn,
				GameOver:    room.GameOver,
				Mode:        room.Mode,
				Winner:      room.CurrentTurn,
			}

			// Send plain board for player
			player1 := room.Players[i]
			payload.Players[i] = PlayerInfo{Board: player1.Board.PlainBoard(), Fleet: player1.GetFleetInfo()}
			// Send masked board for opponent
			player2 := room.Players[1-i]
			payload.Players[1-i] = PlayerInfo{Board: player2.Board.MaskBoard(), Fleet: player2.GetFleetInfo()}

			msg, err := func(payload interface{}) ([]byte, error) {
				fmt.Printf("%+v - \n", payload)
				p, err := json.Marshal(payload)
				if err != nil {
					return nil, fmt.Errorf("An error occured while Marshalling outgoingEvent payload")
				}
				outgoingEvent.Payload = p
				p, err = json.Marshal(outgoingEvent)
				if err != nil {
					return nil, fmt.Errorf("An error occured while Marshalling outgoingEvent")
				}
				return p, nil
			}(payload)

			if err != nil {
				return
			}

			c.send <- msg
		}
	} else {
		fmt.Println("Room not found")
	}
}

func FindGameEventHandler(c *Client) {
	c.hub.mu.Lock()
	defer c.hub.mu.Unlock()
	for _, room := range c.hub.rooms {
		if room == nil || len(room.clients) == 2 {
			continue
		}
		room.clients = append(room.clients, c)
		room.BattleshipGame = *g.NewBattleshipGame(10)

		for i := range room.clients {
			player := room.BattleshipGame.Players[i]
			player.GenerateAndPlaceShips()
			log.Println()
			log.Printf("%+v", player.Ships)
			log.Println()
		}

		for i, c := range room.clients {
			outgoingEvent := &Event{Type: "game_start"}
			payload := &GameStartPayload{
				RoomID:      room.id,
				Index:       i,
				Players:     [2]PlayerInfo{},
				Message:     "Hi Captain, you've been given orders to eliminate the enemy, good luck sailor",
				CurrentTurn: room.CurrentTurn,
				GameOver:    room.GameOver,
				Mode:        room.Mode,
				Winner:      1 - room.CurrentTurn,
			}

			// Send plain board for player
			player1 := room.Players[i]
			payload.Players[i] = PlayerInfo{Board: player1.Board.PlainBoard(), Fleet: player1.GetFleetInfo()}
			// Send masked board for opponent
			player2 := room.Players[1-i]
			payload.Players[1-i] = PlayerInfo{Board: player2.Board.MaskBoard(), Fleet: player2.GetFleetInfo()}

			log.Println()
			log.Printf("After setting payload: %+v", payload.Players)
			log.Println()

			msg, err := func(event string, payload interface{}) ([]byte, error) {
				fmt.Printf("%+v - \n", payload)
				p, err := json.Marshal(payload)
				if err != nil {
					return nil, fmt.Errorf("An error occured while Marshalling outgoingEvent payload")
				}
				outgoingEvent.Payload = p
				p, err = json.Marshal(outgoingEvent)
				if err != nil {
					return nil, fmt.Errorf("An error occured while Marshalling outgoingEvent")
				}
				return p, nil
			}("game_start", payload)

			// msg, err = json.Marshal(outgoing)
			if err != nil {
				return
			}

			c.send <- msg
		}

		fmt.Printf("New user joined, number of clients = %d, number of rooms = %d\n", len(c.hub.clients), len(c.hub.rooms))
		fmt.Printf("room %v has %d clients", room.id, len(room.clients))
		for _, client := range room.clients {
			fmt.Printf(" -- client %s", client.userData["name"])
		}
		fmt.Println()
		return
	}

	newRoomID := fmt.Sprintf("room-%d", c.hub.lastRoomID)
	newRoom := &GameRoom{
		id:      newRoomID,
		clients: make([]*Client, 0),
	}
	newRoom.clients = append(newRoom.clients, c)
	c.hub.rooms[newRoomID] = newRoom
	c.hub.lastRoomID++

	newRoom.Broadcast([]byte(fmt.Sprintf("%s joined room %s, game room is ready\n", c.userData["name"], newRoom.id)))
	fmt.Printf("New user joined, number of clients = %d, number of rooms = %d\n", len(c.hub.clients), len(c.hub.rooms))
	fmt.Printf("room %v has %d clients", newRoom.id, len(newRoom.clients))
	for _, client := range newRoom.clients {
		fmt.Printf(" -- client %s", client.userData["name"])
	}
	fmt.Println()
}
