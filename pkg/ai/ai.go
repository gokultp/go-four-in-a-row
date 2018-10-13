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
				position = "1" + position
				mask = "1" + mask
			case 2:
				mask = "1" + mask
				position = "0" + position
			default:
				mask = "0" + mask
				position = "0" + position
			}
		}
		// adds a buffer row for checking false positives
		mask = "0" + mask
		position = "0" + position
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

func connectFour(position int64, height uint64) bool {
	// Horizontal check
	m := position & (position >> (height + 1))
	if m&(m>>14) != 0 {
		return true
	}
	// Diagonal Left
	m = position & (position >> height)
	if m&(m>>12) != 0 {
		return true
	}
	// Diagonal Right
	m = position & (position >> (height + 2))
	if m&(m>>16) != 0 {
		return true
	}
	// Vertical
	m = position & (position >> 1)
	if m&(m>>2) != 0 {
		return true
	}
	// Nothing found
	return false
}

func makeMoveOne(mask int64, col int, height uint64) (int64, bool) {
	newMask := mask | (mask + (1 << (uint64(col) * (height + 1))))
	if !isValidMove(mask ^ newMask) {
		return 0, false
	}
	return newMask, true
}

func makeMoveTwo(position int64, mask int64, col int, height uint64) (int64, bool) {
	newMask := mask | (mask + (1 << (uint64(col) * (height + 1))))
	if !isValidMove((newMask ^ mask)) {
		return 0, false
	}
	newPosition := position | (newMask ^ mask)
	return newPosition, true
}

func isValidMove(position int64) bool {
	bitString := strconv.FormatInt(int64(position), 2)
	if len(bitString)%7 == 0 {
		return false
	}
	return true
}
