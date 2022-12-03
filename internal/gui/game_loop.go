package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameLoop interface {
	Start(w Window)
}

type gameLoop struct {
	isRunning bool
}

func NewGameLoop() GameLoop {
	return gameLoop{false}
}

func (gl gameLoop) Start(w Window) {
	gl.isRunning = true
	for gl.isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				break

			case *sdl.QuitEvent:
				gl.isRunning = false
				break
			}
		}

		w.Redraw()
		sdl.Delay(1000 / 60)
	}
}
