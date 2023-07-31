package main

import (
	"git.sr.ht/~bonbon/gmcts"
)

type Game struct {
	board      *Board
	playerTurn Player
}

func NewGame(startingPlayer Player) *Game {
	return &Game{
		board:      NewBoard(),
		playerTurn: startingPlayer,
	}
}

func (g *Game) GetActions() []gmcts.Action {
	var output []gmcts.Action
	for x, row := range g.board.Get() {
		for y, cell := range row {
			if cell == 0 {
				output = append(output, NewAction(g.playerTurn, x, y))
			}
		}
	}
	return output
}

func (g *Game) ApplyAction(action gmcts.Action) (gmcts.Game, error) {
	var board = g.board.Clone()
	if err := board.Set(action.(*Action).Player, action.(*Action).X, action.(*Action).Y); err != nil {
		return nil, err
	}
	var newPlayerTurn Player
	if g.playerTurn == PlayerX {
		newPlayerTurn = PlayerO
	} else {
		newPlayerTurn = PlayerX
	}
	return &Game{
		board:      board,
		playerTurn: newPlayerTurn,
	}, nil
}

func (g *Game) Player() gmcts.Player {
	return gmcts.Player(g.playerTurn)
}

func (g *Game) IsTerminal() bool {
	for _, row := range g.board.Get() {
		for _, cell := range row {
			if cell == 0 {
				return len(g.Winners()) > 0
			}
		}
	}
	return true
}

func (g *Game) Winners() []gmcts.Player {
	if winner, ok := g.board.Winner(); ok {
		return []gmcts.Player{gmcts.Player(winner)}
	}
	return nil
}
