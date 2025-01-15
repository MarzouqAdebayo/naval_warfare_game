package game

import (
	"reflect"
	"testing"
)

func TestIntializeShips(t *testing.T) {
	t.Run("should intialize ships with correct values", func(t *testing.T) {
		want := []Ship{
			{Type: Carrier, Size: 5, Hits: 0, Sunk: false},
			{Type: Battleship, Size: 4, Hits: 0, Sunk: false},
			{Type: Destroyer, Size: 3, Hits: 0, Sunk: false},
			{Type: Submarine, Size: 3, Hits: 0, Sunk: false},
			{Type: Cruiser, Size: 2, Hits: 0, Sunk: false},
		}

		got := InitializeShips()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %+v got %+v", want, got)
		}
	})
}
