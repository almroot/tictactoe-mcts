package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/inancgumus/screen"

	"git.sr.ht/~bonbon/gmcts"
)

var messages []string

const (
	concurrency = 4
	timeout     = time.Second
)

type playMove func(game gmcts.Game) (gmcts.Action, error)

var playerImplementations = []playMove{
	playMoveHuman,
	playMoveAI,
}

func main() {

	// Initialize required variables
	var playerIdx = 0
	var action gmcts.Action
	var err error

	// Prepare the board
	rand.Seed(time.Now().UnixMilli())
	var game gmcts.Game = NewGame(PlayerX)
	renderBoard(game.(*Game))

	// Take turns playing...
	for !game.IsTerminal() {
		for {
			action, err = playerImplementations[playerIdx](game)
			if err != nil {
				messages = append(messages, err.Error())
				renderBoard(game.(*Game))
				continue
			}
			break
		}
		playerIdx++
		playerIdx = playerIdx % len(playerImplementations)
		if game, err = game.ApplyAction(action); err != nil {
			panic(err)
		}
		renderBoard(game.(*Game))
	}

	// Show the winner...
	if len(game.Winners()) == 0 {
		messages = append(messages, fmt.Sprintf("[%s] <SYSTEM> $ game ended in draw", time.Now().Format(time.DateTime)))
	} else {
		messages = append(messages, fmt.Sprintf("[%s] <SYSTEM> $ the winner is %s!", time.Now().Format(time.DateTime), Player(game.Winners()[0])))
	}
	renderBoard(game.(*Game))
}

func applyMove(game gmcts.Game, move *Action) gmcts.Game {
	if nextState, err := game.ApplyAction(move); err != nil {
		panic(err)
	} else {
		return nextState
	}
}

func playMoveHuman(game gmcts.Game) (gmcts.Action, error) {
	var bestAction gmcts.Action
	var legalActions = game.GetActions()
	var reader = bufio.NewReader(os.Stdin)
	var prefix = fmt.Sprintf("[%s] <HUMAN>  $ ", time.Now().Format(time.DateTime))
	fmt.Print(prefix)
	var line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	messages = append(messages, prefix+line)
	if len(line) != 3 {
		return nil, fmt.Errorf("[%s] <SYSTEM> $ invalid command", time.Now().Format(time.DateTime))
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
	return nil, fmt.Errorf("[%s] <SYSTEM> $ invalid move", time.Now().Format(time.DateTime))
}

func playMoveAI(game gmcts.Game) (gmcts.Action, error) {
	fmt.Printf("[%s] <AI>     $ thinking...", time.Now().Format(time.DateTime))
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
	messages = append(messages, fmt.Sprintf("[%s] <AI>     $ explored %v nodes in %vms", time.Now().Format(time.DateTime), ctr, time.Now().UnixMilli()-start))
	return mcts.BestAction(), nil
}

func renderBoard(game *Game) {
	screen.Clear()
	screen.MoveTopLeft()
	var rows = game.board.Get()
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	fmt.Println("=======")
	for _, row := range rows {
		for _, cell := range row {
			fmt.Print("|")
			fmt.Print(Player(cell).String())
		}
		fmt.Println("|")
	}
	fmt.Println("=======")
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
