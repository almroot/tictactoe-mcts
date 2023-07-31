package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"tictactoe-mcts/history"
	"time"

	"git.sr.ht/~bonbon/gmcts"
)

type Actor struct {
	Player   Player
	TakeTurn func(game gmcts.Game) (gmcts.Action, error)
}

func NewActorAI(history history.History, player Player, ephemeral func(message string), concurrency int, timeout time.Duration) *Actor {
	return &Actor{
		Player: player,
		TakeTurn: func(game gmcts.Game) (gmcts.Action, error) {
			history.Ephemeral(player.String(), "thinking...", ephemeral)
			var wait sync.WaitGroup
			mcts := gmcts.NewMCTS(game)
			start := time.Now().UnixMilli()
			ctr := 0
			wait.Add(concurrency)
			for i := 0; i < concurrency; i++ {
				go func() {
					tree := mcts.SpawnTree()
					tree.Search(timeout)
					ctr += tree.Nodes()
					mcts.AddTree(tree)
					wait.Done()
				}()
			}
			wait.Wait()
			result := mcts.BestAction()
			history.Report(
				player.String(),
				fmt.Sprintf(
					"%dx%d (explored %v nodes in %vms)",
					result.(*Action).X+1,
					result.(*Action).Y+1,
					ctr,
					time.Now().UnixMilli()-start,
				),
			)
			return result, nil
		},
	}
}

func NewActorHuman(history history.History, player Player, ephemeral func(message string)) *Actor {
	return &Actor{
		Player: player,
		TakeTurn: func(game gmcts.Game) (gmcts.Action, error) {
			var bestAction gmcts.Action
			var legalActions = game.GetActions()
			var reader = bufio.NewReader(os.Stdin)
			history.Ephemeral(player.String(), "", ephemeral)
			var line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)
			history.Report(player.String(), line)
			if len(line) != 3 {
				const msg = "invalid command"
				return nil, errors.New(msg)
			}
			horizontal, vertical, _ := strings.Cut(line, "x")
			y, _ := strconv.Atoi(horizontal)
			x, _ := strconv.Atoi(vertical)
			x--
			y--
			for _, a := range legalActions {
				if a.(*Action).X == x && a.(*Action).Y == y && a.(*Action).Player == Player(game.Player()) {
					bestAction = NewAction(Player(game.Player()), x, y)
				}
			}
			if bestAction != nil {
				return bestAction, nil
			}
			return nil, errors.New("invalid move")
		},
	}
}
