package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Digger struct {
	scene            *GameScene
	posX             int32
	posY             int32
	width            int32
	height           int32
	innerOffsetX     int32
	innerOffsetY     int32
	direction        Direction
	shouldMove       bool
	spritePointer    int
	spritePointerInc int
	leftSprites      []*TextureImage
	rightSprites     []*TextureImage
	upSprites        []*TextureImage
	downSprites      []*TextureImage
}

func NewDigger(scene *GameScene) *Digger {
	digger := &Digger{}
	digger.scene = scene

	digger.leftSprites = make([]*TextureImage, 3)
	digger.leftSprites[0] = NewTextureImage("cldig1.png", 0, false, false)
	digger.leftSprites[1] = NewTextureImage("cldig2.png", 0, false, false)
	digger.leftSprites[2] = NewTextureImage("cldig3.png", 0, false, false)

	digger.rightSprites = make([]*TextureImage, 3)
	digger.rightSprites[0] = NewTextureImage("cldig1.png", 0, true, false)
	digger.rightSprites[1] = NewTextureImage("cldig2.png", 0, true, false)
	digger.rightSprites[2] = NewTextureImage("cldig3.png", 0, true, false)

	digger.upSprites = make([]*TextureImage, 3)
	digger.upSprites[0] = NewTextureImage("cldig1.png", 90, false, false)
	digger.upSprites[1] = NewTextureImage("cldig2.png", 90, false, false)
	digger.upSprites[2] = NewTextureImage("cldig3.png", 90, false, false)

	digger.downSprites = make([]*TextureImage, 3)
	digger.downSprites[0] = NewTextureImage("cldig1.png", 90, false, true)
	digger.downSprites[1] = NewTextureImage("cldig2.png", 90, false, true)
	digger.downSprites[2] = NewTextureImage("cldig3.png", 90, false, true)

	//same for all levels
	cellX := int32(7)
	cellY := int32(9)
	digger.innerOffsetX = int32(0)
	digger.innerOffsetY = int32(1)
	digger.spritePointer = 0
	digger.spritePointerInc = 1
	digger.width = 16
	digger.height = 16
	digger.posX = FIELD_OFFSET_X + cellX*CELL_WIDTH - digger.innerOffsetX
	digger.posY = FIELD_OFFSET_Y + cellY*CELL_HEIGHT - digger.innerOffsetY
	digger.direction = RIGHT
	digger.shouldMove = false
	return digger
}

func (digger *Digger) Update(tick int64) {
	if tick%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer, digger.spritePointerInc = GetNextSpritePointerAndInc(digger.spritePointer, digger.spritePointerInc, len(digger.leftSprites))
	}
	if tick%DIGGER_SPEED == 0 && digger.shouldMove && digger.canMove() {
		x := If(digger.direction == RIGHT, 1, If(digger.direction == LEFT, -1, 0))
		y := If(digger.direction == DOWN, 1, If(digger.direction == UP, -1, 0))
		digger.posX += int32(x)
		digger.posY += int32(y)
	}
}

func (digger *Digger) canMove() bool {
	return digger.scene.field.isWithinBounds(digger.direction, digger.getCollisionRec())
}

func (digger *Digger) getSprites() []*TextureImage {
	switch digger.direction {
	case RIGHT:
		return digger.rightSprites
	case LEFT:
		return digger.leftSprites
	case UP:
		return digger.upSprites
	case DOWN:
		return digger.downSprites
	default:
		return digger.rightSprites
	}
}

func (digger *Digger) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(digger.posX + (CELL_WIDTH-digger.width)/2),
		Y:      float32(digger.posY + (CELL_WIDTH-digger.height)/2),
		Width:  float32(digger.width),
		Height: float32(digger.height),
	}
}

func (digger *Digger) Render(drawTarget rl.RenderTexture2D) {
	sprites := digger.getSprites()
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexture(
		sprites[digger.spritePointer].texture,
		digger.posX, digger.posY,
		rl.White)
	rl.DrawRectangleLinesEx(digger.getCollisionRec(), 1.0, TransparentYellow)
	rl.EndTextureMode()
}
