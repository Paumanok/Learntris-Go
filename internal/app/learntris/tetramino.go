package main


type Tetramino struct {
	tet [][]rune
	spawn_col int

	//these are set upon activation. Refer to top right point
	start_x int
	start_y int
	end_x int
	end_y int
}
		

//define ur tetraminos 

var Empty_tet Tetramino = Tetramino {
	tet : [][]rune {{'.','.','.','.'},
					{'.','.','.','.'},
					{'.','.','.','.'},
					{'.','.','.','.'}},
	spawn_col : 0}


var I_tet Tetramino = Tetramino{
	tet : [][]rune {{'.','.','.','.'},
					{'c','c','c','c'},
					{'.','.','.','.'},
					{'.','.','.','.'}},
	spawn_col : 2}
							
var O_tet Tetramino = Tetramino{
	tet : [][]rune {{'y','y'},
            		{'y','y'}},
	spawn_col : 4}


var Z_tet Tetramino = Tetramino{
	tet : [][]rune {{'r','r','.'},
		  	  		{'.','r','r'},
           	  		{'.','.','.'}},
	spawn_col : 3}

var S_tet Tetramino = Tetramino{
	tet : [][]rune {{'.','g','g'},
          	 	    {'g','g','.'},
           		    {'.','.','.'}},
	spawn_col : 3}

	
var J_tet Tetramino = Tetramino{
	tet : [][]rune {{'b','.','.'},
            	    {'b','b','b'},
            	    {'.','.','.'}}, 
	spawn_col : 3}

var L_tet Tetramino = Tetramino{
	tet : [][]rune {{'.','.','o'}, 
          	 	    {'o','o','o'},
           		    {'.','.','.'}}, 
	spawn_col : 3}


var T_tet Tetramino = Tetramino{
	tet : [][]rune {{'.','m','.'},
            	    {'m','m','m'},
            	    {'.','.','.'}}, 
	spawn_col : 3}


var Testramino Tetramino = Tetramino{
	tet :  [][]rune {{'1','2','3','4'},
            	   {'5','6','7','8'},
            	   {'9','A','B','C'},
            	   {'D','E','F','G'}},
	spawn_col : 0}



//var O_tetramino = [][]rune {{'y','y'}, 
//							{'y','y'}}
//
//var Z_tetramino = [][]rune {{'r','r','.'},
//						  {'.','r','r'},
//						  {'.','.','.'}}
//
//var S_tetramino = [][]rune {{'.','g','g'},
//						    {'g','g','.'},
//						    {'.','.','.'}}
//
//var J_tetramino = [][]rune {{'b','.','.'},
//						    {'b','b','b'},
//						    {'.','.','.'}}
//
//var L_tetramino = [][]rune {{'.','.','o'},
//						    {'o','o','o'},
//						    {'.','.','.'}}
//
//var T_tetramino = [][]rune {{'.','m','.'},
//						    {'m','m','m'},
//						    {'.','.','.'}}
//
//var Testrimino = [][]rune {{'1','2','3','4'},
//						   {'5','6','7','8'},
//						   {'9','A','B','C'},
//						   {'D','E','F','G'}}
