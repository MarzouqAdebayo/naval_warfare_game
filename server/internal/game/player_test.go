package game

import (
	"testing"
)

func TestPlaceShips(t *testing.T) {
	t.Run("should set ship on board horizontally", func(t *testing.T) {
		boardSize := 4
		shipSize := 4
		player := &Player{}
		player.Board = GenerateEmptyBoard(boardSize)
		ship := &Ship{Type: Destroyer, Size: shipSize}
		err := player.PlaceShips(ship, Coordinates{X: 0, Y: 0}, X)
		if err != nil {
			t.Errorf("%v", err)
		} else {
			for i := range shipSize {
				if player.Board.Board[0][i].State != HasShip {
					t.Errorf("Cell %d %d should contain Ship", 0, i)
				}
			}
		}
	})

	t.Run("should set ship on board vertically", func(t *testing.T) {
		boardSize := 4
		shipSize := 4
		player := &Player{}
		player.Board = GenerateEmptyBoard(boardSize)
		ship := &Ship{Type: Destroyer, Size: shipSize}
		err := player.PlaceShips(ship, Coordinates{X: 0, Y: 3}, Y)
		if err != nil {
			t.Errorf("%v", err)
		} else {
			for i := range shipSize {
				if player.Board.Board[i][3].State != HasShip {
					t.Errorf("Cell %d %d should contain Ship", i, 3)
				}
			}
		}
	})

	t.Run("should return error if ship cannot be placed on board", func(t *testing.T) {
		boardSize := 5
		shipSize := 5
		player := &Player{}
		player.Board = GenerateEmptyBoard(boardSize)
		ship := &Ship{Type: Destroyer, Size: shipSize}
		err := player.PlaceShips(ship, Coordinates{X: 3, Y: 0}, Y)

		assertError(t, err, ErrCannotPlaceShip)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
