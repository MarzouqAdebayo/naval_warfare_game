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
	Size    int
	Squares [][]Cell
}

// Generates an empty board
func GenerateEmptyBoard(size int) *Board {
	board := &Board{
		Size:    size,
		Squares: make([][]Cell, size),
	}

	for i := 0; i < size; i++ {
		board.Squares[i] = make([]Cell, size)
		for j := 0; j < size; j++ {
			board.Squares[i][j] = Cell{Position: Position{X: i, Y: j}, State: Empty}
		}
	}

	return board
}
