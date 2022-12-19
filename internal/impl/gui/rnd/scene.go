package rnd

import (
	"container/list"
	"github.com/geniot/digger/internal/api"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"strings"
)

type Scene struct {
	level       int
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
	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)

	scn.renderables = list.New()
	scn.renderables.PushBack(scn.field)
	scn.renderables.PushBack(scn.digger)

	rows := strings.Split(strings.TrimSpace(resources.GetLevel(scn.level)), "\n")
	for y := 0; y < len(rows); y++ {
		row := rows[y]
		for x := 0; x < len(row); x++ {
			if row[x] == 'C' {
				scn.renderables.PushBack(NewEmerald(x, y, scn))
			} else if row[x] == 'B' {
				scn.renderables.PushBack(NewBag(x, y, scn))
			} else if row[x] == 'S' {
				isUpCont := If(y > 0 && scn.isTunnel(rows[y-1][x]), true, false)
				isDownCont := If(y < CELLS_VERTICAL-1 && scn.isTunnel(rows[y+1][x]), true, false)
				isRightCont := If(x < CELLS_HORIZONTAL-1 && scn.isTunnel(row[x+1]), true, false)
				isLeftCont := If(x > 0 && scn.isTunnel(row[x-1]), true, false)
				scn.field.eatVertical(x, y, isUpCont, isDownCont)
				scn.field.eatHorizontal(x, y, isRightCont, isLeftCont)
			} else if row[x] == 'V' {
				isUpCont := If(y > 0 && scn.isTunnel(rows[y-1][x]), true, false)
				isDownCont := If(y < CELLS_VERTICAL-1 && scn.isTunnel(rows[y+1][x]), true, false)
				scn.field.eatVertical(x, y, isUpCont, isDownCont)
			} else if row[x] == 'H' {
				isRightCont := If(x < CELLS_HORIZONTAL-1 && scn.isTunnel(row[x+1]), true, false)
				isLeftCont := If(x > 0 && scn.isTunnel(row[x-1]), true, false)
				scn.field.eatHorizontal(x, y, isRightCont, isLeftCont)
			}
		}
	}

	//scn.renderables.PushBack(NewDebugGrid())
	//l.PushBack(NewFpsCounter())

	return scn
}

func (scene *Scene) isTunnel(ch uint8) bool {
	return ch == 'V' || ch == 'H' || ch == 'S'
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
