package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MonstersPool struct {
	scene      *GameScene
	sprite     *TextureImage
	spriteMask *TextureImage
	monsters   mapset.Set[*Monster]
}

func NewMonstersPool(scene *GameScene) *MonstersPool {
	monstersPool := &MonstersPool{}
	monstersPool.scene = scene
	monstersPool.monsters = mapset.NewThreadUnsafeSet[*Monster]()
	//monstersPool.sprite = NewTextureImage("monster.png", 0, false, false, false)
	return monstersPool
}

func (mp *MonstersPool) Update(tick int64) {
	for monster := range mp.monsters.Iter() {
		monster.Update(tick)
	}
}

func (mp *MonstersPool) Render(drawTarget rl.RenderTexture2D) {
	for monster := range mp.monsters.Iter() {
		monster.Render(drawTarget)
	}
}

func (mp *MonstersPool) handle(dg *Digger) {
	for monster := range mp.monsters.Iter() {
		if rl.CheckCollisionRecs(monster.getCollisionRec(), dg.getCollisionRec()) {
			mp.monsters.Remove(monster)
		}
	}
}
