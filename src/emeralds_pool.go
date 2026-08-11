package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EmeraldsPool struct {
	scene      *GameScene
	sprite     *TextureImage
	spriteMask *TextureImage
	emeralds   mapset.Set[*Emerald]
}

func NewEmeraldsPool(scene *GameScene) *EmeraldsPool {
	emeraldsPool := &EmeraldsPool{}
	emeraldsPool.scene = scene
	emeraldsPool.emeralds = mapset.NewThreadUnsafeSet[*Emerald]()
	emeraldsPool.sprite = NewTextureImage("emerald.png", 0, false, false)
	emeraldsPool.spriteMask = NewTextureImage("emerald_mask.png", 0, false, false)
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

func (ep *EmeraldsPool) Render(drawTarget rl.RenderTexture2D) {
	for emerald := range ep.emeralds.Iter() {
		emerald.Render(drawTarget)
	}
}

func (ep *EmeraldsPool) handle(dg *Digger) {
	for emerald := range ep.emeralds.Iter() {
		if rl.CheckCollisionRecs(emerald.getCollisionRec(), dg.getCollisionRec()) {
			ep.scene.field.draw(ep.spriteMask, float32(emerald.posX)-ep.spriteMask.width/2, float32(emerald.posY)-ep.spriteMask.height/2)
			ep.emeralds.Remove(emerald)
		}
	}
}
