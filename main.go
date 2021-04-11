package main

import (
	"fmt"
	"strconv"
	"strings"
	"tic_tac_toe/game"
)

func main() {
	t := game.NewTicTacToe()
	fmt.Println("Welcome to tic tac toe game!")
	t.Print()

	for {
		var position string
		var i int
		var j int
		var resetGame string

		fmt.Printf("Do your move '%s' player (input a position like 0,1)\n", t.CurrentMark)
		fmt.Scanf("%s", &position)

		arr := strings.Split(position, ",")
		i, _ = strconv.Atoi(arr[0])
		j, _ = strconv.Atoi(arr[1])

		t.AddMove(i, j)
		t.Print()

		if t.CheckVictory(i, j) {
			fmt.Printf("Congratulations, '%s' player win!\n", t.CurrentMark)
			fmt.Printf("\nRestart the game? y/n \n")
			fmt.Scanf("%s", &resetGame)

			if resetGame == "y" {
				t.Clean()
			} else {
				return
			}
		}

		t.ChangeTurn()
	}
}
