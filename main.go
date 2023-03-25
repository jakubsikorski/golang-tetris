package main

import (
	"math/rand"
	"tetris/screen"
	"tetris/tetris"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	const animationSpeed = 20 + time.Millisecond
	ticker := time.NewTimer(time.Duration(animationSpeed))

	game := tetris.NewGame()
	scr := screen.New()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					game.Rotate()
				case ev.Key == termbox.KeyArrowDown:
					game.SpeedUp()
				case ev.Key == termbox.KeyArrowLeft:
					game.MoveLeft()
				case ev.Key == termbox.KeyArrowRight:
					game.MoveRight()
				case ev.Key == termbox.KeySpace:
					game.Drop()
				// quit
				case ev.Ch == 'q':
					return
				// start
				case ev.Ch == 's':
					game.Start()
				}
			}
		case <-ticker.C:
			// scr.RenderAscii(game.GetBoard())
			scr.Render(game.GetBoard(), game.Score)
			ticker.Reset(animationSpeed)
		case <-game.FallSpeed.C:
			game.GameLoop()
		}

	}
}
