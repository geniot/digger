package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SPRITE_UPDATE_RATE = 5
)

type Digger struct {
	app              *Application
	posX             int32
	posY             int32
	innerOffsetX     int32
	innerOffsetY     int32
	direction        Direction
	spritePointer    int
	spritePointerInc int
	sprites          []rl.Texture2D
}

func NewDigger(app *Application) *Digger {
	digger := &Digger{}
	digger.app = app

	digger.sprites = make([]rl.Texture2D, 3)

	dg1Bytes := orPanicRes(resList.ReadFile("res/cldig1.png"))
	dg1Image := rl.LoadImageFromMemory(".png", dg1Bytes, int32(len(dg1Bytes)))
	digger.sprites[0] = rl.LoadTextureFromImage(dg1Image)

	dg2Bytes := orPanicRes(resList.ReadFile("res/cldig2.png"))
	dg2Image := rl.LoadImageFromMemory(".png", dg2Bytes, int32(len(dg2Bytes)))
	digger.sprites[1] = rl.LoadTextureFromImage(dg2Image)

	dg3Bytes := orPanicRes(resList.ReadFile("res/cldig3.png"))
	dg3Image := rl.LoadImageFromMemory(".png", dg3Bytes, int32(len(dg3Bytes)))
	digger.sprites[2] = rl.LoadTextureFromImage(dg3Image)

	//same for all levels
	cellX := int32(7)
	cellY := int32(9)
	digger.innerOffsetX = int32(0)
	digger.innerOffsetY = int32(1)
	digger.spritePointer = 0
	digger.spritePointerInc = 1
	digger.posX = FIELD_OFFSET_X + cellX*CELL_WIDTH - digger.innerOffsetX
	digger.posY = FIELD_OFFSET_Y + cellY*CELL_HEIGHT - digger.innerOffsetY
	digger.direction = RIGHT
	return digger
}

func (digger *Digger) Update(drawTarget rl.RenderTexture2D, frame int64) {
	if frame%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer, digger.spritePointerInc = GetNextSpritePointerAndInc(digger.spritePointer, digger.spritePointerInc, len(digger.sprites))
	}
	sourceRec := rl.NewRectangle(0, 0, float32(IfInt(digger.direction == LEFT, CELL_WIDTH, -CELL_WIDTH)), float32(CELL_HEIGHT))
	destRec := rl.NewRectangle(float32(digger.posX), float32(digger.posY), float32(CELL_WIDTH), float32(CELL_HEIGHT))
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(
		digger.sprites[digger.spritePointer],
		sourceRec,
		destRec,
		rl.Vector2{},
		0,
		rl.White)
	rl.EndTextureMode()
}
