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
	Sunk    CellState = "Sunk"
)

type Squares [][]Cell

type Board struct {
	Size    int
	Squares Squares
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

// MaskBoard returns a copy of the board where ship positions are hidden
func (b *Board) MaskBoard() [][]CellState {
	copy := make([][]CellState, len(b.Squares))
	for i := range b.Squares {
		copy[i] = make([]CellState, len(b.Squares[i]))
		for j := range b.Squares[i] {
			if b.Squares[i][j].State == HasShip {
				copy[i][j] = Empty
			} else {
				copy[i][j] = b.Squares[i][j].State
			}
		}
	}
	return copy
}

// PlainBoard returns an empty board
func (b *Board) PlainBoard() [][]CellState {
	copy := make([][]CellState, len(b.Squares))
	for i := range copy {
		copy[i] = make([]CellState, len(b.Squares[i]))
		for j := range copy[i] {
			copy[i][j] = b.Squares[i][j].State
		}
	}
	return copy
}
