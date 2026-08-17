package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Field struct {
	scene            *GameScene
	texture          rl.RenderTexture2D
	image            *rl.Image
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

	fld.textureSourceRec = rl.NewRectangle(0, 0, FieldWidth, -FieldHeight) //see https://github.com/raysan5/raylib/issues/3803
	fld.imageSourceRec = rl.NewRectangle(0, 0, FieldWidth, FieldHeight)
	fld.destRec = rl.NewRectangle(0, 0, FieldWidth, FieldHeight)

	bg := NewTextureImage("graphics/field/cback1.png", 0, false, false, false)

	fld.upBlob = NewTextureImage("graphics/field/cublob.png", 0, false, false, false)
	fld.downBlob = NewTextureImage("graphics/field/cdblob.png", 0, false, false, false)
	fld.leftBlob = NewTextureImage("graphics/field/clblob.png", 0, false, false, false)
	fld.rightBlob = NewTextureImage("graphics/field/crblob.png", 0, false, false, false)

	fld.texture = rl.LoadRenderTexture(FieldWidth, FieldHeight)
	fld.image = rl.GenImageColor(FieldWidth, FieldHeight, rl.Black)

	rl.BeginTextureMode(fld.texture)
	rl.ClearBackground(rl.Black)
	for y := int32(0); y < FieldHeight; y += int32(bg.height) {
		for x := int32(0); x < FieldWidth; x += int32(bg.width) {
			fld.draw(bg, float32(x), float32(y))
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

	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			c := getLevelChar(x, y, lp)
			if c == 'S' || c == 'V' || c == 'H' {
				xp := x*20 + 12
				yp := y*18 + 18
				if c == 'V' || c == 'S' {
					for decr := int32(-15); decr <= -3; decr += 3 {
						fld.draw(fld.downBlob, float32(xp+dX), float32(yp+decr+dY))
					}
					fld.draw(fld.upBlob, float32(xp+uX), float32(yp+3+uY))
				}
				if c == 'H' || c == 'S' {
					for decr := int32(-16); decr <= -4; decr += 4 {
						fld.draw(fld.rightBlob, float32(xp+decr+rX), float32(yp+rY))
					}
					fld.draw(fld.leftBlob, float32(xp+4+lX), float32(yp+lY))
				}
				if x < 14 && (getLevelChar(x+1, y, lp) == 'H' || getLevelChar(x+1, y, lp) == 'S') {
					fld.draw(fld.rightBlob, float32(xp+rX), float32(yp+rY))
				}
				if y < 9 && (getLevelChar(x, y+1, lp) == 'V' || getLevelChar(x, y+1, lp) == 'H') {
					fld.draw(fld.downBlob, float32(xp+dX), float32(yp+dY))
				}
			}
		}
	}
	return fld
}

func (field *Field) drawExt(textureImage *TextureImage, x, y float32, splitX, splitY, splitWidth, splitHeight bool) {
	field.draw(textureImage,
		x, y,
		rl.Rectangle{
			X:      If(splitX, textureImage.width/2, 0),
			Y:      If(splitY, textureImage.height/2, 0),
			Width:  If(splitWidth, textureImage.width/2, textureImage.width),
			Height: If(splitHeight, textureImage.height/2, textureImage.height)},

		rl.Rectangle{X: x, Y: y,
			Width:  If(splitWidth, textureImage.width/2, textureImage.width),
			Height: If(splitHeight, textureImage.height/2, textureImage.height)},
	)
}

func (field *Field) draw(textureImage *TextureImage, x float32, y float32, rects ...rl.Rectangle) {
	rl.BeginTextureMode(field.texture)
	sourceRect := rl.NewRectangle(0, 0, textureImage.width, textureImage.height)
	destRect := rl.NewRectangle(x, y, textureImage.width, textureImage.height)
	if len(rects) >= 1 {
		sourceRect = rects[0]
	}
	if len(rects) >= 2 {
		destRect = rects[1]
	}
	rl.DrawTexturePro(textureImage.texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	rl.ImageDraw(field.image, textureImage.image, sourceRect, destRect, rl.White)
	rl.EndTextureMode()
}

func (field *Field) Update(_ int64) {
}

func (field *Field) Render(drawTarget rl.RenderTexture2D) {
	//field.Debug()
	rl.BeginTextureMode(drawTarget)
	//rl.DrawTextureRec(rl.LoadTextureFromImage(field.image), field.imageSourceRec, ZERO_VECTOR2, rl.White)
	rl.DrawTexturePro(field.texture.Texture, field.textureSourceRec, field.destRec, ZERO_VECTOR2, 0, rl.White)
	rl.EndTextureMode()
}

func (field *Field) Debug() {
	clone1 := rl.ImageCopy(field.image)
	rl.ImageFlipVertical(clone1)
	colors1 := rl.LoadImageColors(clone1)
	defer rl.UnloadImageColors(colors1)

	clone2 := rl.LoadImageFromTexture(field.texture.Texture)
	colors2 := rl.LoadImageColors(clone2)
	defer rl.UnloadImageColors(colors2)

	//println(len(colors1))
	if len(colors1) != len(colors2) {
		panic("colors are different")
	}
	for i := 0; i < len(colors1); i++ {
		if colors1[i].R != colors2[i].R || colors1[i].G != colors2[i].G || colors1[i].B != colors2[i].B || colors1[i].A != colors2[i].A {
			println(colors1[i].R, " ", colors1[i].G, " ", colors1[i].B, " ", colors1[i].A, " ")
			println(colors2[i].R, " ", colors2[i].G, " ", colors2[i].B, " ", colors2[i].A, " ")
			panic("colors are different")
		}
	}
}
