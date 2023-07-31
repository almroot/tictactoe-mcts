package tictactoe

import "errors"

type Board struct {
	rows [][]int
}

func NewBoard() *Board {
	var board = &Board{}
	for i := 0; i < 3; i++ {
		board.rows = append(board.rows, []int{0, 0, 0})
	}
	return board
}

func (b *Board) Clone() *Board {
	var out = &Board{}
	for _, row := range b.rows {
		var clonedRow []int
		for _, cell := range row {
			clonedRow = append(clonedRow, cell)
		}
		out.rows = append(out.rows, clonedRow)
	}
	return out
}

func (b *Board) Get() [][]int {
	return b.Clone().rows
}

func (b *Board) Set(player Player, x, y int) error {
	const msg = "out of bounds"
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return errors.New(msg)
	}
	b.rows[x][y] = int(player)
	return nil
}

func (b *Board) Winner() (Player, bool) {
	for x := 0; x < 3; x++ {
		if b.rows[x][0] != 0 && b.rows[x][0] == b.rows[x][1] && b.rows[x][1] == b.rows[x][2] {
			return Player(b.rows[x][0]), true
		}
	}
	for y := 0; y < 3; y++ {
		if b.rows[0][y] != 0 && b.rows[0][y] == b.rows[1][y] && b.rows[1][y] == b.rows[2][y] {
			return Player(b.rows[0][y]), true
		}
	}
	if b.rows[1][1] != 0 {
		if b.rows[1][1] == b.rows[2][0] && b.rows[2][0] == b.rows[0][2] {
			return Player(b.rows[1][1]), true
		} else if b.rows[1][1] == b.rows[0][0] && b.rows[0][0] == b.rows[2][2] {
			return Player(b.rows[1][1]), true
		}
	}
	return 0, false
}
