package rnd

import (
	"container/list"
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
)

type Scene struct {
	level       int
	world       glb.World
	field       *Field
	digger      *Digger
	renderables *list.List
}

/**
 * INIT
 */

func NewScene() *Scene {

	scn := &Scene{}
	scn.level = 1
	scn.world = glb.ParseWorld(resources.GetLevel(scn.level))
	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)

	scn.renderables = list.New()
	scn.renderables.PushBack(scn.field)
	scn.renderables.PushBack(scn.digger)
	for i := 0; i < 10; i++ {
		scn.renderables.PushBack(NewEmerald(scn))
	}
	//scn.renderables.PushBack(NewDebugGrid())
	//l.PushBack(NewFpsCounter())

	return scn
}

/**
 * MODEL
 */

func (scene *Scene) Step(n uint64) {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Step(n)
		if _, ok := e.Value.(*Emerald); ok {
			if collide(scene.digger, e.Value.(*Emerald)) {
				e.Value.(*Emerald).Destroy()
				scene.renderables.Remove(e)
			}
		}
	}
}

func collide(digger *Digger, emerald *Emerald) bool {
	x1, y1, x2, y2 := digger.getHitBox()
	x3, y3, x4, y4 := emerald.getHitBox()
	// If one rectangle is on left side of other
	if x1 > x4 || x3 > x2 {
		return false
	}
	// If one rectangle is above other
	if y2 < y3 || y4 < y1 {
		return false
	}
	return true
}

/**
 * VIEW
 */

func (scene *Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}
