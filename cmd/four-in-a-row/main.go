package main

import (
	"flag"
	"os"

	"github.com/gokultp/go-four-in-a-row/pkg/ai"
	"github.com/gokultp/go-four-in-a-row/pkg/game"
	"github.com/nsf/termbox-go"
)

func main() {
	// initialise termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	// channel eventQueue for listening events
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// add flag option for ai
	vsAI := flag.Bool("ai", false, "play against ai")
	flag.Parse()

	// game size
	width := 10
	height := 10

	if *vsAI {
		width = 6
		height = 6
	}

	m := game.NewManager()
	g := game.NewGame(width, height, 0, 0)

	// splash screen only compatible with minimal 9x8 window
	if g.Height >= 9 || g.Width >= 8 {
		g.SplashScreen()
	} else {
		g = m.NewGame(g)
		g.Draw()
	}

	defer func() {
		if r := recover(); r != nil {
			f, err := os.Create("dump.txt")
			if err == nil {
				g.WriteState(f)
				f.Close()
			}
			panic(r)
		}
	}()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type != termbox.EventKey {
				continue
			}

			if ev.Key == termbox.KeyEsc {
				break loop
			}

			// reset game
			if g.Winner != 0 {
				g = m.NewGame(g)
				g.Draw()
				continue
			}

			// in the middle of a game
			g.Input(int(ev.Ch) - 48)
			if g.CurrentPlayer == 2 && *vsAI {
				move := ai.MakeMove(g)
				g.Input(move)
			}

			g.Draw()

		}
	}
}
