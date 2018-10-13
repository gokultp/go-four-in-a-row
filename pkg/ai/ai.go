package ai

import (
	"strconv"

	"github.com/gokultp/go-four-in-a-row/pkg/game"
	"github.com/pkg/errors"
)

// C4AI represents the AI for connect four
type C4AI struct{}

// MakeMove takes a game and returns a valid move
func (ai *C4AI) MakeMove(g *game.Game) int {
	return 0
}

func convertBoard(board [][]int) (int64, int64, error) {
	position := ""
	mask := ""

	board = transposeRight(board)

	for _, row := range board {
		for _, col := range row {
			switch col {
			case 1:
				position += "1"
				mask += "1"
			case 2:
				mask += "1"
				position += "0"
			default:
				mask += "0"
				position += "0"
			}
		}
	}

	positionInt, err := strconv.ParseInt(position, 2, 64)

	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to parse position bit")
	}

	maskInt, err := strconv.ParseInt(mask, 2, 64)

	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to parse position bit")
	}

	return positionInt, maskInt, nil
}

func transposeRight(board [][]int) [][]int {
	x := len(board)
	y := len(board[0])
	newBoard := make([][]int, y)
	for i, row := range board {
		for j, col := range row {
			if len(newBoard[j]) == 0 {
				newBoard[j] = make([]int, x)
			}
			newBoard[j][len(board)-1-i] = col
		}
	}
	return newBoard
}
