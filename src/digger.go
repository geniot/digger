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
	sprites          []*TextureImage
}

func NewDigger(app *Application) *Digger {
	digger := &Digger{}
	digger.app = app

	digger.sprites = make([]*TextureImage, 3)
	digger.sprites[0] = NewTextureImage("cldig1.png")
	digger.sprites[1] = NewTextureImage("cldig2.png")
	digger.sprites[2] = NewTextureImage("cldig3.png")

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
		digger.sprites[digger.spritePointer].texture,
		sourceRec,
		destRec,
		rl.Vector2{},
		0,
		rl.White)
	rl.EndTextureMode()
}
