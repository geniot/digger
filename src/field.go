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
	fld.target = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT-20)
	rl.BeginTextureMode(fld.target)
	rl.ClearBackground(rl.Yellow)
	rl.EndTextureMode()
	return fld
}

func (c *Field) Update(drawTarget rl.RenderTexture2D) {
	randomColor := COLORS[rand.Intn(len(COLORS))]

	rl.BeginTextureMode(c.target)
	rl.DrawPixel(int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), randomColor)
	rl.DrawRectangleLines(1, 0, SCREEN_LOGICAL_WIDTH-1, SCREEN_LOGICAL_HEIGHT-21, rl.Red)
	//rl.DrawRectangle(int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), int32(rand.Intn(SCREEN_LOGICAL_WIDTH)), int32(rand.Intn(SCREEN_LOGICAL_HEIGHT)), randomColor)
	rl.EndTextureMode()

	rl.BeginTextureMode(drawTarget)
	rl.DrawTexture(c.target.Texture, 0, 20, rl.White)
	rl.EndTextureMode()

}
