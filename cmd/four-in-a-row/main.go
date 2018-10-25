package main

import (
	"flag"

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
	g.Draw()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Type == termbox.EventKey {
				if g.Winner != 0 {
					g = m.NewGame(g)
					g.Draw()
				} else {
					g.Input(int(ev.Ch) - 48)
					if g.CurrentPlayer == 2 && *vsAI {
						move := ai.MakeMove(g)
						g.Input(move)
					}
				}
				g.Draw()
			}
		}
	}
}
