package model

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	balls mapset.Set[Ball]
}

func NewScene() Scene {
	set := mapset.NewSet[Ball]()
	set.Add(NewBall())
	return Scene{set}
}

func (scene Scene) Render(renderer *sdl.Renderer) {
	scene.balls.Each(func(ball Ball) bool {
		ball.Render(renderer)
		return false
	})
}
