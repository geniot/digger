package rnd

import (
	"container/list"
	"geniot.com/geniot/digger/internal/api"
)

type Scene struct {
	renderables *list.List
}

func (scene *Scene) Step(n uint64) {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Step(n)
	}
}

func NewScene() *Scene {

	l := list.New()
	l.PushBack(NewFpsCounter())
	l.PushBack(NewDigger())

	return &Scene{l}
}

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}
