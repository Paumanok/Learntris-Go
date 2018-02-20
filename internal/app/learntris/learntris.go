package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"learntris/internal/pkg/matrix"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	board := matrix.Create_mat(22,10)
	var active (*[][]rune) = &Empty_tetramino
	score, cleared_lines := 0, 0 //init values
	

	for scanner.Scan() {
		
		text := scanner.Text()
		//fmt.Println("text: " , text[0], " len: ", len(text))
		if len(text) > 0 {
			char := text[0]
			switch {
			case char == 'q':
				//quit
				os.Exit(3)
			case char == 'p':
				//print the state of the matrix
				matrix.Print_mat(board)
			
			case char == 'g':
				//set board to given input
				given(&board, scanner)

			case char == 'c':
				//clear matrix
				clear(&board)

			case char == 's':
				//run single step
				single_step(&board, &cleared_lines, &score)

			case char == 'I':
				active = &I_tetramino

			case char == 'O':
				active = &O_tetramino

			case char == 't':
				matrix.Print_mat(*active)

			case char == '?':
				//query 
				switch {
				//score
				case text[1] == 's':
					 fmt.Println(score)
				//cleared lines
				case text[1] == 'n':
					fmt.Println(cleared_lines)
				}


			default:
				//still quit
				fmt.Println("case not hit")
				os.Exit(3)
			}
		}
	}
}

func given(board *([][]rune), scanner *bufio.Scanner) {
	//expecting 22 lines of 10 items separated by spaces
	i := 0 //my iterator for num lines(rows)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) >= 10{
			//fmt.Println(text)
			//need to split and throw away spaces
			split_line := strings.Split(text, " ")
			for j, j_str := range(split_line){
				(*board)[i][j] = rune(j_str[0])
			}
			i++

		}

		if i == 22 {break} //stop at 22 lines
	}
	if scanner.Err() != nil{
		fmt.Println(scanner.Err())
	}

}

func single_step(board (*[][]rune), n *int, s *int){
	//iterate over board
	//find complete line
	//if complete
	//	clear
	//  update score

	for i := range(*board){
		count := 0
		for j := range((*board)[i]){
			if (*board)[i][j] != '.'{count++}
			//fmt.Println(count)
			//fmt.Println("row" )
			//fmt.Println(i)
			if count == 10 {
				clear_row(board, i) //clear the cool row
				*n += 1
				*s += 100

			}
		}
	}
}

func clear_row(board *([][]rune), row int){
	for i := range((*board)[row]){
		(*board)[row][i] = '.'
	}
}

func clear(board *([][]rune)) {
	for i := range *board {
		for j := range (*board)[i] { 
			(*board)[i][j] = '.'
		}

	}
}
