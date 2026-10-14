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
	direction           Direction
	requestedDirection  Direction
	shouldMove          bool
	sprites             map[Direction]*Sprite
	seqPtr              int
	renderOffsetsMap    map[Direction]Pos
	collisionOffsetsMap map[Direction]Pos
	collisionSizeMap    map[Direction]WidthHeight
	state               DiggerState
}

func NewDigger(scene *GameScene) *Digger {
	dg := &Digger{}
	dg.scene = scene

	prefix := "graphics/digger/cldig"
	dg.sprites = make(map[Direction]*Sprite)
	dg.sprites[DirLeft] = NewSprite([]int{0, 1, 2, 1}, prefix, 0, false, false)
	dg.sprites[DirRight] = NewSprite([]int{0, 1, 2, 1}, prefix, 0, true, false)
	dg.sprites[DirUp] = NewSprite([]int{0, 1, 2, 1}, prefix, 90, false, false)
	dg.sprites[DirDown] = NewSprite([]int{0, 1, 2, 1}, prefix, 90, true, true)

	dg.seqPtr = 0
	dg.posX, dg.posY = scene.moveGrid.getDiggerStartPos()
	dg.direction = DirRight
	dg.shouldMove = false
	dg.state = DiggerAlive

	dg.renderOffsetsMap = map[Direction]Pos{
		DirLeft:  {-10, -10},
		DirRight: {-10, -10},
		DirUp:    {-11, -11},
		DirDown:  {-9, -10},
	}
	dg.collisionOffsetsMap = map[Direction]Pos{
		DirLeft:  {-8, -6},
		DirRight: {-8, -6},
		DirUp:    {-8, -9},
		DirDown:  {-6, -8},
	}
	dg.collisionSizeMap = map[Direction]WidthHeight{
		DirLeft:  {16, 14},
		DirRight: {16, 14},
		DirUp:    {14, 16},
		DirDown:  {14, 16},
	}
	return dg
}

func (dg *Digger) Update(tick int64) {
	if tick%SpriteUpdateRate == 0 {
		dg.seqPtr = dg.sprites[dg.direction].Next(dg.seqPtr)
	}
	if tick%DiggerSpeed == 0 && dg.shouldMove {
		posX, posY, dir := dg.scene.moveGrid.move(dg.posX, dg.posY, dg.direction, dg.requestedDirection)
		if dg.posX != posX || dg.posY != posY || dg.direction != dir { //any change from previous state?
			dg.scene.bagsPool.handle(dg)
			dg.posX, dg.posY, dg.direction = posX, posY, dir
			dg.scene.emeraldsPool.handle(dg)
			if dir == DirRight {
				blob := dg.scene.field.rightBlob
				dg.scene.field.drawExt(blob, dg.posX-dg.posX%4+4, dg.posY-CellHeight/2+1, false, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX+8, dg.posY-CellHeight/2+1, true, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX+7, dg.posY-CellHeight/2+1, true, false, true, false)
			} else if dir == DirLeft {
				blob := dg.scene.field.leftBlob
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2-2, dg.posY-CellHeight/2+1, false, false, true, false)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2-1, dg.posY-CellHeight/2+1, false, false, true, false)
				//drawing all indentation when approaching the left side of the field (<=20)
				dg.scene.field.drawExt(blob, dg.posX-dg.posX%4-CellWidth/2+IfInt(dg.posX <= 20, 2, 6), dg.posY-CellHeight/2+1, true, false, true, false)
			} else if dir == DirUp {
				blob := dg.scene.field.upBlob
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2-dg.posY%3+4, false, true, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2-1, false, false, false, true)
				dg.scene.field.drawExt(blob, dg.posX-CellWidth/2, dg.posY-CellHeight/2, false, false, false, true)
			} else if dir == DirDown {
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
		X:      float32(dg.posX + dg.collisionOffsetsMap[dg.direction].X),
		Y:      float32(dg.posY + dg.collisionOffsetsMap[dg.direction].Y),
		Width:  float32(dg.collisionSizeMap[dg.direction].W),
		Height: float32(dg.collisionSizeMap[dg.direction].H),
	}
}

func (dg *Digger) Render() {
	rl.DrawTexture(
		dg.sprites[dg.direction].GetTexture(dg.seqPtr).texture,
		dg.posX+dg.renderOffsetsMap[dg.direction].X,
		dg.posY+dg.renderOffsetsMap[dg.direction].Y,
		rl.White)
	//rl.DrawRectangleLinesEx(dg.getCollisionRec(), 1.0, TransparentYellow)
}
