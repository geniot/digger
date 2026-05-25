package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	COLORS = []rl.Color{
		rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple, rl.Orange, rl.Pink, rl.Brown, rl.Gray,
		rl.LightGray, rl.DarkGray, rl.DarkGreen, rl.DarkBlue, rl.DarkPurple, rl.DarkBrown,
	}
)

type GameScene struct {
	a      *Application
	target rl.RenderTexture2D
	//cube      *Cube
	isStarted bool
}

func NewGameScene(a *Application) *GameScene {
	gameScene := GameScene{}
	gameScene.a = a
	//gameScene.cube = NewCube(3, split(CubeCorrect), a)
	gameScene.isStarted = false
	//gameScene.Reset()
	gameScene.target = rl.LoadRenderTexture(100, 100)
	rl.BeginTextureMode(gameScene.target)
	rl.ClearBackground(rl.Yellow)
	rl.EndTextureMode()

	return &gameScene
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(camera *rl.Camera) {

	rl.BeginTextureMode(gs.target)
	rl.DrawPixel(int32(rand.Intn(100)), int32(rand.Intn(100)), COLORS[rand.Intn(len(COLORS))])
	//rl.DrawPixel(0, 0, rl.Black)
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(gs.target.Texture, 0, 0, rl.White)
	rl.EndDrawing()

	//image := rl.LoadImageFromTexture(gs.target.Texture)
	////colors := rl.LoadImageColors(image)
	//color1 := rl.GetImageColor(*image, 0, 0)
	//color2 := rl.GetImageColor(*image, 0, 1)
	////color2 := colors[1]
	//println(color1.RGBA())
	//println(color2.RGBA())
}
