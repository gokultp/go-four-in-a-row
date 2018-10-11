package main

import (
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

	g := game.NewGame(10, 10)
	g.Draw()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Type == termbox.EventKey {
				if g.Winner != 0 {
					g.Cancel()
					g = game.NewGame(10, 10)
					g.Draw()
				} else {
					g.Input(int(ev.Ch) - 48)
				}
			}
		}
	}
}
