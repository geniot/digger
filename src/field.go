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
	upBlob           *TextureImage
	downBlob         *TextureImage
	leftBlob         *TextureImage
	rightBlob        *TextureImage
}

func NewField(scene *GameScene) *Field {
	fld := &Field{}
	fld.scene = scene

	fld.textureSourceRec = rl.NewRectangle(0, 0, float32(FieldWidth), -float32(FieldHeight)) //see https://github.com/raysan5/raylib/issues/3803
	fld.imageSourceRec = rl.NewRectangle(0, 0, float32(FieldWidth), float32(FieldHeight))
	fld.destRec = rl.NewRectangle(0, 0, float32(FieldWidth), float32(FieldHeight))

	bg := NewTextureImage("graphics/field/cback1.png", 0, false, false, false)

	fld.upBlob = NewTextureImage("graphics/field/cublob.png", 0, false, false, false)
	fld.downBlob = NewTextureImage("graphics/field/cdblob.png", 0, false, false, false)
	fld.leftBlob = NewTextureImage("graphics/field/clblob.png", 0, false, false, false)
	fld.rightBlob = NewTextureImage("graphics/field/crblob.png", 0, false, false, false)

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

func (field *Field) drawExt(textureImage *TextureImage, x, y int32, splitX, splitY, splitWidth, splitHeight bool) {
	field.draw(textureImage,
		x, y,
		rl.Rectangle{
			X:      float32(If(splitX, textureImage.width/2, 0)),
			Y:      float32(If(splitY, textureImage.height/2, 0)),
			Width:  float32(If(splitWidth, textureImage.width/2, textureImage.width)),
			Height: float32(If(splitHeight, textureImage.height/2, textureImage.height))},
		rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(If(splitWidth, textureImage.width/2, textureImage.width)),
			Height: float32(If(splitHeight, textureImage.height/2, textureImage.height))},
	)
}

func (field *Field) draw(textureImage *TextureImage, x int32, y int32, rects ...rl.Rectangle) {
	rl.BeginTextureMode(field.texture)
	sourceRect := rl.NewRectangle(0, 0, float32(textureImage.width), float32(textureImage.height))
	destRect := rl.NewRectangle(float32(x), float32(y), float32(textureImage.width), float32(textureImage.height))
	if len(rects) >= 1 {
		sourceRect = rects[0]
	}
	if len(rects) >= 2 {
		destRect = rects[1]
	}
	rl.DrawTexturePro(textureImage.texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	for xx := int32(0); xx < int32(len(textureImage.mask)); xx++ {
		for yy := int32(0); yy < int32(len(textureImage.mask[xx])); yy++ {
			if textureImage.mask[xx][yy] && x+xx < FieldWidth && y+yy < FieldHeight {
				field.state[x+xx][y+yy] = true
			}
		}
	}
	rl.EndTextureMode()
}

func (field *Field) Update(_ int64) {
}

func (field *Field) Render() {
	//field.Debug()
	//rl.DrawTextureRec(rl.LoadTextureFromImage(field.image), field.imageSourceRec, ZERO_VECTOR2, rl.White)
	rl.DrawTexturePro(field.texture.Texture, field.textureSourceRec, field.destRec, ZERO_VECTOR2, 0, rl.White)
}

func (field *Field) Debug() {
	image := rl.LoadImageFromTexture(field.texture.Texture)
	for x := range FieldWidth {
		for y := range FieldHeight {
			if IsPixelColored(x, y, image) && field.state[x][y] {
				println(x, " ", y, " ")
				panic("colors are different")
			}
		}
	}
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
