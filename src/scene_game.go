package main

import (
	"image"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	COLORS = []rl.Color{
		rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple, rl.Orange, rl.Pink, rl.Brown, rl.Gray,
		rl.LightGray, rl.DarkGray, rl.DarkGreen, rl.DarkBlue, rl.DarkPurple, rl.DarkBrown,
	}
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
	rgba      *image.RGBA
	target    rl.RenderTexture2D
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	//gameScene.cube = NewCube(3, split(CubeCorrect), a)
	gameScene.isStarted = false
	//gameScene.Reset()
	gameScene.target = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)

	rl.BeginTextureMode(gameScene.target)
	rl.ClearBackground(rl.Yellow)
	rl.EndTextureMode()

	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(camera *rl.Camera) {
	randomColor := COLORS[rand.Intn(len(COLORS))]

	rl.BeginTextureMode(gs.target)
	rl.DrawPixel(int32(rand.Intn(100)), int32(rand.Intn(100)), randomColor)
	rl.DrawRectangle(int32(rand.Intn(100)), int32(rand.Intn(100)), int32(rand.Intn(100)), int32(rand.Intn(100)), randomColor)
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(gs.target.Texture, 0, 0, rl.White)
	rl.EndDrawing()
}
