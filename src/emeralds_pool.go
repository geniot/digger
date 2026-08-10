package main

import rl "github.com/gen2brain/raylib-go/raylib"

type EmeraldsPool struct {
	scene *GameScene
}

func NewEmeraldsPool(scene *GameScene) *EmeraldsPool {
	emeraldsPool := &EmeraldsPool{}
	return emeraldsPool
}

func (ep *EmeraldsPool) Update(tick int64) {

}

func (ep *EmeraldsPool) Render(drawTarget rl.RenderTexture2D) {

}
