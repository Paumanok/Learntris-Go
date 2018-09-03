//file defining board struct and init 
package main

import (
	"learntris/internal/pkg/matrix"
)
type Board struct {
	board [][]rune
	active [][]rune
	score int
	cleared_lines int
	active_row int
	active_col int
}

func init_board() *Board {
	return &Board{
		board : matrix.Create_mat(22,10),
		active : Empty_tetramino,
		score : 0,
		cleared_lines : 0}
	}


