package game

import (
	"testing"
)

func TestGenerateEmptyBoard(t *testing.T) {
	t.Run("should intialize board with default parameters", func(t *testing.T) {
		boardSize := 5
		board_test := GenerateEmptyBoard(boardSize)

		if board_test.Size != boardSize {
			t.Errorf("board size is expected to be %v got %v", boardSize, board_test.Size)
		}
		if len(board_test.Squares) != boardSize {
			t.Errorf("board length is expected to be %v got %v", boardSize, board_test.Size)
		}
		for idxi, i := range board_test.Squares {
			for idxj, j := range i {
				if j.State != Empty {
					t.Errorf("cell (%v, %v) is expected to be %s", idxi, idxj, Empty)
				}
				if j.Position.X != idxi || j.Position.Y != idxj {
					t.Errorf("Cell Position: want (%v, %v), got (%v, %v)", idxi, idxj, j.Position.X, j.Position.Y)
				}
			}
		}
	})
}
