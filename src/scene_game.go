package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	a *Application

	scoreDrawTarget rl.RenderTexture2D
	scoreSourceRect rl.Rectangle
	scoreDestRect   rl.Rectangle

	gameDrawTarget rl.RenderTexture2D
	gameSourceRect rl.Rectangle
	gameDestRect   rl.Rectangle

	score               *Score
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

	gs.scoreDrawTarget = rl.LoadRenderTexture(ScreenLogicalWidth, ScoreHeight)
	gs.scoreSourceRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), -float32(ScoreHeight))
	gs.scoreDestRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), float32(ScoreHeight))

	gs.gameDrawTarget = rl.LoadRenderTexture(ScreenLogicalWidth, ScreenLogicalHeight-ScoreHeight)
	gs.gameSourceRect = rl.NewRectangle(0, 0, float32(ScreenLogicalWidth), -float32(ScreenLogicalHeight-ScoreHeight))
	gs.gameDestRect = rl.NewRectangle(0, float32(ScoreHeight), float32(ScreenLogicalWidth), float32(ScreenLogicalHeight-ScoreHeight))

	gs.field = NewField(&gs)
	gs.moveGrid = NewMoveGrid(&gs)
	gs.digger = NewDigger(&gs)
	gs.fire = NewFire(&gs)
	gs.emeraldsPool = NewEmeraldsPool(&gs)
	gs.bagsPool = NewBagsPool(&gs)
	gs.monstersPool = NewMonstersPool(&gs)
	gs.debugGrid = NewDebugGrid(&gs)
	gs.score = NewScore(&gs)

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
	gs.score.Update(tick)
}

func (gs *GameScene) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(gs.gameDrawTarget)
	{
		gs.field.Render()
		gs.emeraldsPool.Render()
		gs.digger.Render()
		gs.fire.Render()
		gs.monstersPool.Render()
		gs.bagsPool.Render()
		gs.debugGrid.Render()
		gs.moveGrid.Render()
	}
	rl.EndTextureMode()

	rl.BeginTextureMode(gs.scoreDrawTarget)
	{
		gs.score.Render()
	}
	rl.EndTextureMode()

	rl.BeginTextureMode(drawTarget)
	{
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(gs.scoreDrawTarget.Texture, gs.scoreSourceRect, gs.scoreDestRect, ZERO_VECTOR2, 0, rl.White)
		rl.DrawTexturePro(gs.gameDrawTarget.Texture, gs.gameSourceRect, gs.gameDestRect, ZERO_VECTOR2, 0, rl.White)
	}
	rl.EndTextureMode()
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}
