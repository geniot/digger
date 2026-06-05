package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
)

const (
	FIELD_WIDTH  = 320
	FIELD_HEIGHT = 186
)

type Field struct {
	app       *Application
	texture   rl.RenderTexture2D
	image     *rl.Image
	sourceRec rl.Rectangle
	destRec   rl.Rectangle
}

func NewField(app *Application) *Field {
	fld := &Field{}
	fld.app = app

	fld.sourceRec = rl.NewRectangle(0, 0, FIELD_WIDTH, -FIELD_HEIGHT) //see https://github.com/raysan5/raylib/issues/3803
	fld.destRec = rl.NewRectangle(0, 0, FIELD_WIDTH, FIELD_HEIGHT)

	bgBytes := orPanicRes(resList.ReadFile("res/cback1.png"))
	bgImage := rl.LoadImageFromMemory(".png", bgBytes, int32(len(bgBytes)))
	bgTexture := rl.LoadTextureFromImage(bgImage)

	upBlobBytes := orPanicRes(resList.ReadFile("res/cublob.png"))
	upBlobImage := rl.LoadImageFromMemory(".png", upBlobBytes, int32(len(upBlobBytes)))
	upBlobTexture := rl.LoadTextureFromImage(upBlobImage)

	downBlobBytes := orPanicRes(resList.ReadFile("res/cdblob.png"))
	downBlobImage := rl.LoadImageFromMemory(".png", downBlobBytes, int32(len(downBlobBytes)))
	downBlobTexture := rl.LoadTextureFromImage(downBlobImage)

	leftBlobBytes := orPanicRes(resList.ReadFile("res/clblob.png"))
	leftBlobImage := rl.LoadImageFromMemory(".png", leftBlobBytes, int32(len(leftBlobBytes)))
	leftBlobTexture := rl.LoadTextureFromImage(leftBlobImage)

	rightBlobBytes := orPanicRes(resList.ReadFile("res/crblob.png"))
	rightBlobImage := rl.LoadImageFromMemory(".png", rightBlobBytes, int32(len(rightBlobBytes)))
	rightBlobTexture := rl.LoadTextureFromImage(rightBlobImage)

	fld.texture = rl.LoadRenderTexture(FIELD_WIDTH, FIELD_HEIGHT)
	fld.image = rl.GenImageColor(FIELD_WIDTH, FIELD_HEIGHT, rl.Black)

	rl.BeginTextureMode(fld.texture)
	rl.ClearBackground(rl.Black)
	for y := int32(0); y < FIELD_HEIGHT; y += bgTexture.Height {
		for x := int32(0); x < FIELD_WIDTH; x += bgTexture.Width {
			fld.draw(float32(x), float32(y), float32(bgTexture.Width),
				float32(If(y+bgTexture.Height > FIELD_HEIGHT, bgTexture.Height/2, bgTexture.Height)),
				&bgTexture, bgImage)
		}
	}
	//little offsets as copied from the original code
	dX := int32(-2)
	dY := int32(1)
	uX := int32(-2)
	uY := int32(-20)
	rX := int32(16)
	rY := int32(-15)
	lX := int32(-8)
	lY := int32(-15)

	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			c := getLevelChar(x, y, levplan())
			if c == 'S' || c == 'V' || c == 'H' {
				xp := x*20 + 12
				yp := y*18 + 18
				if c == 'V' || c == 'S' {
					for decr := int32(-15); decr <= -3; decr += 3 {
						fld.draw(float32(xp+dX), float32(yp+decr+dY), float32(downBlobImage.Width), float32(downBlobImage.Height), &downBlobTexture, downBlobImage)
					}
					fld.draw(float32(xp+uX), float32(yp+3+uY), float32(upBlobImage.Width), float32(upBlobImage.Height), &upBlobTexture, upBlobImage)
				}
				if c == 'H' || c == 'S' {
					for decr := int32(-16); decr <= -4; decr += 4 {
						fld.draw(float32(xp+decr+rX), float32(yp+rY), float32(rightBlobImage.Width), float32(rightBlobImage.Height), &rightBlobTexture, rightBlobImage)
					}
					fld.draw(float32(xp+4+lX), float32(yp+lY), float32(leftBlobImage.Width), float32(leftBlobImage.Height), &leftBlobTexture, leftBlobImage)
				}
				if x < 14 && (getLevelChar(x+1, y, levplan()) == 'H' || getLevelChar(x+1, y, levplan()) == 'S') {
					fld.draw(float32(xp+rX), float32(yp+rY), float32(rightBlobImage.Width), float32(rightBlobImage.Height), &rightBlobTexture, rightBlobImage)
				}
				if y < 9 && (getLevelChar(x, y+1, levplan()) == 'V' || getLevelChar(x, y+1, levplan()) == 'H') {
					fld.draw(float32(xp+dX), float32(yp+dY), float32(downBlobImage.Width), float32(downBlobImage.Height), &downBlobTexture, downBlobImage)
				}
			}
		}
	}
	rl.EndTextureMode()
	return fld
}

func (c *Field) draw(x, y, width, height float32, texture *rl.Texture2D, image *rl.Image) {
	sourceRect := rl.NewRectangle(0, 0, width, height)
	destRect := rl.NewRectangle(x, y, width, height)
	rl.DrawTexturePro(*texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	rl.ImageDraw(c.image, image, sourceRect, destRect, rl.White)
}

func (c *Field) Update(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	//c.Debug()
	//rl.DrawTextureRec(rl.LoadTextureFromImage(clone), c.sourceRect, c.zeroVector, rl.White)
	rl.DrawTexturePro(c.texture.Texture, c.sourceRec, c.destRec, ZERO_VECTOR2, 0, rl.White)
	rl.EndTextureMode()
}

func (c *Field) Debug() {
	clone1 := rl.ImageCopy(c.image)
	rl.ImageFlipVertical(clone1)
	colors1 := rl.LoadImageColors(clone1)

	clone2 := rl.LoadImageFromTexture(c.texture.Texture)
	colors2 := rl.LoadImageColors(clone2)

	if len(colors1) != len(colors2) {
		panic("colors are different")
	}
	for i := 0; i < len(colors1); i++ {
		if colors1[i].R != colors2[i].R || colors1[i].G != colors2[i].G || colors1[i].B != colors2[i].B || colors1[i].A != colors2[i].A {
			panic("colors are different")
		}
	}
}
