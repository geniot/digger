package main

const (
	menuSceneKey = iota
	gameSceneKey
	tutorialSceneKey
	controlsSceneKey
)

type Scene interface {
	Update()
	ShouldExit() bool
}
