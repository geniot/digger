package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
	field   [150]int32
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

	for x := 0; x < 15; x++ {
		for y := 0; y < 10; y++ {
			field[y*15+x] = -1
			c := getLevelChar(x, y, levplan())
			if c == 'S' || c == 'V' {
				field[y*15+x] &= 0xd03f
			}
			if c == 'S' || c == 'H' {
				field[y*15+x] &= 0xdfe0
			}
		}
	}

	bgBytes := orPanicRes(resList.ReadFile("res/cback1.png"))
	bgTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", bgBytes, int32(len(bgBytes))))
	upBlobBytes := orPanicRes(resList.ReadFile("res/cublob.png"))
	upBlobTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", upBlobBytes, int32(len(upBlobBytes))))
	downBlobBytes := orPanicRes(resList.ReadFile("res/cdblob.png"))
	downBlobTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", downBlobBytes, int32(len(downBlobBytes))))
	leftBlobBytes := orPanicRes(resList.ReadFile("res/clblob.png"))
	leftBlobTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", leftBlobBytes, int32(len(leftBlobBytes))))
	rightBlobBytes := orPanicRes(resList.ReadFile("res/crblob.png"))
	rightBlobTexture := rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", rightBlobBytes, int32(len(rightBlobBytes))))

	fld.target = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)

	rl.BeginTextureMode(fld.target)
	rl.ClearBackground(rl.Black)
	for y := 14; y < 200; y += 4 {
		for x := 0; x < 320; x += 20 {
			sourceRect := rl.NewRectangle(0, 0, 20, float32(If(y+4 > 200, 2, 4)))
			rl.DrawTextureRec(bgTexture, sourceRect, rl.Vector2{X: float32(x), Y: float32(y)}, rl.White)
		}
	}
	dX := int32(-2)
	dY := int32(15)
	uX := int32(-2)
	uY := int32(-6)
	rX := int32(16)
	rY := int32(-1)
	lX := int32(-8)
	lY := int32(-1)

	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			if field[y*15+x]&0x2000 == 0 {
				xp := x*20 + 12
				yp := y*18 + 18
				if field[y*15+x]&0xfc0 != 0xfc0 {
					field[y*15+x] &= 0xd03f
					rl.DrawTexture(downBlobTexture, xp+dX, yp-15+dY, rl.White)
					rl.DrawTexture(downBlobTexture, xp+dX, yp-12+dY, rl.White)
					rl.DrawTexture(downBlobTexture, xp+dX, yp-9+dY, rl.White)
					rl.DrawTexture(downBlobTexture, xp+dX, yp-6+dY, rl.White)
					rl.DrawTexture(downBlobTexture, xp+dX, yp-3+dY, rl.White)
					rl.DrawTexture(upBlobTexture, xp+uX, yp+3+uY, rl.White)
				}
				if field[y*15+x]&0x1f != 0x1f {
					field[y*15+x] &= 0xdfe0
					rl.DrawTexture(rightBlobTexture, xp-16+rX, yp+rY, rl.White)
					rl.DrawTexture(rightBlobTexture, xp-12+rX, yp+rY, rl.White)
					rl.DrawTexture(rightBlobTexture, xp-8+rX, yp+rY, rl.White)
					rl.DrawTexture(rightBlobTexture, xp-4+rX, yp+rY, rl.White)
					rl.DrawTexture(leftBlobTexture, xp+4+lX, yp+lY, rl.White)
				}
				if x < 14 && field[y*15+x+1]&0xfdf != 0xfdf {
					rl.DrawTexture(rightBlobTexture, xp+rX, yp+rY, rl.White)
				}
				if y < 9 && field[(y+1)*15+x]&0xfdf != 0xfdf {
					rl.DrawTexture(downBlobTexture, xp+dX, yp+dY, rl.White)
				}
			}
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
