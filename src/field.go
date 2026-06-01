package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
)

type Field struct {
	app        *Application
	target     rl.RenderTexture2D
	sourceRect rl.Rectangle
	zeroVector rl.Vector2
}

func NewField(app *Application) *Field {
	fld := &Field{}
	fld.app = app

	fld.sourceRect = rl.NewRectangle(0, 0, SCREEN_LOGICAL_WIDTH, -SCREEN_LOGICAL_HEIGHT) //see https://github.com/raysan5/raylib/issues/3803
	fld.zeroVector = rl.Vector2{X: 0, Y: 0}

	bgBytes := orPanicRes(resList.ReadFile("res/cback1.png"))
	bgTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", bgBytes, int32(len(bgBytes))))
	fld.target = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)

	rl.BeginTextureMode(fld.target)
	rl.ClearBackground(rl.Black)
	for y := 14; y < 200; y += 4 {
		for x := 0; x < 320; x += 20 {
			sourceRect := rl.NewRectangle(0, 0, 20, float32(If(y+4 > 200, 2, 4)))
			rl.DrawTextureRec(bgTexture, sourceRect, rl.Vector2{X: float32(x), Y: float32(y)}, rl.White)
		}
	}
	rl.EndTextureMode()
	return fld
}

func (c *Field) Update(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	rl.DrawTextureRec(c.target.Texture, c.sourceRect, c.zeroVector, rl.White)
	rl.EndTextureMode()

}
