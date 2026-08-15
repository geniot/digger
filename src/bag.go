package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bag struct {
	posX     int32
	posY     int32
	bagsPool *BagsPool
}

func NewBag(bagsPool *BagsPool, x int32, y int32) *Bag {
	bg := &Bag{}
	bg.posX = x*CELL_WIDTH + FIELD_OFFSET_X + CELL_WIDTH/2
	bg.posY = y*CELL_HEIGHT + FIELD_OFFSET_Y + CELL_HEIGHT/2
	bg.bagsPool = bagsPool
	return bg
}

func (bg *Bag) Update(_ int64) {
}

func (bg *Bag) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	sprite := bg.bagsPool.sprite
	rl.DrawTexture(
		sprite.texture,
		bg.posX-int32(sprite.width/2),
		bg.posY-int32(sprite.height/2),
		rl.White)
	rl.DrawRectangleLinesEx(bg.getCollisionRec(), 1.0, TransparentBlue)
	rl.EndTextureMode()
}

func (bg *Bag) getCollisionRec() rl.Rectangle {
	sprite := bg.bagsPool.sprite
	return rl.Rectangle{
		X:      float32(bg.posX - int32(sprite.width/2) + 2),
		Y:      float32(bg.posY - int32(sprite.height/2) + 3),
		Width:  sprite.width - 4,
		Height: sprite.height - 5,
	}
}
