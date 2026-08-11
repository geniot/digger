package main

import rl "github.com/gen2brain/raylib-go/raylib"

type EmeraldsPool struct {
	scene    *GameScene
	emeralds map[int32]*Emerald
}

func NewEmeraldsPool(scene *GameScene) *EmeraldsPool {
	emeraldsPool := &EmeraldsPool{}
	emeraldsPool.emeralds = make(map[int32]*Emerald)
	lp := LevelPlan(scene.level)
	counter := int32(0)
	for x := int32(0); x < 15; x++ {
		for y := int32(0); y < 10; y++ {
			c := getLevelChar(x, y, lp)
			if c == 'C' {
				emeraldsPool.emeralds[counter] = NewEmerald(emeraldsPool, x, y)
				counter++
			}
		}
	}
	return emeraldsPool
}

func (ep *EmeraldsPool) Update(tick int64) {
	for _, emerald := range ep.emeralds {
		emerald.Update(tick)
	}
}

func (ep *EmeraldsPool) Render(drawTarget rl.RenderTexture2D) {
	for _, emerald := range ep.emeralds {
		emerald.Render(drawTarget)
	}
}
