package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	t := newTicTacToe()
	fmt.Println("Welcome to tic tac toe game!")
	t.print()

	for {
		var position string
		var i int
		var j int
		var resetGame string

		fmt.Printf("Do your move '%s' player (input a position like 0,1)\n", t.currentMark)
		fmt.Scanf("%s", &position)

		arr := strings.Split(position, ",")
		i, _ = strconv.Atoi(arr[0])
		j, _ = strconv.Atoi(arr[1])

		t.addMove(i, j)
		t.print()

		if t.checkVictory(i, j) {
			fmt.Printf("Congratulations, '%s' player win!\n", t.currentMark)
			fmt.Printf("\nRestart the game? y/n \n")
			fmt.Scanf("%s", &resetGame)

			if resetGame == "y" {
				t.clean()
			} else {
				return
			}
		}

		t.changeTurn()
	}
}
