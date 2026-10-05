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
	posX     int32
	posY     int32
	bagsPool *BagsPool
	state    BagState
}

func NewBag(bagsPool *BagsPool, x int32, y int32) *Bag {
	bg := &Bag{}
	bg.posX = x*CellWidth + FieldOffsetX + CellWidth/2
	bg.posY = y*CellHeight + FieldOffsetY + CellHeight/2 + FieldVerticalOffset
	bg.bagsPool = bagsPool
	bg.state = BagSet
	return bg
}

func (bg *Bag) Update(_ int64) {
}

func (bg *Bag) Render() {
	sprite := bg.bagsPool.sprite
	rl.DrawTexture(
		sprite.texture,
		bg.posX-sprite.width/2,
		bg.posY-sprite.height/2,
		rl.White)
	//rl.DrawRectangleLinesEx(bg.getCollisionRec(), 1.0, TransparentBlue)
}

func (bg *Bag) getCollisionRec() rl.Rectangle {
	sprite := bg.bagsPool.sprite
	return rl.Rectangle{
		X:      float32(bg.posX - sprite.width/2 + 2),
		Y:      float32(bg.posY - sprite.height/2 + 3),
		Width:  float32(sprite.width - 4),
		Height: float32(sprite.height - 5),
	}
}
