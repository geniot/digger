package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Field struct {
	scene            *GameScene
	texture          rl.RenderTexture2D
	state            [FieldWidth][FieldHeight]bool
	textureSourceRec rl.Rectangle
	imageSourceRec   rl.Rectangle
	destRec          rl.Rectangle
	upBlob           *MaskedTexture
	downBlob         *MaskedTexture
	leftBlob         *MaskedTexture
	rightBlob        *MaskedTexture
}

func NewField(scene *GameScene) *Field {
	fld := &Field{}
	fld.scene = scene

	fld.textureSourceRec = rl.NewRectangle(0, 0, float32(FieldWidth), -float32(FieldHeight)) //see https://github.com/raysan5/raylib/issues/3803
	fld.imageSourceRec = rl.NewRectangle(0, 0, float32(FieldWidth), float32(FieldHeight))
	fld.destRec = rl.NewRectangle(0, 0, float32(FieldWidth), float32(FieldHeight))

	bg := NewMaskedTexture("graphics/field/cback1.png", 0, false, false, false)

	fld.upBlob = NewMaskedTexture("graphics/field/cublob.png", 0, false, false, false)
	fld.downBlob = NewMaskedTexture("graphics/field/cdblob.png", 0, false, false, false)
	fld.leftBlob = NewMaskedTexture("graphics/field/clblob.png", 0, false, false, false)
	fld.rightBlob = NewMaskedTexture("graphics/field/crblob.png", 0, false, false, false)

	fld.texture = rl.LoadRenderTexture(FieldWidth, FieldHeight)
	fld.state = [FieldWidth][FieldHeight]bool{}
	for x := range FieldWidth {
		for y := range FieldHeight {
			fld.state[x][y] = true
		}
	}

	rl.BeginTextureMode(fld.texture)
	rl.ClearBackground(rl.Black)
	for x := int32(0); x < FieldWidth; x += bg.width {
		for y := int32(0); y < FieldHeight; y += bg.height {
			fld.draw(bg, x, y)
		}
	}
	rl.EndTextureMode()
	//little offsets as copied from the original code
	dX := int32(-2)
	dY := int32(1)
	uX := int32(-2)
	uY := int32(-20)
	rX := int32(16)
	rY := int32(-15)
	lX := int32(-8)
	lY := int32(-15)

	lp := LevelPlan(scene.level)

	for x := range int32(15) {
		for y := range int32(10) {
			c := getLevelChar(x, y, lp)
			if c == 'S' || c == 'V' || c == 'H' {
				xp := x*20 + 12
				yp := y*18 + 18
				if c == 'V' || c == 'S' {
					for decr := int32(-15); decr <= -3; decr += 3 {
						fld.draw(fld.downBlob, xp+dX, yp+decr+dY)
					}
					fld.draw(fld.upBlob, xp+uX, yp+3+uY)
				}
				if c == 'H' || c == 'S' {
					for decr := int32(-16); decr <= -4; decr += 4 {
						fld.draw(fld.rightBlob, xp+decr+rX, yp+rY)
					}
					fld.draw(fld.leftBlob, xp+4+lX, yp+lY)
				}
				if x < 14 && (getLevelChar(x+1, y, lp) == 'H' || getLevelChar(x+1, y, lp) == 'S') {
					fld.draw(fld.rightBlob, xp+rX, yp+rY)
				}
				if y < 9 && (getLevelChar(x, y+1, lp) == 'V' || getLevelChar(x, y+1, lp) == 'H') {
					fld.draw(fld.downBlob, xp+dX, yp+dY)
				}
			}
		}
	}
	return fld
}

func (field *Field) drawExt(maskedTexture *MaskedTexture, x, y int32, splitX, splitY, splitWidth, splitHeight bool) {
	width := float32(If(splitWidth, maskedTexture.width/2, maskedTexture.width))
	height := float32(If(splitHeight, maskedTexture.height/2, maskedTexture.height))
	field.draw(maskedTexture,
		x, y,
		rl.Rectangle{
			X:      float32(If(splitX, maskedTexture.width/2, 0)),
			Y:      float32(If(splitY, maskedTexture.height/2, 0)),
			Width:  width,
			Height: height},
		rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  width,
			Height: height},
	)
}

func (field *Field) draw(maskedTexture *MaskedTexture, x int32, y int32, rects ...rl.Rectangle) {
	rl.BeginTextureMode(field.texture)
	sourceRect := rl.NewRectangle(0, 0, float32(maskedTexture.width), float32(maskedTexture.height))
	destRect := rl.NewRectangle(float32(x), float32(y), float32(maskedTexture.width), float32(maskedTexture.height))
	if len(rects) >= 1 {
		sourceRect = rects[0]
	}
	if len(rects) >= 2 {
		destRect = rects[1]
	}
	rl.DrawTexturePro(maskedTexture.texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	rl.EndTextureMode()
	// updating the masked state
	for xx := int32(destRect.X); xx < int32(destRect.X)+int32(destRect.Width); xx++ {
		for yy := int32(destRect.Y); yy < int32(destRect.Y)+int32(destRect.Height); yy++ {
			xxx := xx - int32(destRect.X) + int32(sourceRect.X)
			yyy := yy - int32(destRect.Y) + int32(sourceRect.Y)
			if maskedTexture.mask[xxx][yyy] {
				field.state[xx][yy] = false
			}
		}
	}
}

func (field *Field) Update(_ int64) {
}

func (field *Field) Render() {
	rl.DrawTexturePro(field.texture.Texture, field.textureSourceRec, field.destRec, ZERO_VECTOR2, 0, rl.White)
	//field.Debug()
}

func (field *Field) IsColliding(rec rl.Rectangle) bool {
	recX, recY, recW, recH := int32(rec.X), int32(rec.Y), int32(rec.Width), int32(rec.Height)
	for x := recX; x < recX+recW; x++ {
		for y := recY; y < recY+recH; y++ {
			if field.state[x][y] {
				return true
			}
		}
	}
	return false
}

func (field *Field) Debug() {
	for x := range FieldWidth {
		for y := range FieldHeight {
			if field.state[x][y] {
				rl.DrawPixel(x, y, TransparentBlue)
			}
		}
	}
}
