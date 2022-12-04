package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Loop struct {
	application *Application
	isRunning   bool
}

func NewLoop(app *Application) Loop {
	return Loop{app, false}
}

func (loop Loop) Start() {
	loop.isRunning = true
	for loop.isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				break

			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_CLOSE {
					loop.application.window.OnBeforeClose()
				}
				break

			case *sdl.QuitEvent:
				loop.isRunning = false
				break
			}
		}

		//loop.application.window.Redraw()
		//sdl.Delay(1000 / 60)
	}
}
