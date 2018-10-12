package ai

import (
	"strconv"
	"testing"

	"github.com/gokultp/go-four-in-a-row/pkg/game"
)

func TestMakeMove(t *testing.T) {
	g := game.NewGame(10, 10)
	ai := C4AI{}
	move := ai.MakeMove(g)
	playerBeforeMove := g.CurrentPlayer
	g.Input(move)
	playerAfterMove := g.CurrentPlayer

	// expect move to be a valid column number
	if move < 0 || move >= g.Width {
		t.Fatalf("Invalid move: %v. Should be less than %v, more than 0", move, g.Width)
	}

	// expect playerBeforeMove to not be the same player after move
	if playerBeforeMove == playerAfterMove {
		t.Fatalf("Invalid move: %v should not equal %v", playerBeforeMove, playerAfterMove)
	}
}

func TestConvertBoard(t *testing.T) {
	sampleBoard := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}

	position, mask, err := convertBoard(sampleBoard)
	if err != nil {
		t.Fatalf("Failed to converboard with error: %s", err.Error())
	}

	expectedPosition, err := strconv.ParseInt("000000000000000000111000010100000000000000", 2, 64)
	if err != nil || position != expectedPosition {
		t.Fatalf("Invalid convertion: position %v should equal %v", position, expectedPosition)
	}
	expectedMask, err := strconv.ParseInt("000000000000100000111000111100100000000000", 2, 64)

	if err != nil || mask != expectedMask {
		t.Fatalf("Invalid convertion: mask %v should equal %v", mask, expectedMask)
	}
}
