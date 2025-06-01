package game

import (
	"fmt"
	"go-game/utils"
)

type GameState struct {
	board        Board
	currentTurn  Color
	hasStarted   bool
	hasCheckmate bool
	timer        int
	moves        []Move
	message      string
}

type Move struct {
	from      string
	to        string
	isCapture bool
	isCastle  bool
}

func (gs GameState) InitGame(board Board) GameState {
	newGame := GameState{
		board:        board,
		currentTurn:  WHITE,
		hasStarted:   false,
		hasCheckmate: false,
		timer:        60,
		moves:        make([]Move, 40),
	}

	return newGame
}

func (gs *GameState) Move(fromRow int, fromCol int, toRow int, toCol int) {

	fromPiece := &gs.board[fromRow][fromCol]

	gs.message = "You have moved " + fromPiece.piece.String()

	utils.Log("Before"+fromPiece.piece.getIcon(fromPiece.color), utils.Cyan)

	toPiece := &gs.board[toRow][toCol]

	toPiece.piece = fromPiece.piece
	toPiece.color = fromPiece.color

	fromPiece.piece = EMPTY

	if gs.currentTurn == BLACK {
		gs.currentTurn = WHITE
	} else {
		gs.currentTurn = BLACK
	}
}

func ParseInput(input string) (fromRow, fromCol, toRow, toCol int, err error) {

	if len(input) != 4 {
		return 0, 0, 0, 0, fmt.Errorf("invalid input length: expected 4 characters like 'b2b4'")
	}

	columns := map[byte]int{
		'a': 0, 'b': 1, 'c': 2, 'd': 3,
		'e': 4, 'f': 5, 'g': 6, 'h': 7,
	}
	fromX := input[0]
	fromY := input[1]
	toX := input[2]
	toY := input[3]

	fromCol, ok1 := columns[fromX]
	toCol, ok2 := columns[toX]

	if !ok1 || !ok2 {
		return 0, 0, 0, 0, fmt.Errorf("invalid column character")
	}

	// ASCII math: '1' (49) → 8-1 = 7, '2' (50) → 8-2 = 6, ..., '8' (56) → 8-8 = 0
	fromRow = 8 - int(fromY-'0')
	toRow = 8 - int(toY-'0')

	return fromRow, fromCol, toRow, toCol, nil
}
