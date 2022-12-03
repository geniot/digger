package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameLoop interface{}

type gameLoop struct {
	isRunning bool
}

func New() GameLoop {
	return gameLoop{false}
}

func (gl gameLoop) Start() {
	gl.isRunning = true
	for gl.isRunning {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.QuitEvent:
				gl.isRunning = false
				break
			}
		}
	}
}
