package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BagsPool struct {
	scene  *GameScene
	sprite *TextureImage
	bags   mapset.Set[*Bag]
}

func NewBagsPool(scene *GameScene) *BagsPool {
	bagsPool := &BagsPool{}
	bagsPool.scene = scene
	bagsPool.bags = mapset.NewThreadUnsafeSet[*Bag]()
	bagsPool.sprite = NewTextureImage("graphics/bag/csbag.png", 0, false, false, false)
	lp := LevelPlan(scene.level)
	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			c := getLevelChar(x, y, lp)
			if c == 'B' {
				bagsPool.bags.Add(NewBag(bagsPool, x, y))
			}
		}
	}
	return bagsPool
}

func (ep *BagsPool) Update(tick int64) {
	for bag := range ep.bags.Iter() {
		bag.Update(tick)
	}
}

func (ep *BagsPool) Render(drawTarget rl.RenderTexture2D) {
	for bag := range ep.bags.Iter() {
		bag.Render(drawTarget)
	}
}

func (ep *BagsPool) handle(dg *Digger) {
	for bag := range ep.bags.Iter() {
		if rl.CheckCollisionRecs(bag.getCollisionRec(), dg.getCollisionRec()) {
			ep.bags.Remove(bag)
		}
	}
}
