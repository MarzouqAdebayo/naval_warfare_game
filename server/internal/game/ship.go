package game

import (
	"strings"
)

type ShipType string

const (
	Carrier    ShipType = "Carrier"
	Battleship ShipType = "Battleship"
	Submarine  ShipType = "Submarine"
	Destroyer  ShipType = "Destroyer"
	Cruiser    ShipType = "Cruiser"
)

type Position struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Axis string

const (
	X Axis = "X"
	Y Axis = "Y"
)

type Ship struct {
	Type      ShipType
	Size      int
	Axis      Axis
	Positions []Position
	Hits      uint8
	Sunk      bool
}

type ShipInfoStruct struct {
	Type   string `json:"type"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Axis   Axis   `json:"axis"`
	Length int    `json:"length"`
	Sunk   bool   `json:"sunk"`
}

// Initialize ships for a new game
func InitializeShips() []Ship {
	return []Ship{
		{Type: Carrier, Size: 5, Hits: 0, Sunk: false},
		{Type: Battleship, Size: 4, Hits: 0, Sunk: false},
		{Type: Destroyer, Size: 3, Hits: 0, Sunk: false},
		{Type: Submarine, Size: 3, Hits: 0, Sunk: false},
		{Type: Cruiser, Size: 2, Hits: 0, Sunk: false},
	}
}

// For getting information about a players ship
// for client side rendering
func (ship *Ship) GetShipInfo() ShipInfoStruct {
	return ShipInfoStruct{
		Type:   strings.ToLower(string(ship.Type)),
		X:      ship.Positions[0].X,
		Y:      ship.Positions[0].Y,
		Axis:   ship.Axis,
		Sunk:   ship.Sunk,
		Length: len(ship.Positions),
	}
}
