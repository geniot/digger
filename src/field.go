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

type Field struct {
	app    *Application
	target rl.RenderTexture2D
}

func NewField(app *Application) *Field {
	fld := &Field{}
	fld.app = app
	fld.target = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	rl.BeginTextureMode(fld.target)
	rl.ClearBackground(rl.Yellow)
	rl.EndTextureMode()
	return fld
}

func (c *Field) Update() {
	randomColor := COLORS[rand.Intn(len(COLORS))]

	rl.BeginTextureMode(c.target)
	rl.DrawPixel(int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), randomColor)
	//rl.DrawRectangle(int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), randomColor)
	rl.EndTextureMode()

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	sourceRect := rl.NewRectangle(0, float32(-SCREEN_LOGICAL_HEIGHT), float32(SCREEN_LOGICAL_WIDTH), float32(-SCREEN_LOGICAL_HEIGHT))
	ratioX := screenWidth / float32(SCREEN_LOGICAL_WIDTH)
	ratioY := screenHeight / float32(SCREEN_LOGICAL_HEIGHT)
	resizeRatio := If(ratioX < ratioY, ratioX, ratioY)
	destRect := rl.NewRectangle(
		(screenWidth-(SCREEN_LOGICAL_WIDTH*resizeRatio))*0.5,
		(screenHeight-(SCREEN_LOGICAL_HEIGHT*resizeRatio))*0.5,
		SCREEN_LOGICAL_WIDTH*resizeRatio,
		SCREEN_LOGICAL_HEIGHT*resizeRatio,
	)

	rl.DrawTexturePro(c.target.Texture,
		sourceRect,
		destRect,
		rl.Vector2{}, 0, rl.White)
}
