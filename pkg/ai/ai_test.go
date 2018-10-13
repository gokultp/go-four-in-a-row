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
		t.Errorf("Invalid move: %v. Should be less than %v, more than 0", move, g.Width)
	}

	// expect playerBeforeMove to not be the same player after move
	if playerBeforeMove == playerAfterMove {
		t.Errorf("Invalid move: %v should not equal %v", playerBeforeMove, playerAfterMove)
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
		t.Errorf("Failed to converboard with error: %s", err.Error())
	}

	expectedPosition, err := strconv.ParseInt("0000000000000000010100000111000000000000000000000", 2, 64)
	if err != nil || position != expectedPosition {
		t.Errorf("Invalid convertion: position %v should equal %v", position, expectedPosition)
	}
	expectedMask, err := strconv.ParseInt("0000000000000100011110000111000000100000000000000", 2, 64)

	if err != nil || mask != expectedMask {
		t.Errorf("Invalid convertion: mask %v should equal %v", mask, expectedMask)
	}
}

func TestConnectFour(t *testing.T) {
	height := uint64(6)

	sampleVertical := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{2, 0, 2, 1, 2, 2, 0},
	}
	sampleDiagonalRight := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 2, 1, 0, 0},
		{0, 0, 2, 1, 2, 0, 0},
		{2, 2, 1, 1, 1, 0, 0},
		{2, 2, 2, 1, 2, 2, 1},
	}
	sampleDiagonalLeft := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 2, 0, 0},
		{0, 0, 2, 1, 2, 0, 0},
		{0, 2, 1, 2, 1, 0, 0},
		{2, 1, 2, 1, 2, 1, 0},
	}
	sampleHorizontal := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 2, 0, 0},
		{0, 0, 2, 1, 2, 0, 0},
		{0, 2, 1, 2, 1, 0, 0},
		{2, 1, 1, 1, 1, 2, 0},
	}
	sampleNone := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}
	positionV, maskV, err := convertBoard(sampleVertical)
	expectedV := connectFour(positionV, height)
	expectedTwoV := connectFour(positionV^maskV, height)

	positionDR, maskDR, err := convertBoard(sampleDiagonalRight)
	expectedDR := connectFour(positionDR, height)
	expectedTwoDR := connectFour(positionDR^maskDR, height)

	positionDL, maskDL, err := convertBoard(sampleDiagonalLeft)
	expectedDL := connectFour(positionDL, height)
	expectedTwoDL := connectFour(positionDL^maskDL, height)

	positionH, maskH, err := convertBoard(sampleHorizontal)
	expectedH := connectFour(positionH, height)
	expectedTwoH := connectFour(positionH^maskH, height)

	positionN, maskN, err := convertBoard(sampleNone)
	expectedN := connectFour(positionN, height)
	expectedTwoN := connectFour(positionN^maskN, height)

	if err != nil {
		t.Errorf("Failed to converboard with error: %s", err.Error())
	}

	if expectedV != true || expectedTwoV != false {
		t.Errorf("Failed to find vertical connect four: %v should be true and %v should be false", expectedV, expectedTwoV)
	}

	if expectedDR != false || expectedTwoDR != true {
		t.Errorf("Failed to find diagonal right connect four: %v should be false and %v should be true", expectedDR, expectedTwoDR)
	}

	if expectedDL != true || expectedTwoDL != false {
		t.Errorf("Failed to find diagonal left connect four: %v should be true and %v should be false", expectedDL, expectedTwoDL)
	}

	if expectedH != true || expectedTwoH != false {
		t.Errorf("Failed to find horizontal connect four: %v should be true and %v should be false", expectedH, expectedTwoH)
	}

	if expectedN != false || expectedTwoN != false {
		t.Errorf("Failed to find connect four: %v should be false and %v should be false", expectedN, expectedTwoN)
	}
}

func TestAddMove(t *testing.T) {
	moveCol := 2
	height := uint64(6)
	sampleBoard := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}
	sampleBoardResult := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}
	sampleBoardResultTwo := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 2, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}
	position, mask, err := convertBoard(sampleBoard)
	positionR, _, err := convertBoard(sampleBoardResult)
	_, maskR, err := convertBoard(sampleBoardResultTwo)
	if err != nil {
		t.Errorf("Failed to converboard with error: %s", err.Error())
	}

	newMask := makeMoveOne(mask, moveCol, height)
	newPosition := makeMoveTwo(position, mask, moveCol, height)

	if newMask != maskR {
		t.Errorf("Failed to make move: %v should equal %v", newMask, maskR)
	}

	if newPosition != positionR {
		t.Errorf("Failed to make move: %v should equal %v", newPosition, positionR)
	}
}
