package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DiggerState int64

const (
	DiggerAlive DiggerState = iota
	DiggerDie
	DiggerGrave
)

type Digger struct {
	scene               *GameScene
	posX                int32
	posY                int32
	width               int32
	height              int32
	direction           Direction
	requestedDirection  Direction
	shouldMove          bool
	spritePtr           int
	spritePtrInc        int
	sprites             map[Direction][]*MaskedTexture
	renderOffsetsMap    map[Direction]Pos
	collisionOffsetsMap map[Direction]Pos
	collisionSizeMap    map[Direction]WidthHeight
	innerOffsetX        int32
	innerOffsetY        int32
	state               DiggerState
}

func NewDigger(scene *GameScene) *Digger {
	dg := &Digger{}
	dg.scene = scene

	prefix := "graphics/digger/cldig"
	dg.sprites = make(map[Direction][]*MaskedTexture)
	dg.sprites[LEFT] = initMaskedTextures(3, prefix, 0, false, false)
	dg.sprites[RIGHT] = initMaskedTextures(3, prefix, 0, true, false)
	dg.sprites[UP] = initMaskedTextures(3, prefix, 90, false, false)
	dg.sprites[DOWN] = initMaskedTextures(3, prefix, 90, true, true)

	dg.spritePtr = 0
	dg.spritePtrInc = 1
	dg.width = 16
	dg.height = 16
	dg.posX, dg.posY = scene.moveGrid.getDiggerStartPos()
	dg.direction = RIGHT
	dg.shouldMove = false
	dg.state = DiggerAlive
	dg.innerOffsetX = 0
	dg.innerOffsetY = 1

	dg.renderOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 0},
		RIGHT: {0, 0},
		UP:    {-1, 0},
		DOWN:  {1, 1},
	}
	dg.collisionOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 2},
		RIGHT: {0, 2},
		UP:    {0, 0},
		DOWN:  {2, 1},
	}
	dg.collisionSizeMap = map[Direction]WidthHeight{
		LEFT:  {0, -2},
		RIGHT: {0, -2},
		UP:    {-2, 0},
		DOWN:  {-2, 0},
	}
	return dg
}

func (dg *Digger) Update(tick int64) {
	if tick%SpriteUpdateRate == 0 {
		dg.spritePtr, dg.spritePtrInc = GetNextSpritePtrAndInc(dg.spritePtr, dg.spritePtrInc, len(dg.sprites[dg.direction]))
	}
	if tick%DiggerSpeed == 0 && dg.shouldMove {
		posX, posY, dir := dg.scene.moveGrid.move(dg.posX, dg.posY, dg.direction, dg.requestedDirection)
		if dg.posX != posX || dg.posY != posY || dg.direction != dir { //any change from previous state?
			dg.scene.bagsPool.handle(dg)
			dg.posX, dg.posY, dg.direction = posX, posY, dir
			dg.scene.emeraldsPool.handle(dg)
			if dir == RIGHT {
				blob := dg.scene.field.rightBlob
				dg.scene.field.drawExt(blob, dg.posX-dg.posX%4+4, dg.posY-CellHeight/2+1, false, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX+8, dg.posY-CellHeight/2+1, true, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX+7, dg.posY-CellHeight/2+1, true, false, true, false)
			} else if dir == LEFT {
				blob := dg.scene.field.leftBlob
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2-2, dg.posY-CellHeight/2+1, false, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2-1, dg.posY-CellHeight/2+1, false, false, true, false)
				//drawing all indentation when approaching the left side of the field (<=20)
				dg.scene.field.drawExt(blob, dg.posX-dg.posX%4-CellWidth/2+IfInt(dg.posX <= 20, 2, 6), dg.posY-CellHeight/2+1, true, false, true, false)
			} else if dir == UP {
				blob := dg.scene.field.upBlob
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2-dg.posY%3+4, false, true, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2-1, false, false, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2, false, false, false, true)
			} else if dir == DOWN {
				blob := dg.scene.field.downBlob
				//drawing all indentation when approaching the bottom side of the field (>=173)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY+CellHeight/2-dg.posY%3-IfInt(dg.posY >= 173, 2, 5), false, false, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY+CellHeight/2-1, false, true, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY+CellHeight/2-2, false, true, false, true)
			}

		}
	}
}

func (dg *Digger) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(dg.posX + (CellWidth-dg.width)/2 - CellWidth/2 - dg.innerOffsetX + dg.collisionOffsetsMap[dg.direction].X),
		Y:      float32(dg.posY + (CellWidth-dg.height)/2 - CellHeight/2 - dg.innerOffsetY + dg.collisionOffsetsMap[dg.direction].Y),
		Width:  float32(dg.width + dg.collisionSizeMap[dg.direction].W),
		Height: float32(dg.height + dg.collisionSizeMap[dg.direction].H),
	}
}

func (dg *Digger) Render() {
	sprites := dg.sprites[dg.direction]
	rl.DrawTexture(
		sprites[dg.spritePtr].texture,
		dg.posX-CellWidth/2-dg.innerOffsetX+dg.renderOffsetsMap[dg.direction].X,
		dg.posY-CellHeight/2-dg.innerOffsetY+dg.renderOffsetsMap[dg.direction].Y,
		rl.White)
	//rl.DrawRectangleLinesEx(dg.getCollisionRec(), 1.0, TransparentYellow)
}
