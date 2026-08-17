package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int32

type GameScene struct {
	a            *Application
	field        *Field
	digger       *Digger
	fire         *Fire
	emeraldsPool *EmeraldsPool
	bagsPool     *BagsPool
	monstersPool *MonstersPool
	debugGrid    *DebugGrid
	moveGrid     *MoveGrid
	isStarted    bool
	level        int32
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	gameScene.field = NewField(&gameScene)
	gameScene.moveGrid = NewMoveGrid(&gameScene)
	gameScene.digger = NewDigger(&gameScene)
	gameScene.fire = NewFire(&gameScene)
	gameScene.emeraldsPool = NewEmeraldsPool(&gameScene)
	gameScene.bagsPool = NewBagsPool(&gameScene)
	gameScene.monstersPool = NewMonstersPool(&gameScene)
	gameScene.debugGrid = NewDebugGrid(&gameScene)
	gameScene.isStarted = false
	gameScene.level = LevelPlan(1)
	return &gameScene
}

func (gs *GameScene) ProcessInput() {
	gs.digger.shouldMove = false
	for k, v := range keysToDirectionsMap {
		if rl.IsKeyDown(k) {
			gs.digger.requestedDirection = v
			gs.digger.shouldMove = true
		}
	}
}

func (gs *GameScene) Update(tick int64) {
	gs.debugGrid.Update(tick)
	gs.moveGrid.Update(tick)
	gs.field.Update(tick)
	gs.emeraldsPool.Update(tick)
	gs.bagsPool.Update(tick)
	gs.fire.Update(tick)
	gs.monstersPool.Update(tick)
	gs.digger.Update(tick)
}

func (gs *GameScene) Render(drawTarget rl.RenderTexture2D) {
	gs.field.Render(drawTarget)
	gs.emeraldsPool.Render(drawTarget)
	gs.digger.Render(drawTarget)
	gs.fire.Render(drawTarget)
	gs.monstersPool.Render(drawTarget)
	gs.bagsPool.Render(drawTarget)
	gs.debugGrid.Render(drawTarget)
	gs.moveGrid.Render(drawTarget)
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}
