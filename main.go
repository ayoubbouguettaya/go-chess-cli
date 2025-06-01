package main

import (
	"bufio"
	"fmt"
	"go-game/game"
	"go-game/utils"
	"os"
	"strings"
)

func main() {

	board := game.InitialiseBoard()

	gameState := game.GameState{}.InitGame(board)

	game.DisplayBoard(gameState)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to CLI Chess!")
	for {
		fmt.Print("Enter your move (e.g., e2e4 or 'quit'): ")

		if !scanner.Scan() {
			// Exit on EOF or error
			fmt.Println("\nExiting game.")
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "quit" || input == "exit" {
			fmt.Println("Thanks for playing!")
			break
		}

		// TODO: Validate and apply the move
		fmt.Printf("You entered move: %s\n", input)

		// Update board state and display it

		fromRow, fromCol, toRow, toCol, err := game.ParseInput(input)

		if err != nil {
			utils.Log("Error", utils.Blue)
		}

		gameState.Move(fromRow, fromCol, toRow, toCol)

		game.DisplayBoard(gameState)

	}

}
