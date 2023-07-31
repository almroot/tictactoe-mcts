package tictactoe

type Player int

const (
	PlayerX Player = +1
	PlayerO Player = -1
)

func (p Player) String() string {
	switch p {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	default:
		return ""
	}
}
