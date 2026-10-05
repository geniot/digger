package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Score struct {
	scene *GameScene
}

func NewScore(scene *GameScene) *Score {
	score := &Score{}
	score.scene = scene
	return score
}

func (ep *Score) Update(tick int64) {

}

func (ep *Score) Render() {
	rl.DrawText("Score: ", 0, 0, 10, rl.White)
}
