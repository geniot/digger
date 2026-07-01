package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int32

type GameScene struct {
	a         *Application
	field     *Field
	digger    *Digger
	debugGrid *DebugGrid
	moveGrid  *MoveGrid
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.field = NewField(&gameScene)
	gameScene.digger = NewDigger(&gameScene)
	gameScene.debugGrid = NewDebugGrid(&gameScene)
	gameScene.moveGrid = NewMoveGrid(&gameScene)
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
	gs.moveGrid.Update(tick)
}

func (gs *GameScene) Render(drawTarget rl.RenderTexture2D) {
	gs.field.Render(drawTarget)
	gs.digger.Render(drawTarget)
	gs.debugGrid.Render(drawTarget)
	gs.moveGrid.Render(drawTarget)
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}
