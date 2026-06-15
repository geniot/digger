package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int32

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

var (
	keysToDirectionsMap = map[int32]Direction{
		rl.KeyLeft:  LEFT,
		rl.KeyRight: RIGHT,
		rl.KeyUp:    UP,
		rl.KeyDown:  DOWN,
	}
)

const (
	SCREEN_LOGICAL_WIDTH  = int32(320)
	SCREEN_LOGICAL_HEIGHT = int32(240)
	CELLS_HORIZONTAL      = int32(15)
	CELLS_VERTICAL        = int32(10)
	CELL_WIDTH            = int32(20)
	CELL_HEIGHT           = int32(18)
	FIELD_OFFSET_X        = int32(10)
	FIELD_OFFSET_Y        = int32(2)
)

type GameScene struct {
	a         *Application
	field     *Field
	digger    *Digger
	debugGrid *DebugGrid
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.field = NewField(a)
	gameScene.digger = NewDigger(a)
	gameScene.debugGrid = NewDebugGrid(a)
	gameScene.isStarted = false
	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(drawTarget rl.RenderTexture2D, frame int64) {
	gs.handleUserInput()
	gs.field.Update(drawTarget, frame)
	gs.digger.Update(drawTarget, frame)
	gs.debugGrid.Update(drawTarget, frame)
}

func (gs *GameScene) handleUserInput() {
	gs.digger.shouldMove = false
	gs.digger.direction = NONE
	if direction, ok := keysToDirectionsMap[rl.GetKeyPressed()]; ok {
		gs.digger.direction = direction
		gs.digger.shouldMove = true
	}
}
