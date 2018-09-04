//file defining board struct and init 
package main

import (
	"fmt"
)
type Board struct {
	board [][]rune
	active Tetramino
	score int
	cleared_lines int
	active_row int
	active_col int
}

type Active_tet struct{
	active [][]rune
	top_row int
	left_col int
	bot_row int
	right_col int
}

func Place_active(board(*Board)){
	start_col := board.active.spawn_col
	board.active.start_x = start_col
	board.active.start_y = 0
	board.active.end_x = start_col + len(board.active.tet[0])-1
	board.active.end_y = len(board.active.tet)-1
}

func Print_active_board(board(Board)){
	a_i := 0
	a_j := 0
	for i := range board.board { //for row
		line := ""
		for j := range board.board[i] { //for col
			//fmt.Println(j)
			//fmt.Println(board.active.end_x)
			//fmt.Println(board.active.start_x <= j)
			//fmt.Println(board.active.end_x >= j)
			//fmt.Println(board.active.start_y <= i)
			//fmt.Println(board.active.end_y >=i)
			//fmt.Println()
			if board.active.start_x <= j && j <= board.active.end_x && board.active.start_y <= i && i <= board.active.end_y {
				line +=string(board.active.tet[a_i][a_j]) + " "
				a_j++
			}else{

				line +=string( board.board[i][j]) + " " 
			}
			
		}
		a_j = 0
		if board.active.start_y <= i && i <= board.active.end_y { a_i++}
		fmt.Println(line)
	}
}

func init_board() *Board {
	return &Board{
		board : Create_board(22,10),
		active : Empty_tet,
		score : 0,
		cleared_lines : 0}
	}


func Create_board(row(int), col(int))([][]rune) {

	board := make([][]rune, row)
	for i := 0; i < row; i++ {
		board[i] = make([]rune, col)
		for j := 0; j < col; j++ {
			board[i][j] = '.'
		}
	}
	return board
}


func Print_board(board([][]rune)) {
	for i := range board {
		line := ""
		for j := range board[i] {
			line +=string( board[i][j]) + " " 
		
		}
		fmt.Println(line)
	}
}

