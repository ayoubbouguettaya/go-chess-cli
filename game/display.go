package game

import (
	"fmt"
	"go-game/utils"
)

func DisplayBoard(gameState GameState) {

	board := gameState.board
	utils.ClearConsole()

	fmt.Println("  a b c d e f g h")
	for i, row := range board {
		fmt.Printf("%d ", 8-i)
		for _, square := range row {
			utils.Log(square.piece.getIcon(square.color), utils.Cyan)
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("  a b c d e f g h")

	utils.Log("===================================================\n", utils.Green)

	utils.Log("current turn : "+gameState.currentTurn.String()+" | "+gameState.message+"\n", utils.Green)

	utils.Log("===================================================\n", utils.Green)

}
