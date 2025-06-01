package game

type Piece int

const (
	KING Piece = iota
	QUEEN
	BISHOP
	KNIGHT
	CASTLE
	PAWN
	EMPTY
)

var PieceStateName = map[Piece]string{
	KING:   "KING",
	QUEEN:  "QUEEN",
	BISHOP: "BISHOP",
	KNIGHT: "KNIGHT",
	CASTLE: "CASTLE",
	PAWN:   "PAWN",
	EMPTY:  "EMPTY",
}

func (ss Piece) String() string {
	return PieceStateName[ss]
}

func (ss Piece) getIcon(color Color) string {
	switch ss {
	case BISHOP:
		if color == BLACK {
			return "♝"
		} else {
			return "♗"
		}
	case CASTLE:
		if color == BLACK {
			return "♜"
		} else {
			return "♖"
		}

	case PAWN:
		if color == BLACK {
			return "♟"
		} else {
			return "♙"
		}

	case KNIGHT:
		if color == BLACK {
			return "♞"

		} else {
			return "♘"
		}

	case KING:
		if color == BLACK {
			return "♚"
		} else {
			return "♔"
		}

	case QUEEN:
		if color == BLACK {
			return "♛"
		} else {
			return "♕"
		}
	case EMPTY:
		return " "
	}

	return " "
}

type Color int

const (
	BLACK Color = iota
	WHITE
)

var ColorStateName = map[Color]string{
	BLACK: "BLACK",
	WHITE: "WHITE",
}

func (ss Color) String() string {
	return ColorStateName[ss]
}

type Square struct {
	piece     Piece
	positionX string
	positionY int
	color     Color
}

type Board [8][8]Square

func InitialiseBoard() Board {
	var board [8][8]Square

	columns := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	blackBackRank := [8]Piece{CASTLE, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, CASTLE}

	for col := 0; col < 8; col++ {
		board[0][col] = Square{
			piece:     blackBackRank[col],
			positionX: columns[col],
			positionY: 8,
			color:     BLACK,
		}
	}

	// Place black pawns (row 7)
	for col := 0; col < 8; col++ {
		board[1][col] = Square{
			piece:     PAWN,
			positionX: columns[col],
			positionY: 7,
			color:     BLACK,
		}
	}

	// Place white pawns (row 2)
	for col := 0; col < 8; col++ {
		board[6][col] = Square{
			piece:     PAWN,
			positionX: columns[col],
			positionY: 2,
			color:     WHITE,
		}
	}

	// Place white major pieces (row 1)
	whiteBackRank := [8]Piece{CASTLE, KNIGHT, BISHOP, QUEEN, KING, BISHOP, KNIGHT, CASTLE}
	for col := 0; col < 8; col++ {
		board[7][col] = Square{
			piece:     whiteBackRank[col],
			positionX: columns[col],
			positionY: 1,
			color:     WHITE,
		}
	}

	// Empty squares
	for row := 2; row < 6; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = Square{
				piece:     -1, // No piece
				positionX: columns[col],
				positionY: 8 - row,
			}
		}
	}

	return board
}
