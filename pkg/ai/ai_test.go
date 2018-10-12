package ai

import (
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
