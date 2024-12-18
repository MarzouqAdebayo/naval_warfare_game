package game

type Cell struct {
	Position Position
	State    CellState
}

type CellState string

const (
	Empty   CellState = "Empty"
	HasShip CellState = "Ship"
	Hit     CellState = "Hit"
	Miss    CellState = "Miss"
)

type Board struct {
	Size  int
	Board [][]Cell
}

func GenerateEmptyBoard(size int) *Board {
	board := &Board{
		Size:  size,
		Board: make([][]Cell, size),
	}

	for i := range size {
		board.Board[i] = make([]Cell, size)
		for j := range size {
			board.Board[i][j] = Cell{Position: Position{X: i, Y: j}, State: Empty}
		}
	}

	return board
}
