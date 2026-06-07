package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	menuSceneKey = iota
	gameSceneKey
	controlsSceneKey
)

type Scene interface {
	Update(drawTarget rl.RenderTexture2D, frame int64)
	ShouldExit() bool
}
