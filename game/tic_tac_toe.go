package game

import "fmt"

type ticTacToe struct {
	board       [3][3]string
	CurrentMark string
}

func NewTicTacToe() ticTacToe {
	board := [3][3]string{}
	t := ticTacToe{board: board, CurrentMark: "X"}
	t.Clean()
	return t
}

func (t *ticTacToe) Clean() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.board[i][j] = " "
		}
	}
}

func (t *ticTacToe) AddMove(i, j int) {
	t.board[i][j] = t.CurrentMark
}

func (t *ticTacToe) ChangeTurn() {
	if t.CurrentMark == "X" {
		t.CurrentMark = "O"
	} else {
		t.CurrentMark = "X"
	}
}

func (t ticTacToe) CheckVictory(i, j int) bool {

	if t.board[i][0] == t.board[i][1] && t.board[i][1] == t.board[i][2] {
		return true
	}

	if t.board[0][j] == t.board[1][j] && t.board[1][j] == t.board[2][j] {
		return true
	}

	if i == j && (t.board[0][0] == t.board[1][1]) && (t.board[1][1] == t.board[2][2]) {
		return true
	}

	if (i+j == 2) && (t.board[0][2] == t.board[1][1]) && (t.board[1][1] == t.board[2][0]) {
		return true
	}

	return false
}

func (t ticTacToe) Print() {

	/*
		create a print like this
		X | X | O
		---------
		X | X | O
		---------
		X | X | O

	*/
	fmt.Printf("\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				fmt.Printf("%s |", t.board[i][j])
			}

			if j == 1 {
				fmt.Printf(" %s |", t.board[i][j])
			}

			if j == 2 {
				fmt.Printf(" %s\n", t.board[i][j])
				if i != 2 {
					fmt.Println("---------")
				}
			}
		}
	}
	fmt.Printf("\n")
}
