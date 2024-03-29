package rnd

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

type Digger struct {
	offsetX int32
	offsetY int32
	width   int32
	height  int32

	direction Direction
	state     DiggerState

	innerOffsetX int32
	innerOffsetY int32

	spritePointer    int
	spritePointerInc int

	dieCounter      int
	diePauseCounter int

	spriteGravePointer int

	collisionObject *resolv.Object

	processedTimeStamp int64

	soundChannel     int
	isDieSoundPlayed bool

	killerBag *Bag
	scene     *Scene
}

/**
 * INIT
 */

func NewDigger(scn *Scene) *Digger {
	dg := &Digger{}
	dg.scene = scn

	dg.spritePointerInc = 1

	dg.innerOffsetX = 2
	dg.innerOffsetY = 2
	dg.width = 16
	dg.height = 16

	dg.collisionObject = resolv.NewObject(float64(dg.offsetX+dg.innerOffsetX), float64(dg.offsetY+dg.innerOffsetY), float64(dg.width), float64(dg.height), DIGGER_COLLISION_TAG)
	dg.collisionObject.Data = dg
	scn.collisionSpace.Add(dg.collisionObject)

	dg.reborn()

	return dg
}

func (digger *Digger) reborn() {
	//same for all levels
	cellX := 7
	cellY := 9
	digger.dieCounter = CELL_HEIGHT / 3
	digger.diePauseCounter = CELL_HEIGHT
	digger.spriteGravePointer = 0
	digger.spritePointer = 0
	digger.soundChannel = -1
	digger.isDieSoundPlayed = false

	digger.offsetX = int32(CELLS_OFFSET + cellX*CELL_WIDTH)
	digger.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cellY*CELL_HEIGHT)
	digger.direction = RIGHT
	digger.state = DIGGER_ALIVE
	digger.collisionObject.X = float64(digger.offsetX + digger.innerOffsetX)
	digger.collisionObject.Y = float64(digger.offsetY + digger.innerOffsetY)
	digger.collisionObject.Update()
}

/**
 * MODEL
 */

func (digger *Digger) handleMove(dir1 Direction, dir2 Direction, dir3 Direction, mod int32) {
	if digger.direction == dir1 {
		if cM, _ := digger.canMoveShouldTurn(dir1); cM {
			digger.move(dir1)
		}
	} else if digger.direction == Opposite(dir1) {
		digger.direction = dir1
	} else {
		if mod != 0 {
			if digger.direction == dir2 {
				cM, sT := digger.canMoveShouldTurn(dir2)
				if cM {
					digger.move(dir2)
				}
				if sT {
					digger.direction = Opposite(dir2)
				}
			} else if digger.direction == dir3 {
				cM, sT := digger.canMoveShouldTurn(dir3)
				if cM {
					digger.move(dir3)
				}
				if sT {
					digger.direction = Opposite(dir3)
				}
			}
		} else {
			digger.direction = dir1
		}
	}
}

