package game

import (
	"log"
	"testing"
)

var boardSize int
var shipSize int
var player1 *PlayerGameStruct
var player1ship *Ship
var player2 *PlayerGameStruct
var player2ship *Ship
var game *BattleshipGame

// func TestMain(m *testing.M) {
// 	setup()
// 	m.Run()
// 	teardown()
// }

func setup() func() {
	boardSize = 5
	shipSize = 5

	player1 = &PlayerGameStruct{Index: 0, RemainingShips: 1}
	player1.Board = GenerateEmptyBoard(boardSize)
	player1ship = &Ship{Type: Destroyer, Size: shipSize}
	player1.PlaceShip(player1ship, Position{X: 0, Y: 0}, X)

	player2 = &PlayerGameStruct{Index: 1, RemainingShips: 1}
	player2.Board = GenerateEmptyBoard(boardSize)
	player2ship = &Ship{Type: Destroyer, Size: shipSize}
	player2.PlaceShip(player2ship, Position{X: 0, Y: 0}, Y)

	game = &BattleshipGame{
		Players: [2]*PlayerGameStruct{player1, player2},
		Mode:    SingleFire,
	}
	return func() {
	}
}

func teardown() {
	log.Println("\n-----Teardown complete-----")
}

func TestAttack(t *testing.T) {
	t.Run("should attack other player's ship", func(t *testing.T) {
		defer setup()()
		target := Position{X: 0, Y: 0}
		_, err := game.Attack(1, target)
		assertError(t, err, ErrNotPlayerTurn)
		_, err = game.Attack(0, target)
		if err != nil {
			t.Errorf("got error %s", err.Error())
		}
		if player2.Ships[0].Hits != 1 {
			t.Errorf("ships hits should be 1")
		}
		if player2.Board.Squares[target.X][target.Y].State != Hit {
			t.Errorf("state at %d %d is supposed to be HIT", target.X, target.Y)
		}
		if player2ship.Positions[0] != target {
			t.Errorf("ship position %d %d, target Position %d %d", player2ship.Positions[0].X, player2ship.Positions[0].Y, target.X, target.Y)
		}
	})

	t.Run("should run in correct game mode", func(t *testing.T) {
		defer setup()()
		target := Position{X: 0, Y: 0}
		_, err := game.Attack(0, target)

		if err != nil {
			t.Error(err.Error())
		}

		if game.CurrentTurn != 1 {
			t.Errorf("should be player 1's turn in single fire mode")
		}

		defer setup()()
		game = &BattleshipGame{
			Players:     [2]*PlayerGameStruct{player1, player2},
			Mode:        ContinousFire,
			CurrentTurn: 1,
		}

		target = Position{X: 2, Y: 0}
		_, err = game.Attack(1, target)

		if game.CurrentTurn != 0 {
			t.Errorf("should be player 0's turn in continous fire mode")
		}

		for i := range 4 {
			target := Position{X: i, Y: 0}
			_, err := game.Attack(0, target)

			if err != nil {
				t.Error(err.Error())
			}
		}

		if game.CurrentTurn != 0 {
			t.Errorf("should be player 0's turn in continous fire mode")
		}
	})

	t.Run("should set correct winner after game is over", func(t *testing.T) {
		defer setup()()
		i := 0
		j := 0
		for {
			if game.CurrentTurn == 0 {
				game.Attack(game.CurrentTurn, Position{X: i, Y: 0})
				i++
			} else {
				game.Attack(game.CurrentTurn, Position{X: 0, Y: j})
				j++
			}
			if i == 5 || j == 5 {
				break
			}
		}

		for i := range boardSize {
			for j := range 1 {
				if player2.Board.Squares[i][j].State != Sunk {
					t.Errorf("Ship should be sunk but its is not")
				}
			}
		}

		if !game.GameOver {
			t.Errorf("game should be over")
		}

		if game.Winner != player1 {
			t.Errorf("game winner is %p %d should be %p %d", game.Winner, game.Winner.Index, player1, player1.Index)
		}

		result, err := game.Attack(game.CurrentTurn, Position{X: 0, Y: j})

		if result {
			t.Errorf("expected %v, got %v", false, result)
		}

		assertError(t, err, ErrGameOver)
	})

	t.Run("should set correct state for Miss", func(t *testing.T) {
		defer setup()()
		target := Position{X: 2, Y: 2}
		result, err := game.Attack(0, target)

		if err != nil {
			t.Error(err.Error())
		}

		if result {
			t.Errorf("expected %v, got %v", false, result)
		}

		player2TargetState := player2.Board.Squares[target.X][target.Y].State
		if player2TargetState != Miss {
			t.Errorf("expected %v, got %v", Miss, player2TargetState)
		}
	})
}
