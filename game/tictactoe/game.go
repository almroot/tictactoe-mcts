package tictactoe

import (
	"errors"
	"fmt"
	"tictactoe-mcts/game"
	"tictactoe-mcts/history"

	"git.sr.ht/~bonbon/gmcts"
	"github.com/inancgumus/screen"
)

type Game struct {
	history history.History
	board   *Board
	actors  []*Actor
	turn    int
}

func NewGame(history history.History, board *Board, actors ...*Actor) *Game {
	return &Game{
		history: history,
		board:   board,
		actors:  actors,
		turn:    0,
	}
}

func (g *Game) Winner() string {
	var winners = g.Winners()
	if len(winners) > 0 {
		return Player(g.Winners()[0]).String()
	}
	return ""
}

func (g *Game) Progress() (game.Game, error) {
	var action gmcts.Action
	var err error
	for {
		action, err = g.actors[g.turn].TakeTurn(g)
		if err != nil {
			g.history.Error(err)
			g.Render()
			continue
		}
		break
	}
	var stateChange gmcts.Game
	if stateChange, err = g.ApplyAction(action); err != nil {
		return nil, err
	} else if casted, ok := stateChange.(*Game); ok {
		return casted, nil
	} else {
		const msg = "unable to cast"
		return nil, errors.New(msg)
	}
}

func (g *Game) Render() {
	const boardBoundaries = "======="
	const cellDivider = "|"
	const sp = " "
	screen.Clear()
	screen.MoveTopLeft()
	var rows = g.board.Get()
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	fmt.Println(boardBoundaries)
	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cellDivider)
			var symbol = Player(cell).String()
			if len(symbol) == 0 {
				symbol = sp
			}
			fmt.Print(symbol)
		}
		fmt.Println(cellDivider)
	}
	fmt.Println(boardBoundaries)
	for msg := range g.history.List() {
		fmt.Println(msg)
	}
}
