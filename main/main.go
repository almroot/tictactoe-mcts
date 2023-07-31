package main

import (
	"fmt"
	"math/rand"
	"os"
	"tictactoe-mcts/game"
	"tictactoe-mcts/game/tictactoe"
	"tictactoe-mcts/history/general"
)

func main() {

	// Parse the CLI options...
	var options = NewOptions()
	if status, terminate := options.Parse(os.Args[1:], os.Stderr); terminate {
		os.Exit(status)
	}

	// Prepare the board
	rand.Seed(options.Seed)
	var history = general.NewHistory()
	var ephemeral = func(message string) {
		_, _ = fmt.Print(message)
	}
	var state game.Game = tictactoe.NewGame(
		history,
		tictactoe.NewBoard(),
		tictactoe.NewActorAI(history, tictactoe.PlayerO, ephemeral, options.Parallelization, options.Timeout),
		tictactoe.NewActorHuman(history, tictactoe.PlayerX, ephemeral),
	)
	state.Render()

	// Take turns playing...
	var err error
	for !state.IsTerminal() {
		for {
			state, err = state.Progress()
			if err == nil {
				state.Render()
				break
			}
			history.Error(err)
		}
	}

	// Show the winner...
	if len(state.Winners()) == 0 {
		history.System("game ended in draw")
	} else {
		history.System(fmt.Sprintf("the winner is player %v!", state.Winner()))
	}
	state.Render()
}
