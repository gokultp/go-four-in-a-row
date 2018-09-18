package main

import (
	"github.com/gokultp/go-four-in-a-row/game"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	g := game.NewGame(10, 10)
	g.Draw()

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Type == termbox.EventKey {
				g.Trigger()
				g.AddEntry(int(ev.Ch)-48, g.GetPlayer())
			}

		}
	}
}
