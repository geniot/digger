package model

import (
	"geniot.com/geniot/digger/internal/utils"
	"github.com/veandco/go-sdl2/sdl"
)

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

func NewBall() *Ball {
	return &Ball{}
}

func (ball Ball) Render(renderer *sdl.Renderer) {
	width, height, _ := renderer.GetOutputSize()

	alpha := uint8(255)
	radius := width / 3
	if height/3 < radius {
		radius = height / 3
	}

	renderer.SetDrawColor(0, 0, 0, alpha)

	colourInside := [3]uint8{0, 0, 0}
	//colourAround := [3]uint8{192, 192, 192}

	//DrawSimpleCircle(renderer, 100.0, 100.0, 100.0, 100.0, 100.0)
	//DrawCircle(renderer, 100.0, 100.0, 100.0, float64(width/2), float64(height/2))
	//utils.AaellipseRGBA(renderer, width/2, height/2, radius+1, radius, colourAround[0], colourAround[1], colourAround[2], alpha)
	utils.AaellipseRGBA(renderer, width/2, height/2, radius, radius, colourInside[0], colourInside[1], colourInside[2], alpha)
	//utils.AaellipseRGBA(renderer, width/2, height/2, radius-1, radius-1, colourAround[0], colourAround[1], colourAround[2], alpha)
	//ball.drawCircle(renderer, width/2, height/2, radius)
}