func (digger *Digger) Step(n uint64) {
	switch digger.state {
	case DIGGER_ALIVE:
		if n%SPRITE_UPDATE_RATE == 0 {
			digger.spritePointer, digger.spritePointerInc = GetNextSpritePointerAndInc(digger.spritePointer, digger.spritePointerInc, len(digger.scene.media.diggerSprites))
		}

		if n%DIGGER_SPEED == 0 {
			if _, ok := ctx.PressedKeysCodesSetIns[GCW_BUTTON_RIGHT]; ok {
				digger.handleMove(RIGHT, UP, DOWN, (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT)
			} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_LEFT]; ok {
				digger.handleMove(LEFT, UP, DOWN, (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT)
			} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_UP]; ok {
				digger.handleMove(UP, LEFT, RIGHT, (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH)
			} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_DOWN]; ok {
				digger.handleMove(DOWN, LEFT, RIGHT, (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH)
			}
		}
	case DIGGER_DIE:
		if digger.killerBag != nil && digger.killerBag.state == BAG_FALLING { //fall with the bag
			if digger.killerBag.offsetY > digger.offsetY {
				digger.offsetY = digger.killerBag.offsetY
				digger.collisionObject.Y = float64(digger.offsetY + digger.innerOffsetY)
				digger.collisionObject.Update()
			}
		} else {
			if !digger.isDieSoundPlayed {
				digger.soundChannel, _ = digger.scene.media.soundDie.Play(digger.soundChannel, 0)
				digger.isDieSoundPlayed = true
			}
			if n%DIGGER_DIE_SPEED == 0 { //sink at the end of the fall
				if digger.dieCounter > 0 {
					digger.offsetY += 1
					digger.dieCounter -= 1
					digger.collisionObject.Y = float64(digger.offsetY + digger.innerOffsetY)
					digger.collisionObject.Update()
				} else {
					if digger.diePauseCounter > 0 {
						digger.diePauseCounter -= 1
					} else {
						runtime.GC()
						digger.state = DIGGER_GRAVE
						digger.soundChannel, _ = digger.scene.media.soundGrave.Play(digger.soundChannel, 0)
					}
				}
			}
		}

	case DIGGER_GRAVE:
		if n%DIGGER_GRAVE_SPEED == 0 {
			if digger.spriteGravePointer < len(digger.scene.media.diggerSpritesGraveFrameSequence)-1 {
				digger.spriteGravePointer += 1
			} else {
				digger.scene.onDiggerDie()
			}
		}
	}

	if p, ok := ctx.PressedKeysCodesSetIns[GCW_BUTTON_A]; ok && p != digger.processedTimeStamp {
		digger.processedTimeStamp = p
		digger.fire()
	}
}

func (digger *Digger) fire() {
	digger.scene.fire = NewFire(digger, digger.scene)
}

func (digger *Digger) move(dir Direction) {
	digger.direction = dir
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	digger.offsetX += int32(x)
	digger.offsetY += int32(y)
	digger.collisionObject.X = float64(digger.offsetX + digger.innerOffsetX)
	digger.collisionObject.Y = float64(digger.offsetY + digger.innerOffsetY)
	digger.collisionObject.Update()
}

func (digger *Digger) canMoveShouldTurn(dir Direction) (bool, bool) {
	if !digger.scene.field.isWithinBounds(dir, digger.offsetX, digger.offsetY) {
		return false, false
	}
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	if collision := digger.collisionObject.Check(float64(x), float64(y)); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if em, ok1 := collision.Objects[i].Data.(*Emerald); ok1 {
				em.soundEat()
				em.Destroy()
			} else if bag, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				bag.push(dir, digger)
				return false, !bag.canMove(dir)
			}
		}
	}
	return true, false
}

func (digger *Digger) getHitBox() *sdl.Rect {
	return &sdl.Rect{X: digger.offsetX + digger.innerOffsetX, Y: digger.offsetY + digger.innerOffsetY, W: digger.width, H: digger.height}
}

/**
 * VIEW
 */

func (digger *Digger) Render() {
	switch digger.state {
	case DIGGER_ALIVE:
		flip := sdl.FLIP_NONE
		if digger.direction == RIGHT {
			flip = sdl.FLIP_HORIZONTAL
		}
		angle := 0.0
		if digger.direction == UP {
			angle = 90
		}
		if digger.direction == DOWN {
			angle = 270
		}
		ctx.RendererIns.CopyEx(
			digger.scene.media.diggerSprites[digger.spritePointer],
			nil,
			&sdl.Rect{X: digger.offsetX, Y: digger.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT},
			angle,
			&sdl.Point{X: CELL_WIDTH / 2, Y: CELL_HEIGHT / 2},
			flip)

		digger.eatField()
	case DIGGER_DIE:
		ctx.RendererIns.Copy(digger.scene.media.diggerDieTexture, nil, &sdl.Rect{X: digger.offsetX, Y: digger.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	case DIGGER_GRAVE:
		ctx.RendererIns.CopyEx(
			digger.scene.media.diggerSpritesGrave[digger.scene.media.diggerSpritesGraveFrameSequence[digger.spriteGravePointer]],
			nil,
			&sdl.Rect{X: digger.offsetX, Y: digger.offsetY - CELL_HEIGHT/3, W: CELL_WIDTH, H: CELL_HEIGHT},
			0,
			&sdl.Point{X: CELL_WIDTH / 2, Y: CELL_HEIGHT / 2}, sdl.FLIP_NONE)
	}

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(digger.getHitBox())
	}

}

func (digger *Digger) eatField() {
	field := digger.scene.field

	switch digger.direction {
	case RIGHT:
		for i := 0; i <= CELL_WIDTH/2; i++ {
			field.drawEatRight(digger.offsetX-int32(i), digger.offsetY)
		}
	case LEFT:
		for i := CELL_WIDTH / 2; i >= 0; i-- {
			field.drawEatLeft(digger.offsetX+int32(i), digger.offsetY)
		}
	case UP:
		for i := CELL_WIDTH / 2; i >= 0; i-- {
			field.drawEatUp(digger.offsetX, digger.offsetY+int32(i))
		}
	case DOWN:
		for i := 0; i <= CELL_WIDTH/2; i++ {
			field.drawEatDown(digger.offsetX, digger.offsetY-int32(i))
		}
	}
}

func (digger *Digger) kill() {
	if digger.state == DIGGER_ALIVE {
		digger.state = DIGGER_DIE
	}
}
