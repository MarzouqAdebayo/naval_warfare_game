package ws

import (
	"encoding/json"
	"fmt"
	"time"

	g "server/internal/game"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventFindGame        = "find_game"
	EventAttack          = "attack"
	EventBroadcastAttack = "new_attack"
	EventChangeRoom      = "change_room"
)

type ChangeRoomEvent struct {
	Name string `json:"name"`
}

type FindGameEvent struct {
	From string `json:"from"`
}

type AttackEvent struct {
	AttackPosition g.Position
}

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

func (e Event) Marshal() ([]byte, error) {
	d, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal broadcast message: %v", err)
	}
	return d, nil
}

func SendMessageHandler(event Event, h *Hub, c *Client) error {
	var chatevent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventBroadcastAttack
	// Broadcast to all other Clients
	for _, client := range c.hub.clients {
		// Only send to clients inside the same chatroom
		if client.gameroom == c.gameroom {
			msg, err := outgoingEvent.Marshal()
			if err != nil {
				return err
			}
			client.send <- msg
		}
	}
	return nil
}
