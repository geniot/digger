package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BagState int64

const (
	BagSet BagState = iota
	BagPushed
	BagHold
	BagMoving
	BagWobble
	BagFalling
	BagGoldFalling
	BagGold
)

type Bag struct {
	posX      int32
	posY      int32
	bagsPool  *BagsPool
	state     BagState
	spritePtr int
}

func NewBag(bagsPool *BagsPool, x int32, y int32) *Bag {
	bg := &Bag{}
	bg.posX = x*CellWidth + FieldOffsetX + CellWidth/2
	bg.posY = y*CellHeight + FieldOffsetY + CellHeight/2
	bg.bagsPool = bagsPool
	bg.state = BagSet
	bg.spritePtr = 0
	return bg
}

func (bg *Bag) Update(n int64) {
	switch bg.state {
	case BagSet:
		{
			if bg.bagsPool.scene.field.IsColliding(bg.getFallThroughRec(), false) {
				bg.spritePtr = 0
				bg.state = BagWobble
			}
		}
	case BagWobble:
		{
			if n%(SpriteUpdateRate*3) == 0 {
				bg.spritePtr += 1
				if bg.spritePtr >= len(bg.bagsPool.bagSpritesShakeFrameSequence) {
					bg.spritePtr = 0
					bg.state = BagFalling
				}
			}
		}
	default:
		panic("unhandled default case")
	}
}

func (bg *Bag) Render() {
	sprite := bg.bagsPool.sprites[bg.state][bg.spritePtr]
	rl.DrawTexture(
		sprite.texture,
		bg.posX-sprite.width/2,
		bg.posY-sprite.height/2,
		rl.White)
	rl.DrawRectangleLinesEx(bg.getCollisionRec(), 1.0, TransparentBlue)
	rl.DrawRectangleLinesEx(bg.getFallThroughRec(), 1.0, TransparentYellow)
}

func (bg *Bag) getCollisionRec() rl.Rectangle {
	sprite := bg.bagsPool.sprites[bg.state][bg.spritePtr]
	return rl.Rectangle{
		X:      float32(bg.posX - sprite.width/2 + 2),
		Y:      float32(bg.posY - sprite.height/2 + 3),
		Width:  float32(sprite.width - 4),
		Height: float32(sprite.height - 5),
	}
}

func (bg *Bag) getFallThroughRec() rl.Rectangle {
	collisionRec := bg.getCollisionRec()
	return rl.Rectangle{
		X:      collisionRec.X,
		Y:      collisionRec.Y + collisionRec.Height,
		Width:  collisionRec.Width,
		Height: collisionRec.Height + 3,
	}
}
