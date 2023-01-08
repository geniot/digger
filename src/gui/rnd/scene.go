package rnd

import (
	mapset "github.com/deckarep/golang-set/v2"
	. "github.com/geniot/digger/src/chs"
	. "github.com/geniot/digger/src/glb"
	"github.com/geniot/digger/src/res"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/mix"
	"strconv"
	"strings"
	"time"
)

type Scene struct {
	level    int
	field    *Field
	digger   *Digger
	fire     *Fire
	emeralds mapset.Set[*Emerald]
	bags     mapset.Set[*Bag]
	monsters mapset.Set[*Monster]

	diggerTune        *mix.Chunk
	eatEmerald        [8]*mix.Chunk
	eatEmeraldPointer int
	lastEat           int64
	eatGold           *mix.Chunk

	collisionSpace *resolv.Space
	chaseWorld     *ChaseWorld

	debugGrid  *DebugGrid
	fpsCounter *DebugFpsCounter
}

/**
 * INIT
 */

func NewScene() *Scene {

	scn := &Scene{}
	scn.level = 1
	scn.collisionSpace = resolv.NewSpace(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT, 1, 1)
	scn.chaseWorld = &ChaseWorld{}

	for y := 0; y < CELLS_VERTICAL*2-1; y++ {
		for x := 0; x < CELLS_HORIZONTAL*2-1; x++ {
			if y%2 == 0 && x%2 == 0 {
				scn.chaseWorld.SetTile(&ChaseTile{Kind: KindTunnel, X: x, Y: y, W: *scn.chaseWorld}, x, y)
			} else {
				scn.chaseWorld.SetTile(&ChaseTile{Kind: KindField, X: x, Y: y, W: *scn.chaseWorld}, x, y)
			}
		}
	}

	scn.field = NewField(scn)
	scn.digger = NewDigger(scn)
	scn.emeralds = mapset.NewSet[*Emerald]()
	scn.bags = mapset.NewSet[*Bag]()
	scn.monsters = mapset.NewSet[*Monster]()

	scn.monsters.Add(NewMonster(scn))

	rows := strings.Split(strings.TrimSpace(res.GetLevel(scn.level)), "\n")
	for y := 0; y < len(rows); y++ {
		row := strings.TrimSuffix(rows[y], "\n")
		row = strings.TrimSuffix(rows[y], "\r")
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

	for i := 0; i <= 7; i++ {
		scn.eatEmerald[i], _ = mix.LoadWAVRW(res.GetAudio("emerald"+strconv.FormatInt(int64(i), 10)+".wav"), true)
	}
	scn.eatEmeraldPointer = 0
	scn.lastEat = time.Now().UnixMilli()

	scn.diggerTune, _ = mix.LoadWAVRW(res.GetAudio("digger.wav"), true)
	scn.eatGold, _ = mix.LoadWAVRW(res.GetAudio("gold.wav"), true)

	//scn.diggerTune.Play(-1, 10)

	scn.debugGrid = NewDebugGrid(scn)
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
	for monster := range scene.monsters.Iter() {
		monster.Step(n)
	}
}

/**
 * VIEW
 */

func (scene *Scene) Render() {
	scene.field.Render()
	if scene.fire != nil {
		scene.fire.Render()
	}
	for emerald := range scene.emeralds.Iter() {
		emerald.Render()
	}
	for bag := range scene.bags.Iter() {
		bag.Render()
	}
	for monster := range scene.monsters.Iter() {
		monster.Render()
	}
	scene.digger.Render()

	if IS_DEBUG_ON {
		scene.debugGrid.Render()
		//
	}

	scene.fpsCounter.Render()
}

func (scene *Scene) soundEat() {
	delta := time.Now().UnixMilli() - scene.lastEat
	if delta < EM_SOUND_DELTA_MS {
		scene.eatEmeraldPointer += 1
		if scene.eatEmeraldPointer >= len(scene.eatEmerald) {
			scene.eatEmeraldPointer = 0
		}
	} else {
		scene.eatEmeraldPointer = 0
	}
	scene.lastEat = time.Now().UnixMilli()
	scene.eatEmerald[scene.eatEmeraldPointer].Play(1, 0)
}
