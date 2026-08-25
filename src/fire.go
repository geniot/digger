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
	posX                 int32
	posY                 int32
	shouldShoot          bool
	state                FireState
	sprites              map[Direction][]*TextureImage
	expSprites           map[Direction][]*TextureImage
	initialPosOffsetsMap map[Direction]Pos
	spritePtr            int
	spritePtrInc         int
	spriteExpPtr         int
	spriteExpPtrInc      int
	direction            Direction
	scene                *GameScene
}

func NewFire(scene *GameScene) *Fire {
	f := &Fire{}
	prefix := "graphics/fire/cfire"
	f.sprites = make(map[Direction][]*TextureImage)
	f.sprites[LEFT] = initTextureImages(3, prefix, 0, false, false)
	f.sprites[RIGHT] = initTextureImages(3, prefix, 0, false, false)
	f.sprites[UP] = initTextureImages(3, prefix, 0, false, false)
	f.sprites[DOWN] = initTextureImages(3, prefix, 0, false, false)
	prefix = "graphics/fire/cexp"
	f.expSprites = make(map[Direction][]*TextureImage)
	f.expSprites[LEFT] = initTextureImages(3, prefix, 0, false, false)
	f.expSprites[RIGHT] = initTextureImages(3, prefix, 0, false, false)
	f.expSprites[UP] = initTextureImages(3, prefix, 0, false, false)
	f.expSprites[DOWN] = initTextureImages(3, prefix, 0, false, false)

	f.initialPosOffsetsMap = map[Direction]Pos{
		LEFT:  {-int32(f.sprites[LEFT][f.spritePtr].width) * 2, -2},
		RIGHT: {int32(f.sprites[RIGHT][f.spritePtr].width), -2},
		UP:    {-5, -int32(f.sprites[UP][f.spritePtr].height) * 2},
		DOWN:  {-2, int32(f.sprites[DOWN][f.spritePtr].height) + 1},
	}

	f.scene = scene
	f.spritePtr = 0
	f.spriteExpPtr = 0
	f.spritePtrInc = 1
	f.spriteExpPtrInc = 1
	f.state = FireNone
	f.shouldShoot = false
	f.direction = NONE

	return f
}

func (f *Fire) Update(tick int64) {
	if f.shouldShoot && f.state == FireNone {
		f.direction = f.scene.digger.direction
		newPosX := f.scene.digger.posX + f.initialPosOffsetsMap[f.direction].X
		newPosY := f.scene.digger.posY + f.initialPosOffsetsMap[f.direction].Y
		newColRec := f.getCollisionRec(newPosX, newPosY)
		if f.scene.field.IsColliding(newColRec) {
			f.state = FireExploding
		} else {
			f.posX, f.posY = newPosX, newPosY
			f.state = FireMoving
		}
	}
	if f.state != FireNone {
		if tick%SpriteUpdateRate == 0 {
			if f.state == FireMoving {
				f.spritePtr, f.spritePtrInc = GetNextSpritePtrAndInc(f.spritePtr, f.spritePtrInc, len(f.sprites[f.direction]))
			} else if f.state == FireExploding {
				f.spriteExpPtr, f.spriteExpPtrInc = GetNextSpritePtrAndInc(f.spriteExpPtr, f.spriteExpPtrInc, len(f.expSprites[f.direction]))
			}
		}
		if f.state == FireExploding && f.spriteExpPtr+1 >= len(f.expSprites[f.direction]) {
			f.spritePtr = 0
			f.spriteExpPtr = 0
			f.state = FireNone
		} else if tick%FireSpeed == 0 {
			newPosX, newPosY := f.posX+MoveMap[f.direction].X, f.posY+MoveMap[f.direction].Y
			newColRec := f.getCollisionRec(newPosX, newPosY)
			if f.scene.field.IsColliding(newColRec) {
				f.state = FireExploding
			} else {
				f.posX, f.posY = newPosX, newPosY
			}
		}
	}
}

func (f *Fire) Render(drawTarget rl.RenderTexture2D) {
	if f.state != FireNone {
		sprites := If(f.state == FireMoving, f.sprites[f.direction], f.expSprites[f.direction])
		ptr := If(f.state == FireMoving, f.spritePtr, f.spriteExpPtr)
		rl.BeginTextureMode(drawTarget)
		rl.DrawTexture(
			sprites[ptr].texture,
			f.posX,
			f.posY,
			rl.White)
		//rl.DrawRectangleLinesEx(f.getCollisionRec(f.posX, f.posY), 1.0, TransparentYellow)
		rl.EndTextureMode()
	}
}

func (f *Fire) getCollisionRec(posX int32, posY int32) rl.Rectangle {
	sprites := If(f.state == FireMoving, f.sprites[f.direction], f.expSprites[f.direction])
	ptr := If(f.state == FireMoving, f.spritePtr, f.spriteExpPtr)
	return rl.Rectangle{
		X:      float32(posX),
		Y:      float32(posY),
		Width:  sprites[ptr].width,
		Height: sprites[ptr].height,
	}
}
