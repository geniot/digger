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
	gameScene.field = NewField(&gameScene)
	gameScene.digger = NewDigger(&gameScene)
	gameScene.debugGrid = NewDebugGrid(&gameScene)
	gameScene.isStarted = false
	return &gameScene
}

func (gs *GameScene) ProcessInput() {
	gs.digger.shouldMove = false
	for k, v := range keysToDirectionsMap {
		if rl.IsKeyDown(k) {
			gs.digger.direction = v
			gs.digger.shouldMove = true
		}
	}
}

func (gs *GameScene) Update(tick int64) {
	gs.field.Update(tick)
	gs.digger.Update(tick)
	gs.debugGrid.Update(tick)
}

func (gs *GameScene) Render(drawTarget rl.RenderTexture2D) {
	gs.field.Render(drawTarget)
	gs.digger.Render(drawTarget)
	gs.debugGrid.Render(drawTarget)
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}
