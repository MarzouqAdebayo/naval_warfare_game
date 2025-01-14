package ws

import (
	g "server/internal/game"
)

// Message represents the structure of WebSocket messages
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// GameRoom represents the BattleshipGame with extra info
type GameRoom struct {
	id      string
	clients []*Client
	g.BattleshipGame
}

func (g *GameRoom) newRoom() {}
