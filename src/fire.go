package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FireState int64

const (
	FireNone FireState = iota
	FireMoving
	FireExploding
)

type Fire struct {
	posX                 int32
	posY                 int32
	shouldShoot          bool
	state                FireState
	sprites              map[FireState]*Sprite
	initialPosOffsetsMap map[Direction]Pos
	seqPtr               int
	direction            Direction
	scene                *GameScene
}

func NewFire(scene *GameScene) *Fire {
	f := &Fire{}
	f.sprites = make(map[FireState]*Sprite)
	f.sprites[FireMoving] = NewSprite([]int{0, 1, 2, 1}, "graphics/fire/cfire", 0, false, false)
	f.sprites[FireExploding] = NewSprite([]int{0, 1, 2, 1}, "graphics/fire/cexp", 0, false, false)

	f.scene = scene
	f.seqPtr = 0
	f.state = FireNone
	f.shouldShoot = false
	f.direction = DirNone

	f.initialPosOffsetsMap = map[Direction]Pos{
		DirLeft:  {-f.sprites[FireMoving].GetTexture(f.seqPtr).width * 2, -2},
		DirRight: {f.sprites[FireMoving].GetTexture(f.seqPtr).width, -2},
		DirUp:    {-5, -f.sprites[FireMoving].GetTexture(f.seqPtr).height * 2},
		DirDown:  {-2, int32(f.sprites[FireMoving].GetTexture(f.seqPtr).height) + 1},
	}

	return f
}

func (f *Fire) Update(tick int64) {
	switch f.state {
	case FireNone:
		{
			if f.shouldShoot { //initiate the fire shot
				f.direction = f.scene.digger.direction
				newPosX := f.scene.digger.posX + f.initialPosOffsetsMap[f.direction].X
				newPosY := f.scene.digger.posY + f.initialPosOffsetsMap[f.direction].Y
				newColRec := f.getCollisionRec(newPosX, newPosY)
				f.posX, f.posY = newPosX, newPosY
				if f.scene.field.IsColliding(newColRec, true) {
					f.state = FireNone
				} else {
					f.state = FireMoving
				}
			}
		}
	case FireMoving:
		{
			if tick%SpriteUpdateRate == 0 {
				f.seqPtr = f.sprites[f.state].Next(f.seqPtr)
			}
			if tick%FireSpeed == 0 {
				newPosX, newPosY := f.posX+MoveMap[f.direction].X, f.posY+MoveMap[f.direction].Y
				newColRec := f.getCollisionRec(newPosX, newPosY)
				if f.scene.field.IsColliding(newColRec, true) {
					f.state = FireExploding
				} else {
					f.posX, f.posY = newPosX, newPosY
				}
			}
		}
	case FireExploding:
		{
			if tick%SpriteUpdateRate == 0 {
				f.seqPtr = f.sprites[f.state].Next(f.seqPtr)
				if f.seqPtr == 0 {
					f.state = FireNone
				}
			}
		}
	default:
		panic(fmt.Sprintf("unhandled default case: %d", f.state))
	}
}

func (f *Fire) Render() {
	if f.state != FireNone {
		rl.DrawTexture(
			f.sprites[f.state].GetTexture(f.seqPtr).texture,
			f.posX,
			f.posY,
			rl.White)
		//rl.DrawRectangleLinesEx(f.getCollisionRec(f.posX, f.posY), 1.0, TransparentYellow)
	}
}

func (f *Fire) getCollisionRec(posX int32, posY int32) rl.Rectangle {
	state := If(f.state == FireNone, FireMoving, f.state) //there are no sprites for FireNone
	return rl.Rectangle{
		X:      float32(posX),
		Y:      float32(posY),
		Width:  float32(f.sprites[state].GetTexture(f.seqPtr).width),
		Height: float32(f.sprites[state].GetTexture(f.seqPtr).height),
	}
}
