package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Digger struct {
	scene              *GameScene
	posX               int32
	posY               int32
	width              int32
	height             int32
	direction          Direction
	requestedDirection Direction
	shouldMove         bool
	spritePointer      int
	spritePointerInc   int
	leftSprites        []*TextureImage
	rightSprites       []*TextureImage
	upSprites          []*TextureImage
	downSprites        []*TextureImage
	speed              int64
}

func NewDigger(scene *GameScene) *Digger {
	dg := &Digger{}
	dg.scene = scene

	dg.leftSprites = dg.initSprites(0, false, false)
	dg.rightSprites = dg.initSprites(0, true, false)
	dg.upSprites = dg.initSprites(90, false, false)
	dg.downSprites = dg.initSprites(90, true, true)

	dg.spritePointer = 0
	dg.spritePointerInc = 1
	dg.width = 16
	dg.height = 16
	dg.posX, dg.posY = scene.moveGrid.getDiggerStartPos()
	dg.direction = RIGHT
	dg.shouldMove = false
	dg.speed = DIGGER_SPEED
	return dg
}

func (dg *Digger) initSprites(degrees int32, flipHorizontal bool, flipVertical bool) []*TextureImage {
	sprites := make([]*TextureImage, 3)
	sprites[0] = NewTextureImage("cldig1.png", degrees, flipHorizontal, flipVertical, false)
	sprites[1] = NewTextureImage("cldig2.png", degrees, flipHorizontal, flipVertical, false)
	sprites[2] = NewTextureImage("cldig3.png", degrees, flipHorizontal, flipVertical, false)
	return sprites
}

func (dg *Digger) Update(tick int64) {
	if tick%SPRITE_UPDATE_RATE == 0 {
		dg.spritePointer, dg.spritePointerInc = GetNextSpritePointerAndInc(dg.spritePointer, dg.spritePointerInc, len(dg.leftSprites))
	}
	if tick%dg.speed == 0 && dg.shouldMove {
		posX, posY, dir := dg.scene.moveGrid.move(dg.posX, dg.posY, dg.direction, dg.requestedDirection)
		if dg.posX != posX || dg.posY != posY || dg.direction != dir { //any change from previous state?
			dg.scene.emeraldsPool.handle(dg)
			dg.posX, dg.posY, dg.direction = posX, posY, dir
			if dir == RIGHT {
				blob := dg.scene.field.rightBlob
				dg.scene.field.drawExt(blob, float32(dg.posX-dg.posX%4+4), float32(dg.posY-CELL_HEIGHT/2+1), false, false, true, false)
				dg.scene.field.drawExt(blob, float32(dg.posX+8), float32(dg.posY-CELL_HEIGHT/2+1), true, false, true, false)
				dg.scene.field.drawExt(blob, float32(dg.posX+7), float32(dg.posY-CELL_HEIGHT/2+1), true, false, true, false)
			} else if dir == LEFT {
				blob := dg.scene.field.leftBlob
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2-2), float32(dg.posY-CELL_HEIGHT/2+1), false, false, true, false)
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2-1), float32(dg.posY-CELL_HEIGHT/2+1), false, false, true, false)
				dg.scene.field.drawExt(blob, float32(dg.posX-dg.posX%4-CELL_WIDTH/2+IfInt(dg.posX <= 20, 2, 6)), float32(dg.posY-CELL_HEIGHT/2+1), true, false, true, false)
			} else if dir == UP {
				blob := dg.scene.field.upBlob
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY-CELL_HEIGHT/2-dg.posY%3+4), false, true, false, true)
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY-CELL_HEIGHT/2-1), false, false, false, true)
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY-CELL_HEIGHT/2), false, false, false, true)
			} else if dir == DOWN {
				blob := dg.scene.field.downBlob
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY+CELL_HEIGHT/2-dg.posY%3-IfInt(dg.posY >= 173, 2, 5)), false, false, false, true)
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY+CELL_HEIGHT/2-1), false, true, false, true)
				dg.scene.field.drawExt(blob, float32(dg.posX-CELL_WIDTH/2), float32(dg.posY+CELL_HEIGHT/2-2), false, true, false, true)
			}

		}
	}
}

func (dg *Digger) getSprites() []*TextureImage {
	switch dg.direction {
	case RIGHT:
		return dg.rightSprites
	case LEFT:
		return dg.leftSprites
	case UP:
		return dg.upSprites
	case DOWN:
		return dg.downSprites
	default:
		return dg.rightSprites
	}
}

func (dg *Digger) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(dg.posX + (CELL_WIDTH-dg.width)/2 - CELL_WIDTH/2 - DIGGER_INNER_OFFSET_X + collisionOffsetsMap[dg.direction].X),
		Y:      float32(dg.posY + (CELL_WIDTH-dg.height)/2 - CELL_HEIGHT/2 - DIGGER_INNER_OFFSET_Y + collisionOffsetsMap[dg.direction].Y),
		Width:  float32(dg.width + collisionSizeMap[dg.direction].W),
		Height: float32(dg.height + collisionSizeMap[dg.direction].H),
	}
}

var (
	renderOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 0},
		RIGHT: {0, 0},
		UP:    {-1, 0},
		DOWN:  {1, 1},
	}
	collisionOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 2},
		RIGHT: {0, 2},
		UP:    {0, 0},
		DOWN:  {2, 1},
	}
	collisionSizeMap = map[Direction]WidthHeight{
		LEFT:  {0, -2},
		RIGHT: {0, -2},
		UP:    {-2, 0},
		DOWN:  {-2, 0},
	}
)

func (dg *Digger) Render(drawTarget rl.RenderTexture2D) {
	sprites := dg.getSprites()
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexture(
		sprites[dg.spritePointer].texture,
		dg.posX-CELL_WIDTH/2-DIGGER_INNER_OFFSET_X+renderOffsetsMap[dg.direction].X,
		dg.posY-CELL_HEIGHT/2-DIGGER_INNER_OFFSET_Y+renderOffsetsMap[dg.direction].Y,
		rl.White)
	//rl.DrawRectangleLinesEx(digger.getCollisionRec(), 1.0, TransparentYellow)
	rl.EndTextureMode()
}
