package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameScene struct {
	a *Application
	//cube      *Cube
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	//gameScene.cube = NewCube(3, split(CubeCorrect), a)
	gameScene.isStarted = false
	//gameScene.Reset()
	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(camera *rl.Camera) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}
