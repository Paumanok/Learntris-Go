package matrix

import (
	"fmt"
)

func Create_mat(row(int), col(int))([][]rune) {

	board := make([][]rune, row)
	for i := 0; i < row; i++ {
		board[i] = make([]rune, col)
		for j := 0; j < col; j++ {
			board[i][j] = '.'
		}
	}
	return board
}


func Print_mat(board([][]rune)) {
	for i := range board {
		line := ""
		for j := range board[i] {
			line +=string( board[i][j]) + " " 
		
		}
		fmt.Println(line)
		
	}
				
}	

func main() {
	board := Create_mat(22, 10)
	Print_mat(board)
}
