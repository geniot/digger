package model

import "github.com/veandco/go-sdl2/sdl"

type Ball struct {
	x         float32
	y         float32
	width     float32
	height    float32
	xVelocity float32
	yVelocity float32

	born       uint32
	lastUpdate uint32
}

func NewBall() Ball {
	return Ball{}
}

func (ball Ball) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.DrawRect(&sdl.Rect{20, 20, 100, 100})
}
