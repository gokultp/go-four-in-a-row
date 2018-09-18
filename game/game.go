package game

import (
	"context"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var PlayerChar = []string{"\\//\\", "/\\\\/"}

type Game struct {
	State         [][]int
	CurrentPlayer int
	Width         int
	Height        int
	offsetX       int
	offsetY       int
	Winner        int
	players       []string
	wonState      [][]int
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewGame(width, height int) *Game {

	state := [][]int{}
	for i := 0; i < height; i++ {
		rowState := make([]int, width)
		state = append(state, rowState)
	}
	ctx, cancel := context.WithCancel(context.Background())

	game := &Game{
		Width:         width,
		Height:        height,
		State:         state,
		CurrentPlayer: 1,
		ctx:           ctx,
		cancel:        cancel,
	}
	game.getOffset()
	return game
}

func (g *Game) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for x := 0; x < g.Width; x++ {
		termbox.SetCell(g.offsetX+x*3, g.offsetY-2, rune(48+x), termbox.ColorYellow, termbox.ColorDefault)
	}
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.setContent(x, y, true)
		}
	}
	termbox.Flush()
}

func (g *Game) getOffset() {
	sw, sh := termbox.Size()
	g.offsetX = (sw - g.Width*3) / 2
	g.offsetY = (sh - g.Height*2) / 2
}

func (g *Game) setContent(x, y int, show bool) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 1; j++ {
			ch, fore, bg := getplayerDisplayPropsLarge(g.State[y][x], i, j)
			if show {
				termbox.SetCell((g.offsetX + x*3 + i), (g.offsetY + y*2 + j), ch, fore, bg)
			} else {
				termbox.SetCell((g.offsetX + x*3 + i), (g.offsetY + y*2 + j), ch, fore, termbox.ColorDefault)
			}
		}
	}

	// ch, fore, bg := getplayerDisplayProps(g.State[y][x])
	// termbox.SetCell((g.offsetX + x*3), (g.offsetY + y*2), ch, fore, bg)
}
func getplayerDisplayProps(player int) (rune, termbox.Attribute, termbox.Attribute) {
	if player == 1 {
		return 'X', termbox.ColorRed, termbox.ColorDefault
	}
	if player == 2 {
		return 'O', termbox.ColorBlue, termbox.ColorDefault
	}
	return ' ', termbox.ColorDefault, termbox.ColorBlack
}
func getplayerDisplayPropsLarge(player, col, row int) (rune, termbox.Attribute, termbox.Attribute) {
	if player == 1 {
		return ' ', termbox.ColorDefault, termbox.ColorRed
	}
	if player == 2 {
		return ' ', termbox.ColorDefault, termbox.ColorBlue
	}
	return ' ', termbox.ColorDefault, termbox.ColorBlack
}

func (g *Game) AddEntry(col, player int) {
	if g.Winner != 0 {
		return
	}
	if g.State[0][col] != 0 {
		return
	}
	i := 0
	for ; i < g.Height && g.State[i][col] == 0; i++ {
		g.State[i][col] = player
		if i > 0 {
			g.State[i-1][col] = 0
		}
		g.Draw()
		time.Sleep(30 * time.Millisecond)
	}
	g.togglePlayer()
	if g.isWon(col, i-1, player) {
		g.declareWinner()
	}
}

func (g *Game) isWon(col, row, player int) bool {
	// check vertically
	wonState := make([][]int, 0)
	count := 0
	for i := row; i < g.Width && g.State[i][col] == player; i++ {
		count++
		wonState = append(wonState, []int{i, col})
	}
	if count >= 4 {
		g.Winner = player
		g.wonState = wonState
		return true
	}
	// check horizontally left
	wonState = make([][]int, 0)
	count = 0
	i := col
	for ; i > 0 && g.State[row][i] == player; i-- {
	}
	i++
	// check horizontally right
	for ; i < g.Width && g.State[row][i] == player; i++ {
		wonState = append(wonState, []int{row, i})
		count++
	}
	if count >= 4 {
		g.Winner = player
		g.wonState = wonState
		return true
	}

	// check positive diagonal
	wonState = make([][]int, 0)
	i = 0
	count = 0
	for ; col-i > 0 && row-i > 0 && g.State[row-i][col-i] == player; i++ {
	}
	i--

	for ; col-i < g.Width && row-i < g.Height && g.State[row-i][col-i] == player; i-- {
		wonState = append(wonState, []int{row - i, col - i})
		count++
	}
	if count >= 4 {
		g.wonState = wonState
		g.Winner = player
		return true
	}

	// check negative diagonal
	wonState = make([][]int, 0)
	i = 0
	count = 0
	for ; col+i < g.Width && row-i > 0 && g.State[row-i][col+i] == player; i++ {
	}
	i--
	for ; col+i > 0 && row-i < g.Height && g.State[row-i][col+i] == player; i-- {
		wonState = append(wonState, []int{row - i, col + i})
		count++
	}
	if count >= 4 {
		g.wonState = wonState
		g.Winner = player
		return true
	}
	return false
}

func (g *Game) declareWinner() {
	if g.Winner != 0 {
		g.renderText(12, "% won the game")
	}
	go func(g *Game) {
		show := true
	inf_loop:
		for {
			select {
			case <-g.ctx.Done():
				g = NewGame(g.Width, g.Height)
				g.Draw()
				break inf_loop
			default:
			}
			for i := 0; i < 4; i++ {
				g.setContent(g.wonState[i][1], g.wonState[i][0], show)
			}
			show = !show
			termbox.Flush()
			time.Sleep(time.Millisecond * 500)
		}
	}(g)
}

func (g *Game) renderText(row int, text string) {
	x, y := g.offsetX, g.offsetY+row*2

	for i := 0; i < len(text); i++ {
		if text[i] == byte('%') {
			ch, fore, bg := getplayerDisplayProps(g.Winner)
			termbox.SetCell(x, y+i, ch, fore, bg)
		} else {
			termbox.SetCell(x+i, y, rune(text[i]), termbox.ColorWhite, termbox.ColorDefault)
		}
	}
	termbox.Flush()

}

func (g *Game) GetPlayer() int {
	return g.CurrentPlayer
}

func (g *Game) togglePlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
		return
	}
	g.CurrentPlayer = 1
}

func (g *Game) Trigger() {
	if g.Winner != 0 {
		g.cancel()
	}
}
