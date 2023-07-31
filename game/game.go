package game

import "git.sr.ht/~bonbon/gmcts"

type Game interface {
	gmcts.Game
	Winner() string
	Render()
	Progress() (Game, error)
}
