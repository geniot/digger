package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EmeraldsPool struct {
	scene      *GameScene
	sprite     *MaskedTexture
	spriteMask *MaskedTexture
	emeralds   mapset.Set[*Emerald]
}

func NewEmeraldsPool(scene *GameScene) *EmeraldsPool {
	emeraldsPool := &EmeraldsPool{}
	emeraldsPool.scene = scene
	emeraldsPool.emeralds = mapset.NewThreadUnsafeSet[*Emerald]()
	emeraldsPool.sprite = NewMaskedTexture("graphics/emerald/emerald.png", 0, false, false, false)
	emeraldsPool.spriteMask = NewMaskedTexture("graphics/emerald/emerald.png", 0, false, false, true)
	lp := LevelPlan(scene.level)
	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			c := getLevelChar(x, y, lp)
			if c == 'C' {
				emeraldsPool.emeralds.Add(NewEmerald(emeraldsPool, x, y))
			}
		}
	}
	return emeraldsPool
}

func (ep *EmeraldsPool) Update(tick int64) {
	for emerald := range ep.emeralds.Iter() {
		emerald.Update(tick)
	}
}

func (ep *EmeraldsPool) Render() {
	for emerald := range ep.emeralds.Iter() {
		emerald.Render()
	}
}

func (ep *EmeraldsPool) handle(dg *Digger) {
	for emerald := range ep.emeralds.Iter() {
		if rl.CheckCollisionRecs(emerald.getCollisionRec(), dg.getCollisionRec()) {
			ep.scene.field.draw(ep.spriteMask, emerald.posX-ep.spriteMask.width/2, emerald.posY-ep.spriteMask.height/2)
			ep.emeralds.Remove(emerald)
		}
	}
}
