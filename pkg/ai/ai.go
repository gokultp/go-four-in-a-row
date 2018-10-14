package ai

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/berto/go-four-in-a-row/pkg/game"
	"github.com/pkg/errors"
)

// C4AI represents the AI for connect four
type C4AI struct{}

type pathNode struct {
	column  int
	score   int
	invalid bool
}

// MakeMove takes a game and returns a valid move
func (ai *C4AI) MakeMove(g *game.Game) int {
	position, mask, err := convertBoard(g.State)
	if err != nil {
		log.Fatalf("Failed to convert board with error: %s", err.Error())
	}
	depth := 5
	path := makeSmartMove(position, mask, g.Width, g.Height, depth)
	return path.column
}

func makeSmartMove(position int64, mask int64, width int, height int, depth int) pathNode {
	paths := make([]pathNode, width)

	uHeight := uint64(height)
	isAIMove := depth%2 == 1

	var wg sync.WaitGroup
	wg.Add(width)
	for i := 0; i < width; i++ {
		go func(position int64, mask int64, depth int, i int) {
			defer wg.Done()
			change, score, ok := position, depth, true

			if isAIMove {
				position, mask, ok = makeMoveTwo(position, mask, i, uHeight)
				change = position
			} else {
				mask, ok = makeMoveOne(mask, i, uHeight)
				change = position ^ mask
				score = depth * -1
			}
			win := connectFour(change, uHeight)
			if !ok {
				paths[i] = pathNode{score: 0, column: i, invalid: true}
			} else if !win && depth > 1 {
				path := makeSmartMove(position, mask, width, height, depth-1)
				path.column = i
				paths[i] = path
			} else {
				paths[i] = pathNode{score: score, column: i}
			}
		}(position, mask, depth, i)
	}

	wg.Wait()

	score := 0
	column := 0
	var validColumns []int

	for i, path := range paths {
		if !path.invalid {
			if i == 0 {
				score = path.score
			}
			validColumns = append(validColumns, path.column)
			if path.score != 0 && (isAIMove && path.score > score || !isAIMove && path.score < score) {
				score = path.score
				column = path.column
			}
		}
	}

	validMoves := len(validColumns)
	if score == 0 && validMoves > 0 {
		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)
		randomColNum := random.Intn(validMoves)
		column = validColumns[randomColNum]
	}

	return pathNode{score: score, column: column}
}

func makeRandomMove(mask int64, width int, height int) (int, error) {
	triedMoves := make(map[int]bool)
	count := 0
	for {
		randomColNum, ok := randomCol(triedMoves, width, count)
		if !ok {
			return 0, errors.New("No moves possible")
		}
		_, ok = makeMoveOne(mask, randomColNum, uint64(height))
		if !ok {
			triedMoves[randomColNum] = true
			count++
		} else {
			return randomColNum, nil
		}
	}
}

func randomCol(tried map[int]bool, width int, count int) (int, bool) {
	if count >= width {
		return 0, false
	}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	randomColNum := random.Intn(width)
	if tried[randomColNum] == true {
		return randomCol(tried, width, count)
	}
	return randomColNum, true
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

func makeMoveTwo(position int64, mask int64, col int, height uint64) (int64, int64, bool) {
	newMask := mask | (mask + (1 << (uint64(col) * (height + 1))))
	if !isValidMove((newMask ^ mask)) {
		return 0, 0, false
	}
	newPosition := position | (newMask ^ mask)
	return newPosition, newMask, true
}

func isValidMove(position int64) bool {
	bitString := strconv.FormatInt(int64(position), 2)
	if len(bitString)%7 == 0 {
		return false
	}
	return true
}
