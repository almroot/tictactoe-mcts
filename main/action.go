package main

type Action struct {
	Player Player
	X      int
	Y      int
}

func NewAction(player Player, x, y int) *Action {
	return &Action{
		Player: player,
		X:      x,
		Y:      y,
	}
}
