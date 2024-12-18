package game

type Player struct {
	Board          *Board
	Ships          []Ship
	Shotsfired     []Position
	RemainingShips int
}

func (p *Player) GenerateShips() {}
func (p *Player) PlaceShips()    {}
func (p *Player) Attack()        {}
func (p *Player) Defend()        {}
