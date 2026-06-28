package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	menuSceneKey = iota
	gameSceneKey
	controlsSceneKey
)

type Scene interface {
	ProcessInput()
	Update(tick int64)
	Render(drawTarget rl.RenderTexture2D)
	ShouldExit() bool
}
