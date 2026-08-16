package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fire struct {
	posX  int32
	posY  int32
	scene *GameScene
}

func NewFire(scene *GameScene) *Fire {
	fire := &Fire{}
	//fire.posX = x*CellWidth + FieldOffsetX + CellWidth/2 - 1 // +1 in the original game, not centered, why?
	//fire.posY = y*CellHeight + FieldOffsetY + CellHeight/2 + 1
	fire.scene = scene
	return fire
}

func (e *Fire) Update(_ int64) {
}

func (e *Fire) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	//rl.DrawTexture(
	//	sprite.texture,
	//	e.posX-int32(sprite.width/2),
	//	e.posY-int32(sprite.height/2),
	//	rl.White)
	//rl.DrawRectangleLinesEx(e.getCollisionRec(), 1.0, TransparentBlue)
	rl.EndTextureMode()
}

func (e *Fire) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{}
	//return rl.Rectangle{
	//	X:      float32(e.posX - int32(sprite.width/2)),
	//	Y:      float32(e.posY - int32(sprite.height/2)),
	//	Width:  sprite.width,
	//	Height: sprite.height,
	//}
}
