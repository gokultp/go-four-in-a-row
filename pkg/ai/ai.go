package ai

import (
	"strconv"

	"github.com/gokultp/go-four-in-a-row/pkg/game"
	"github.com/pkg/errors"
)

// C4AI represents the AI for connect four
type C4AI struct{}

// MakeMove takes a game and returns a valie move
func (ai *C4AI) MakeMove(g *game.Game) int {
	return 0
}

func convertBoard(board [][]int) (int64, int64, error) {
	position := ""
	mask := ""

	for _, row := range board {
		for _, num := range row {
			switch num {
			case 1:
				position += "1"
				mask += "1"
			case 2:
				mask += "1"
			default:
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
