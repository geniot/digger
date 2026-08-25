package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fire struct {
	posX                      int32
	posY                      int32
	shouldShoot               bool
	isOn                      bool
	sprites                   []*TextureImage
	explosionSprites          []*TextureImage
	spritePointer             int
	spritePointerInc          int
	spriteExplosionPointer    int
	spriteExplosionPointerInc int
	direction                 Direction
	scene                     *GameScene
}

func NewFire(scene *GameScene) *Fire {
	fire := &Fire{}
	//fire.posX = x*CellWidth + FieldOffsetX + CellWidth/2 - 1 // +1 in the original game, not centered, why?
	//fire.posY = y*CellHeight + FieldOffsetY + CellHeight/2 + 1
	fire.scene = scene
	fire.isOn = false
	fire.shouldShoot = false
	fire.direction = NONE
	return fire
}

func (e *Fire) Update(tick int64) {
	if tick%SpriteUpdateRate == 0 {
		e.spritePointer, e.spritePointerInc = GetNextSpritePointerAndInc(e.spritePointer, e.spritePointerInc, len(e.sprites))
	}
	if tick%FireSpeed == 0 && e.shouldShoot {

	}
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
