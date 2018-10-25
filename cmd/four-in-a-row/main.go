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

	m := game.NewManager(width, height)
	m.Draw()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Type == termbox.EventKey {
				if m.CurrentGame.Winner != 0 {
					m.NewGame(m.CurrentGame.Winner)
					m.Draw()
				} else {
					m.CurrentGame.Input(int(ev.Ch) - 48)
					if m.CurrentGame.CurrentPlayer == 2 && *vsAI {
						move := ai.MakeMove(m.CurrentGame)
						m.CurrentGame.Input(move)
					}
				}
				m.Draw()
			}
		}
	}
}
