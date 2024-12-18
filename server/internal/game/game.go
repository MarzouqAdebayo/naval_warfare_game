package game

type Game struct {
	Players     [2]*Player
	CurrentTurn int
	GameOver    bool
	Winner      *Player
}

func NewGame(boardSize int) *Game {
	game := &Game{
		Players:     [2]*Player{},
		CurrentTurn: 0,
		GameOver:    false,
	}

	for i := range 2 {
		initialShips := InitializeShips()
		game.Players[i] = &Player{
			Board:          GenerateEmptyBoard(boardSize),
			Ships:          initialShips,
			RemainingShips: len(initialShips),
		}
	}

	return game
}
