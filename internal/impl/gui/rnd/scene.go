package rnd

import (
	"container/list"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"strings"
)

type Scene struct {
	level    int
	field    *Field
	digger   *Digger
	fire     *Fire
	emeralds *list.List
	bags     *list.List

	debugGrid  *DebugGrid
	fpsCounter *FpsCounter
}

/**
 * INIT
 */

func NewScene() *Scene {

	scn := &Scene{}
	scn.level = 1
	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)
	scn.emeralds = list.New()
	scn.bags = list.New()

	rows := strings.Split(strings.TrimSpace(resources.GetLevel(scn.level)), "\n")
	for y := 0; y < len(rows); y++ {
		row := rows[y]
		for x := 0; x < len(row); x++ {
			if row[x] == 'C' {
				scn.emeralds.PushBack(NewEmerald(x, y, scn))
			} else if row[x] == 'B' {
				scn.bags.PushBack(NewBag(x, y, scn))
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

	scn.debugGrid = NewDebugGrid()
	scn.fpsCounter = NewFpsCounter()

	return scn
}

func (scene *Scene) isTunnel(ch uint8) bool {
	return ch == 'V' || ch == 'H' || ch == 'S'
}

/**
 * MODEL
 */

func (scene *Scene) Step(n uint64) {
	scene.digger.Step(n)
	if scene.fire != nil {
		scene.fire.Step(n)
	}
}

/**
 * VIEW
 */

func (scene *Scene) Render() {
	scene.field.Render()
	scene.digger.Render()
	if scene.fire != nil {
		scene.fire.Render()
	}
	for e := scene.emeralds.Front(); e != nil; e = e.Next() {
		e.Value.(*Emerald).Render()
	}
	for e := scene.bags.Front(); e != nil; e = e.Next() {
		e.Value.(*Bag).Render()
	}
	if IS_DEBUG_ON {
		scene.debugGrid.Render()
		scene.fpsCounter.Render()
	}
}
