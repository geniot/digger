package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FireState int64

const (
	FireMoving FireState = iota
	FireExploding
	FireNone
)

type Fire struct {
	posX                      int32
	posY                      int32
	shouldShoot               bool
	state                     FireState
	sprites                   map[Direction][]*TextureImage
	explosionSprites          map[Direction][]*TextureImage
	spritePointer             int
	spritePointerInc          int
	spriteExplosionPointer    int
	spriteExplosionPointerInc int
	direction                 Direction
	scene                     *GameScene
}

func NewFire(scene *GameScene) *Fire {
	f := &Fire{}

	prefix := "graphics/fire/cfire"
	f.sprites = make(map[Direction][]*TextureImage)
	f.sprites[LEFT] = initSprites(3, prefix, 0, false, false)
	f.sprites[RIGHT] = initSprites(3, prefix, 0, true, false)
	f.sprites[UP] = initSprites(3, prefix, 90, false, false)
	f.sprites[DOWN] = initSprites(3, prefix, 90, true, true)
	//no rotations for explosion
	prefix = "graphics/fire/cexp"
	f.explosionSprites = make(map[Direction][]*TextureImage)
	f.explosionSprites[LEFT] = initSprites(3, prefix, 0, false, false)
	f.explosionSprites[RIGHT] = initSprites(3, prefix, 0, true, false)
	f.explosionSprites[UP] = initSprites(3, prefix, 0, false, false)
	f.explosionSprites[DOWN] = initSprites(3, prefix, 0, false, false)

	f.scene = scene
	f.state = FireNone
	f.shouldShoot = false
	f.direction = NONE
	return f
}

func (f *Fire) Update(tick int64) {
	if f.shouldShoot && f.state == FireNone {
		f.posX, f.posY = f.scene.digger.posX, f.scene.digger.posY
		f.direction = f.scene.digger.direction
		f.state = FireMoving
	}
	if f.state != FireNone {
		if tick%SpriteUpdateRate == 0 {
			f.spritePointer, f.spritePointerInc = GetNextSpritePointerAndInc(f.spritePointer, f.spritePointerInc, len(f.sprites))
			f.spriteExplosionPointer, f.spriteExplosionPointerInc = GetNextSpritePointerAndInc(f.spriteExplosionPointer, f.spriteExplosionPointerInc, len(f.explosionSprites))
		}
		if tick%FireSpeed == 0 {

		}
	}
}

func (f *Fire) Render(drawTarget rl.RenderTexture2D) {
	if f.state != FireNone {
		sprites := f.sprites[f.direction]
		rl.BeginTextureMode(drawTarget)
		rl.DrawTexture(
			sprites[f.spritePointer].texture,
			f.posX,
			f.posY,
			rl.White)
		//rl.DrawRectangleLinesEx(e.getCollisionRec(), 1.0, TransparentBlue)
		rl.EndTextureMode()
	}
}

func (f *Fire) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{}
	//return rl.Rectangle{
	//	X:      float32(e.posX - int32(sprite.width/2)),
	//	Y:      float32(e.posY - int32(sprite.height/2)),
	//	Width:  sprite.width,
	//	Height: sprite.height,
	//}
}
