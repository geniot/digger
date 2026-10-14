package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BagsPool struct {
	scene   *GameScene
	sprites map[BagState]*Sprite
	bags    mapset.Set[*Bag]
}

func NewBagsPool(scene *GameScene) *BagsPool {
	bagsPool := &BagsPool{}
	bagsPool.scene = scene
	bagsPool.bags = mapset.NewThreadUnsafeSet[*Bag]()

	bagsPool.sprites = make(map[BagState]*Sprite)
	bagsPool.sprites[BagSet] = NewSprite([]int{0}, "graphics/bag/csbag", 0, false, false)
	bagsPool.sprites[BagFalling] = NewSprite([]int{0}, "graphics/bag/cfbag", 0, false, false)
	bagsPool.sprites[BagWobble] = NewSprite([]int{0, 1, 2, 1, 2}, "graphics/bag/csbag", 0, false, false)
	bagsPool.sprites[BagGold] = NewSprite([]int{0, 1, 2}, "graphics/bag/cgold", 0, false, false)

	levelPlan := LevelPlan(scene.level)
	for x := range CellsHorizontal {
		for y := range CellsVertical {
			levelChar := getLevelChar(x, y, levelPlan)
			if levelChar == 'B' {
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
