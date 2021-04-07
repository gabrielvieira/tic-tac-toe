package main

import "fmt"

type TicTacToe struct {
	board       [3][3]string
	currentMark string
}

func newTicTacToe() TicTacToe {
	board := [3][3]string{}
	t := TicTacToe{board: board, currentMark: "X"}
	t.clean()
	return t
}

func (t *TicTacToe) clean() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.board[i][j] = " "
		}
	}
}

func (t *TicTacToe) addMove(i, j int) {
	t.board[i][j] = t.currentMark
}

func (t *TicTacToe) changeTurn() {
	if t.currentMark == "X" {
		t.currentMark = "O"
	} else {
		t.currentMark = "X"
	}
}

func (t TicTacToe) checkVictory(i, j int) bool {

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

func (t TicTacToe) print() {

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
