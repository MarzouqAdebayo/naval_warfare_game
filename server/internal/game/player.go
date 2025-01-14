package game

import (
	"errors"
	"math/rand/v2"
)

type PlayerGameStruct struct {
	Index          int        `json:"index"`
	Board          *Board     `json:"board"`
	Ships          []Ship     `json:"ships"`
	Shotsfired     []Position `json:"shotsFired"`
	RemainingShips int        `json:"remainingShips"`
}

var ErrCannotPlaceShip = errors.New("cannot place ship in this position")
var ErrShipCollision = errors.New("there is a collision in this position")

func NewPlayer(index int, boardSize int) *PlayerGameStruct {
	initialShips := InitializeShips()
	return &PlayerGameStruct{
		Index:          index,
		Board:          GenerateEmptyBoard(boardSize),
		Ships:          initialShips,
		RemainingShips: len(initialShips),
	}
}

func (p *PlayerGameStruct) GenerateAndPlaceShips() {
	// FIXME Due to the size of the ship, board size cannot
	// BUG be less than 5, or else and infinite loop occurs
	if p.Board.Size < 5 {
		return
	}
	ships := InitializeShips()
	axes := [2]Axis{X, Y}

	// Using brute force method lol
	// FIXME Use goroutines to hasten placement
	// FIXME Don't forget to use mutex
	for _, ship := range ships {
		for {
			randomStartPos := Position{X: rand.IntN(p.Board.Size), Y: rand.IntN(p.Board.Size)}
			randomAxis := axes[rand.IntN(2)]

			err := p.PlaceShip(&ship, randomStartPos, randomAxis)
			if err == nil {
				break
			}
		}
	}
}

func (p *PlayerGameStruct) PlaceShip(ship *Ship, startPos Position, axis Axis) error {
	positions, err := p.getShipCoordinates(ship.Size, startPos, axis)
	if err != nil {
		return err
	}

	ship.Positions = positions
	p.Ships = append(p.Ships, *ship)

	for _, pos := range positions {
		p.Board.Squares[pos.X][pos.Y].State = HasShip
	}
	return nil
}

func (p *PlayerGameStruct) getShipCoordinates(shipSize int, startPos Position, axis Axis) ([]Position, error) {
	pos := calculateShipPositions(shipSize, startPos, axis)

	for _, pos := range pos {
		if pos.X < 0 || pos.X >= p.Board.Size || pos.Y < 0 || pos.Y >= p.Board.Size {
			return nil, ErrCannotPlaceShip
		}
		if p.Board.Squares[pos.X][pos.Y].State != Empty {
			return nil, ErrShipCollision
		}
	}
	return pos, nil
}

func calculateShipPositions(shipSize int, startPos Position, axis Axis) []Position {
	positions := make([]Position, shipSize)

	for i := range shipSize {
		if axis == X {
			positions[i] = Position{X: startPos.X, Y: startPos.Y + i}
		} else {
			positions[i] = Position{X: startPos.X + i, Y: startPos.Y}
		}
	}

	return positions
}
