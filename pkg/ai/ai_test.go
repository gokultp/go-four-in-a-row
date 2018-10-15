package ai

import (
	"strconv"
	"testing"

	"github.com/gokultp/go-four-in-a-row/pkg/game"
)

func TestMakeMove(t *testing.T) {
	g := game.NewGame(10, 10)
	move := MakeMove(g)
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

func TestSmartMove(t *testing.T) {
	width := 7
	height := 6
	depth := 7

	sampleBoardOneMove := [][]int{
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 0, 0},
		{1, 0, 0, 1, 2, 0, 2},
		{2, 1, 0, 2, 1, 0, 2},
		{2, 1, 2, 1, 2, 2, 2},
	}

	sampleBoardTwoMoves := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 2, 2, 0, 0},
	}

	position, mask, _ := convertBoard(sampleBoardOneMove)
	position2, mask2, _ := convertBoard(sampleBoardTwoMoves)

	path := makeSmartMove(position, mask, width, height, depth)

	if path.column != 2 {
		t.Fatalf("Failed to make winning move: %v should be 2", path.column)
	}

	if path.score != depth {
		t.Fatalf("Failed to make calculate move score: %v should be %v", path.score, depth)
	}

	path2 := makeSmartMove(position2, mask2, width, height, depth)

	if path2.column != 2 && path2.column != 5 {
		t.Fatalf("Failed to make best move: %v should be 2 or 5", path2.column)
	}
}

func TestRandomMove(t *testing.T) {
	width := 7
	height := 6

	sampleBoard := [][]int{
		{2, 1, 1, 0, 1, 2, 1},
		{1, 2, 2, 1, 2, 1, 2},
		{2, 1, 1, 2, 1, 2, 1},
		{1, 2, 2, 1, 2, 1, 2},
		{2, 1, 1, 2, 1, 2, 1},
		{1, 2, 2, 1, 2, 1, 2},
	}

	sampleBoard2 := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}

	_, mask, _ := convertBoard(sampleBoard)
	col, err := makeRandomMove(mask, width, height)

	_, mask2, _ := convertBoard(sampleBoard2)
	col2, err := makeRandomMove(mask2, width, height)

	if err != nil {
		t.Errorf("Unable to make random move: %s", err.Error())
	}

	if col != 3 {
		t.Errorf("Invalid random move: %v should be the only available spot is 3", col)
	}

	if col2 < 0 || col2 >= width {
		t.Errorf("Invalid random move: %v should be between 0 and %v", col, width)
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
	positionR, maskR, err := convertBoard(sampleBoardResult)
	_, maskR2, err := convertBoard(sampleBoardResultTwo)
	if err != nil {
		t.Errorf("Failed to converboard with error: %s", err.Error())
	}

	newMask, ok := makeMoveOne(mask, moveCol, height)
	newPosition2, newMask2, ok := makeMoveTwo(position, mask, moveCol, height)

	if newMask != maskR2 || ok != true {
		t.Errorf("Failed to make move: %v should equal %v", newMask, maskR)
	}

	if newPosition2 != positionR || ok != true {
		t.Errorf("Failed to make move: %v should equal %v", newPosition2, positionR)
	}

	if newMask2 != maskR || ok != true {
		t.Errorf("Failed to make move: %v should equal %v", newMask2, maskR)
	}
}

func TestValidMove(t *testing.T) {
	invalidMove, err := strconv.ParseInt("0000000000000000000000000000000000010000000000000", 2, 64)
	validMove, err := strconv.ParseInt("0000000000000000000000000001000000000000000000000", 2, 64)
	if err != nil {
		t.Errorf("Failed to parse bit string: %s", err.Error())
	}

	expectInvalid := isValidMove(invalidMove)
	expectValid := isValidMove(validMove)

	if expectInvalid == true {
		t.Errorf("Failed to determine valid move: %v should false", expectInvalid)
	}

	if expectValid != true {
		t.Errorf("Failed to determine valid move: %v should true", expectValid)
	}

	moveCol := 4
	height := uint64(6)

	sampleBoard := [][]int{
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}

	_, mask, _ := convertBoard(sampleBoard)
	_, ok := makeMoveOne(mask, moveCol, height)

	sampleBoard2 := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 2, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 2, 1, 2, 2, 0},
	}

	_, mask2, _ := convertBoard(sampleBoard2)
	_, ok2 := makeMoveOne(mask2, moveCol, height)

	if ok != false {
		t.Errorf("Failed to determine valid move: %v should false", ok)
	}

	if ok2 != true {
		t.Errorf("Failed to determine valid move: %v should true", ok)
	}
}
