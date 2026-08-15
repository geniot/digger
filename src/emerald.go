package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Emerald struct {
	posX         int32
	posY         int32
	emeraldsPool *EmeraldsPool
}

func NewEmerald(emeraldsPool *EmeraldsPool, x int32, y int32) *Emerald {
	emerald := &Emerald{}
	emerald.posX = x*CELL_WIDTH + FIELD_OFFSET_X + CELL_WIDTH/2 - 1 // +1 in the original game, not centered, why?
	emerald.posY = y*CELL_HEIGHT + FIELD_OFFSET_Y + CELL_HEIGHT/2 + 1
	emerald.emeraldsPool = emeraldsPool
	return emerald
}

func (e *Emerald) Update(_ int64) {
}

func (e *Emerald) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	sprite := e.emeraldsPool.sprite
	rl.DrawTexture(
		sprite.texture,
		e.posX-int32(sprite.width/2),
		e.posY-int32(sprite.height/2),
		rl.White)
	//rl.DrawRectangleLinesEx(e.getCollisionRec(), 1.0, TransparentBlue)
	rl.EndTextureMode()
}

func (e *Emerald) getCollisionRec() rl.Rectangle {
	sprite := e.emeraldsPool.sprite
	return rl.Rectangle{
		X:      float32(e.posX - int32(sprite.width/2)),
		Y:      float32(e.posY - int32(sprite.height/2)),
		Width:  sprite.width,
		Height: sprite.height,
	}
}
