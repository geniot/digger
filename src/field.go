package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FIELD_WIDTH  = 320
	FIELD_HEIGHT = 186
)

type Field struct {
	scene            *GameScene
	texture          rl.RenderTexture2D
	image            *rl.Image
	textureSourceRec rl.Rectangle
	imageSourceRec   rl.Rectangle
	destRec          rl.Rectangle
}

func NewField(scene *GameScene) *Field {
	fld := &Field{}
	fld.scene = scene

	fld.textureSourceRec = rl.NewRectangle(0, 0, FIELD_WIDTH, -FIELD_HEIGHT) //see https://github.com/raysan5/raylib/issues/3803
	fld.imageSourceRec = rl.NewRectangle(0, 0, FIELD_WIDTH, FIELD_HEIGHT)
	fld.destRec = rl.NewRectangle(0, 0, FIELD_WIDTH, FIELD_HEIGHT)

	bg := NewTextureImage("cback1.png")
	upBlob := NewTextureImage("cublob.png")
	downBlob := NewTextureImage("cdblob.png")
	leftBlob := NewTextureImage("clblob.png")
	rightBlob := NewTextureImage("crblob.png")

	fld.texture = rl.LoadRenderTexture(FIELD_WIDTH, FIELD_HEIGHT)
	fld.image = rl.GenImageColor(FIELD_WIDTH, FIELD_HEIGHT, rl.Black)

	rl.BeginTextureMode(fld.texture)
	rl.ClearBackground(rl.Black)
	for y := int32(0); y < FIELD_HEIGHT; y += int32(bg.height) {
		for x := int32(0); x < FIELD_WIDTH; x += int32(bg.width) {
			fld.draw(float32(x), float32(y), bg)
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
						fld.draw(float32(xp+dX), float32(yp+decr+dY), downBlob)
					}
					fld.draw(float32(xp+uX), float32(yp+3+uY), upBlob)
				}
				if c == 'H' || c == 'S' {
					for decr := int32(-16); decr <= -4; decr += 4 {
						fld.draw(float32(xp+decr+rX), float32(yp+rY), rightBlob)
					}
					fld.draw(float32(xp+4+lX), float32(yp+lY), leftBlob)
				}
				if x < 14 && (getLevelChar(x+1, y, levplan()) == 'H' || getLevelChar(x+1, y, levplan()) == 'S') {
					fld.draw(float32(xp+rX), float32(yp+rY), rightBlob)
				}
				if y < 9 && (getLevelChar(x, y+1, levplan()) == 'V' || getLevelChar(x, y+1, levplan()) == 'H') {
					fld.draw(float32(xp+dX), float32(yp+dY), downBlob)
				}
			}
		}
	}
	rl.EndTextureMode()
	return fld
}

func (field *Field) draw(x float32, y float32, textureImage *TextureImage) {
	sourceRect := rl.NewRectangle(0, 0, textureImage.width, textureImage.height)
	destRect := rl.NewRectangle(x, y, textureImage.width, textureImage.height)
	rl.DrawTexturePro(textureImage.texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	rl.ImageDraw(field.image, textureImage.image, sourceRect, destRect, rl.White)
}

func (field *Field) Update(drawTarget rl.RenderTexture2D, _ int64) {
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

func (field *Field) isWithinBounds(dir Direction, offsetX int32, offsetY int32) bool {
	//screen bounds
	switch dir {
	case RIGHT:
		return offsetX < CELL_WIDTH*(CELLS_HORIZONTAL-1)
	case LEFT:
		return offsetX > 0
	case UP:
		return offsetY > FIELD_OFFSET_Y
	case DOWN:
		return offsetY < FIELD_OFFSET_Y+CELL_HEIGHT*(CELLS_VERTICAL-1)
	default:
		return true
	}
}
