package rnd

import (
	"container/list"
	"github.com/geniot/digger/internal/api"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/sdl"
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

	if IS_DEBUG_ON {
		scn.renderables.PushBack(NewDebugGrid())
		scn.renderables.PushBack(NewFpsCounter())
	}

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
			if collide(scene.digger.getHitBox(), e.Value.(*Emerald).getHitBox()) {
				e.Value.(*Emerald).Destroy()
				scene.renderables.Remove(e)
			}
		} else if _, ok = e.Value.(*Fire); ok {
			if e.Value.(*Fire).isFinished {
				e.Value.(*Fire).Destroy()
				scene.renderables.Remove(e)
			}
		}
	}
}

func collide(rect1 *sdl.Rect, rect2 *sdl.Rect) bool {
	x1 := rect1.X
	y1 := rect1.Y
	x2 := x1 + rect1.W
	y2 := y1 + rect1.H
	x3 := rect2.X
	y3 := rect2.Y
	x4 := x3 + rect2.W
	y4 := y3 + rect2.H
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
