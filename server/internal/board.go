package internal

type Cell struct {
	x int
	y int
}

type Board struct {
	BoardSize int
	Board     [][]Cell
}

func (b *Board) GenerateEmptyBoard() {
	for i := range b.BoardSize {
		row := []Cell{}
		for j := range b.BoardSize {
			row = append(row, Cell{i, j})
		}
		b.Board = append(b.Board, row)
	}
}
