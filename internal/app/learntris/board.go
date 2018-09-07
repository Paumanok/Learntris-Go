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


func Place_active(board(*Board)){
	start_col := board.active.spawn_col
	board.active.start_x = start_col
	board.active.start_y = 0
	board.active.end_x = start_col + len(board.active.tet[0])-1
	board.active.end_y = len(board.active.tet)-1
}


//shift active tetramino on board
//needs bounds checking
func shift_active_hor(shift_by(int), board(*Board)){
	new_start_x := board.active.start_x + shift_by
	new_end_x := board.active.end_x + shift_by
	//bounds checking:
	//check if start or end is outside size of board
	//if so, check if this only consists of '.'
	//if not, carry on, if so, dont move
	if new_start_x < 0 || new_end_x > (len(board.board[0])-1) {
		for row := range board.active.tet {
			if board.active.tet[row][Neg_offset(board.active.start_x)] != '.'  {
				//fail state
				return
			}else if board.active.tet[row][Pos_offset(board)] != '.' {
				//fmt.Println("fail state")
				return
			}
		}
	}

	board.active.start_x = new_start_x
	board.active.end_x = new_end_x
}

func shift_active_vert(shift_by(int), board(*Board)){
	new_start_y := board.active.start_y + shift_by
	new_end_y := board.active.end_y + shift_by
	
	if 0 {
	//more checking needed here
	}
	board.active.start_y = new_start_y
	board.active.end_y = new_end_y
}

func Print_active_board(board(Board)){
	a_i := 0
	a_j := Neg_offset(board.active.start_x)
	for i := range board.board { //for row
		line := ""
		for j := range board.board[i] { //for col
			
			if board.active.start_x <= j && j <= board.active.end_x && board.active.start_y <= i && i <= board.active.end_y {
				a_char := board.active.tet[a_i][a_j]
				
				if a_char != '.' { a_char = a_char - 0x20 }
				line +=string(a_char) + " "
				a_j++
			}else{

				line +=string( board.board[i][j]) + " " 
			}
			
		}
		a_j = Neg_offset(board.active.start_x)
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

//because golang aparently doesn't have an absolute value
//function for non float64 types?
func Neg_offset(val  int) int{
	if val < 0 {
		return ((-val)-1)
	}
	return 0
}

func Pos_offset(board (*Board)) int{
	if board.active.end_x >= 9{
		//fmt.Println(val - (val - 9))
		return (len(board.active.tet[0])-1) - (board.active.end_x - 9)
	}
	return len(board.active.tet[0])-1
}
