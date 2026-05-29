package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_LOGICAL_WIDTH  = 320
	SCREEN_LOGICAL_HEIGHT = 240
	CELLS_HORIZONTAL      = 15
	CELLS_VERTICAL        = 10
	CELL_WIDTH            = 20
	CELL_HEIGHT           = 20
	CELLS_OFFSET          = 10
	FIELD_OFFSET_Y        = 20
)

type GameScene struct {
	a         *Application
	field     *Field
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.field = NewField(a)
	gameScene.isStarted = false
	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	gs.field.Update()
	rl.EndDrawing()
}
