package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	a                   *Application
	field               *Field
	digger              *Digger
	fire                *Fire
	emeraldsPool        *EmeraldsPool
	bagsPool            *BagsPool
	monstersPool        *MonstersPool
	debugGrid           *DebugGrid
	moveGrid            *MoveGrid
	isStarted           bool
	level               int32
	keysToDirectionsMap map[int32]Direction
}

func NewGameScene(a *Application) *GameScene {
	gs := GameScene{}
	gs.a = a
	gs.field = NewField(&gs)
	gs.moveGrid = NewMoveGrid(&gs)
	gs.digger = NewDigger(&gs)
	gs.fire = NewFire(&gs)
	gs.emeraldsPool = NewEmeraldsPool(&gs)
	gs.bagsPool = NewBagsPool(&gs)
	gs.monstersPool = NewMonstersPool(&gs)
	gs.debugGrid = NewDebugGrid(&gs)
	gs.isStarted = false
	gs.level = LevelPlan(1)
	gs.keysToDirectionsMap = map[int32]Direction{
		rl.KeyLeft:  LEFT,
		rl.KeyRight: RIGHT,
		rl.KeyUp:    UP,
		rl.KeyDown:  DOWN,
	}
	return &gs
}

func (gs *GameScene) ProcessInput() {
	gs.digger.shouldMove = false
	gs.fire.shouldShoot = false
	for k, v := range gs.keysToDirectionsMap {
		if rl.IsKeyDown(k) {
			gs.digger.requestedDirection = v
			gs.digger.shouldMove = true
		}
	}
	if rl.IsKeyDown(rl.KeyF) {
		gs.fire.shouldShoot = true
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
