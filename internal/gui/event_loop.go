package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type EventLoop struct {
	application *Application
}

func NewEventLoop(app *Application) *EventLoop {
	return &EventLoop{app}
}

func (eventLoop EventLoop) Run() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				eventLoop.application.window.OnBeforeClose()
			}
			break

		case *sdl.QuitEvent:
			eventLoop.application.loop.isRunning.UnSet()
			break
		}
	}
}
