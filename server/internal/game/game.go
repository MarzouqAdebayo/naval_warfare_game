package game

import "errors"

type GameMode int

const (
	ContinousFire GameMode = iota
	SingleFire
)

type GameStatus int

const (
	Waiting GameStatus = iota
	Ready
)

var GameModeName = map[GameMode]string{
	ContinousFire: "continous fire",
	SingleFire:    "single fire",
}

func (g GameMode) String() string {
	return GameModeName[g]
}

type BattleshipGame struct {
	Players     [2]*PlayerGameStruct
	CurrentTurn int
	GameOver    bool
	Mode        GameMode
	Winner      *PlayerGameStruct
}

var ErrNotPlayerTurn = errors.New("it is not your turn yet")
var ErrPositionIsAttacked = errors.New("this position has already been attacked")
var ErrGameOver = errors.New("this game is over")

// Starts a new game of Naval Warfare
func NewBattleshipGame(boardSize int) *BattleshipGame {
	game := &BattleshipGame{
		Players:     [2]*PlayerGameStruct{},
		CurrentTurn: 0,
		GameOver:    false,
	}

	for i := range 2 {
		p := NewPlayer(i, boardSize)
		game.Players[i] = p
	}

	return game
}

// Attacks a position on the opponents board
func (game *BattleshipGame) Attack(attackerIndex int, targetCoordinates Position) (bool, error) {
	if game.GameOver {
		return false, ErrGameOver
	}
	if game.CurrentTurn != attackerIndex {
		return false, ErrNotPlayerTurn
	}

	defender := game.Players[1-attackerIndex]

	if containsPosition(defender.Shotsfired, targetCoordinates) {
		return false, ErrPositionIsAttacked
	}

	defender.Shotsfired = append(defender.Shotsfired, targetCoordinates)

	square := &defender.Board.Squares[targetCoordinates.X][targetCoordinates.Y]
	if square.State == HasShip {
		square.State = Hit

		for i := range defender.Ships {
			ship := &defender.Ships[i]
			if isCurrentShipPosition(ship, targetCoordinates) {
				ship.Hits++
				if ship.Hits == uint8(ship.Size) {
					ship.Sunk = true
					for _, pos := range ship.Positions {
						defender.Board.Squares[pos.X][pos.Y].State = Sunk
					}
					defender.RemainingShips--
				}
				break
			}
		}

		if defender.RemainingShips == 0 {
			game.GameOver = true
			game.Winner = game.Players[attackerIndex]
		}

		if game.Mode == SingleFire {
			game.CurrentTurn = 1 - game.CurrentTurn
		}

		return true, nil
	}

	square.State = Miss
	game.CurrentTurn = 1 - game.CurrentTurn
	return false, nil
}

// Checks if a position is a ships position
func isCurrentShipPosition(ship *Ship, targetCoordinates Position) bool {
	for _, shipPosition := range ship.Positions {
		if shipPosition == targetCoordinates {
			return true
		}
	}
	return false
}

// Checks if this position has been attacked
func containsPosition(positions []Position, coordinates Position) bool {
	for _, p := range positions {
		if p == coordinates {
			return true
		}
	}
	return false
}
