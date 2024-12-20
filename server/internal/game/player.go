package game

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Player struct {
	Board          *Board
	Ships          []Ship
	Shotsfired     []Coordinates
	RemainingShips int
}

var ErrCannotPlaceShip = errors.New("cannot place ship in this position")
var ErrShipCollision = errors.New("there is a collision in this position")

func (p *Player) GenerateAndPlaceShips() {
	// FIXME Due to the size of the ship, board size cannot
	// BUG be less than 5, or else and infinite loop occurs
	if p.Board.Size < 5 {
		return
	}
	ships := InitializeShips()
	axes := [2]Axis{X, Y}

	// Using brute force method lol
	for _, ship := range ships {
		for {
			randomStartPos := Coordinates{X: rand.IntN(p.Board.Size), Y: rand.IntN(p.Board.Size)}
			randomAxis := axes[rand.IntN(2)]

			fmt.Println(randomStartPos, randomAxis)

			err := p.PlaceShip(&ship, randomStartPos, randomAxis)
			if err == nil {
				break
			}
		}
	}
}

func (p *Player) PlaceShip(ship *Ship, startPos Coordinates, axis Axis) error {
	positions, err := p.getShipCoordinates(ship.Size, startPos, axis)
	if err != nil {
		return err
	}

	ship.Positions = positions
	p.Ships = append(p.Ships, *ship)

	for _, pos := range positions {
		p.Board.Board[pos.X][pos.Y].State = HasShip
	}
	return nil
}

func (p *Player) getShipCoordinates(shipSize int, startPos Coordinates, axis Axis) ([]Coordinates, error) {
	pos := calculateShipPositions(shipSize, startPos, axis)

	for _, pos := range pos {
		if pos.X < 0 || pos.X >= p.Board.Size || pos.Y < 0 || pos.Y >= p.Board.Size {
			return nil, ErrCannotPlaceShip

		}
		if p.Board.Board[pos.X][pos.Y].State != Empty {
			return nil, ErrShipCollision
		}
	}
	return pos, nil
}

func calculateShipPositions(shipSize int, startPos Coordinates, axis Axis) []Coordinates {
	positions := make([]Coordinates, shipSize)

	for i := range shipSize {
		if axis == X {
			positions[i] = Coordinates{X: startPos.X, Y: startPos.Y + i}
		} else {
			positions[i] = Coordinates{X: startPos.X + i, Y: startPos.Y}
		}
	}

	return positions
}

func (p *Player) Attack() {}
func (p *Player) Defend() {}
