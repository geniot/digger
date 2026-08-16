package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MonsterState int64

const (
	MonsterNobbin MonsterState = iota
	MonsterHobbin
	MonsterDie
)

type Monster struct {
	posX         int32
	posY         int32
	monstersPool *MonstersPool
	state        MonsterState
}

func NewMonster(monstersPool *MonstersPool, x int32, y int32) *Monster {
	mr := &Monster{}
	mr.posX = x*CellWidth + FieldOffsetX + CellWidth/2 - 1 // +1 in the original game, not centered, why?
	mr.posY = y*CellHeight + FieldOffsetY + CellHeight/2 + 1
	mr.monstersPool = monstersPool
	mr.state = MonsterNobbin
	return mr
}

func (mr *Monster) Update(_ int64) {
}

func (mr *Monster) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	sprite := mr.monstersPool.sprite
	rl.DrawTexture(
		sprite.texture,
		mr.posX-int32(sprite.width/2),
		mr.posY-int32(sprite.height/2),
		rl.White)
	//rl.DrawRectangleLinesEx(e.getCollisionRec(), 1.0, TransparentBlue)
	rl.EndTextureMode()
}

func (mr *Monster) getCollisionRec() rl.Rectangle {
	sprite := mr.monstersPool.sprite
	return rl.Rectangle{
		X:      float32(mr.posX - int32(sprite.width/2)),
		Y:      float32(mr.posY - int32(sprite.height/2)),
		Width:  sprite.width,
		Height: sprite.height,
	}
}
