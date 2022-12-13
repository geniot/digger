package rnd

import (
	"container/list"
	"geniot.com/geniot/digger/internal/api"
)

type Scene struct {
	field       *Field
	digger      *Digger
	renderables *list.List
}

func NewScene() *Scene {

	scn := &Scene{}

	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)

	scn.renderables = list.New()
	scn.renderables.PushBack(scn.field)
	scn.renderables.PushBack(scn.digger)
	//scn.renderables.PushBack(NewDebugGrid())
	//l.PushBack(NewFpsCounter())

	return scn
}

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}

func (scene *Scene) Step(n uint64) {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Step(n)
	}
}
