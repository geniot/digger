package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BagsPool struct {
	scene                        *GameScene
	sprites                      map[BagState][]*MaskedTexture
	bagSpritesShakeFrameSequence []int
	bagSpritesGoldFrameSequence  []int
	bags                         mapset.Set[*Bag]
}

func NewBagsPool(scene *GameScene) *BagsPool {
	bagsPool := &BagsPool{}
	bagsPool.scene = scene
	bagsPool.bags = mapset.NewThreadUnsafeSet[*Bag]()

	bagsPool.sprites = make(map[BagState][]*MaskedTexture)
	bagsPool.sprites[BagSet] = initMaskedTextures(1, "graphics/bag/csbag", 0, false, false)
	bagsPool.sprites[BagWobble] = initMaskedTextures(3, "graphics/bag/csbag", 0, false, false)
	bagsPool.sprites[BagGold] = initMaskedTextures(3, "graphics/bag/cgold", 0, false, false)
	bagsPool.sprites[BagFalling] = initMaskedTextures(1, "graphics/bag/cfbag", 0, false, false)
	bagsPool.bagSpritesShakeFrameSequence = []int{0, 1, 2, 1, 2}
	bagsPool.bagSpritesGoldFrameSequence = []int{0, 1, 2}

	lp := LevelPlan(scene.level)
	for x := range int32(15) {
		for y := range int32(10) {
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

func (ep *BagsPool) Render() {
	for bag := range ep.bags.Iter() {
		bag.Render()
	}
}

func (ep *BagsPool) handle(dg *Digger) {
	for bag := range ep.bags.Iter() {
		if rl.CheckCollisionRecs(bag.getCollisionRec(), dg.getCollisionRec()) {
			ep.bags.Remove(bag)
		}
	}
}
