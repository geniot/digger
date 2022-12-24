package rnd

import (
	mapset "github.com/deckarep/golang-set/v2"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/solarlune/resolv"
	"strings"
)

type Scene struct {
	level          int
	field          *Field
	digger         *Digger
	fire           *Fire
	emeralds       mapset.Set[*Emerald]
	bags           mapset.Set[*Bag]
	collisionSpace *resolv.Space

	debugGrid  *DebugGrid
	fpsCounter *FpsCounter
}

/**
 * INIT
 */

func NewScene() *Scene {

	scn := &Scene{}
	scn.level = 1
	scn.collisionSpace = resolv.NewSpace(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT, 1, 1)

	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)
	scn.emeralds = mapset.NewSet[*Emerald]()
	scn.bags = mapset.NewSet[*Bag]()

	rows := strings.Split(strings.TrimSpace(resources.GetLevel(scn.level)), "\n")
	for y := 0; y < len(rows); y++ {
		row := rows[y]
		for x := 0; x < len(row); x++ {
			if row[x] == 'C' {
				scn.emeralds.Add(NewEmerald(x, y, scn))
			} else if row[x] == 'B' {
				scn.bags.Add(NewBag(x, y, scn))
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
	for bag := range scene.bags.Iter() {
		bag.Step(n)
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
	for emerald := range scene.emeralds.Iter() {
		emerald.Render()
	}
	for bag := range scene.bags.Iter() {
		bag.Render()
	}
	if IS_DEBUG_ON {
		scene.debugGrid.Render()
		//scene.fpsCounter.Render()
	}
}
