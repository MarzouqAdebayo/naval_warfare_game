package game

type ShipType string

const (
	Carrier    ShipType = "Carrier"
	Battleship ShipType = "Battleship"
	Cruiser    ShipType = "Cruiser"
	Submarine  ShipType = "Submarine"
	Destroyer  ShipType = "Destroyer"
)

type Position struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Axis string

const (
	X Axis = "Horizontal"
	Y Axis = "Vertical"
)

type Ship struct {
	Type      ShipType
	Size      int
	Positions []Position
	Hits      uint8
	Sunk      bool
}

// Initialize ships for a new game
func InitializeShips() []Ship {
	return []Ship{
		{Type: Carrier, Size: 5, Hits: 0, Sunk: false},
		{Type: Battleship, Size: 4, Hits: 0, Sunk: false},
		{Type: Cruiser, Size: 3, Hits: 0, Sunk: false},
		{Type: Submarine, Size: 3, Hits: 0, Sunk: false},
		{Type: Destroyer, Size: 2, Hits: 0, Sunk: false},
	}
}
