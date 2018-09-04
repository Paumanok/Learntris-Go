package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//board := matrix.Create_mat(22,10)
	//score, cleared_lines := 0, 0 //init values
	
	var board Board = *init_board()


	//var active ([][]rune) = Empty_tetramino

	for scanner.Scan() {
		
		text := scanner.Text()
		skipNext := false
		//split_commands := strings.Split(text, " ")
		commands := strings.Replace(text, " ", "", -1)
		//fmt.Println("text: " , text[0], " len: ", len(text))
		//if len(text) > 0 {
		for i, char := range(commands){
			switch {
			case skipNext == true:
				skipNext = false
			case char == 'q':
				//quit
				os.Exit(3)
			case char == 'p':
				//print the state of the matrix
				Print_board(board.board)
			
			case char == 'P':
				//print state of the board with active placed
				Place_active(&board)
				Print_active_board(board)
			case char == 'g':
				//set board to given input
				given(&board.board, scanner)

			case char == 'c':
				//clear matrix
				clear(&board.board)

			case char == 's':
				//run single step
				single_step(&board.board, &board.cleared_lines, &board.score)
			
			//tetramino assignments
			case char == 'I':
				board.active = I_tet

			case char == 'O':
				board.active = O_tet

			case char == 'Z':
				board.active = Z_tet

			case char == 'S':
				board.active = S_tet

			case char == 'J':
				board.active = J_tet
			
			case char == 'L':
				board.active = L_tet			
			case char == 'T':
				board.active = T_tet
			
			case char == '[':
				board.active = Testramino	
			
			case char == 't':
				//set active tetramino
				Print_board(board.active.tet)
			
			case char == ')':
				rotate_active(&board.active.tet)	

			case char == ';':
				fmt.Println("")
			case char == '?':
				//query 
				skipNext = true
				next_char := commands[i+1]
				switch {
				//score
				case next_char == 's':
					 fmt.Println(board.score)
				//cleared lines
				case next_char == 'n':
					fmt.Println(board.cleared_lines)
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

//rotate active tetramino clockwise
func rotate_active(active *([][]rune)){
	temp_tet := *active
	N := len(temp_tet)
	//matrix is y,x [rows][col]
	for x := 0; x < N/2; x+=1 {
		for y:=x; y < N-x-1; y+=1 {
			x_off := N-1-x
			y_off := N-1-y

			temp := temp_tet[x][y]
			temp_tet[x][y] = temp_tet[y_off][x] //bottom left
			temp_tet[y_off][x] = temp_tet[x_off][y_off]
			temp_tet[x_off][y_off] = temp_tet[y][x_off]
			temp_tet[y][x_off] = temp
			}	
		}
	*active = temp_tet
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
