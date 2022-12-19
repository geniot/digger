package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Digger struct {
	offsetX   int32
	offsetY   int32
	direction api.Direction

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture

	scene *Scene
}

/**
 * INIT
 */

func NewDigger(scn *Scene) *Digger {
	dg := &Digger{}
	dg.scene = scn

	dg.sprites = []*sdl.Texture{resources.LoadTexture("cldig1.png"), resources.LoadTexture("cldig2.png"), resources.LoadTexture("cldig3.png")}

	//same for all levels
	cellX := 7
	cellY := 9

	dg.offsetX = int32(CELLS_OFFSET + cellX*CELL_WIDTH)
	dg.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cellY*CELL_HEIGHT)
	dg.direction = RIGHT
	dg.spritePointer = 0
	dg.spritePointerInc = 1

	return dg
}

/**
 * MODEL
 */

func (digger *Digger) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer += digger.spritePointerInc
		if digger.spritePointer == len(digger.sprites)-1 || digger.spritePointer == 0 {
			digger.spritePointerInc = -digger.spritePointerInc
		}
	}

	if n%DIGGER_SPEED_RATE == 0 {
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_RIGHT) {
			digger.move(RIGHT, digger.moveRight, math.Mod(float64(FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY), CELL_HEIGHT), UP, digger.moveUp, digger.moveDown)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_LEFT) {
			digger.move(LEFT, digger.moveLeft, math.Mod(float64(FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY), CELL_HEIGHT), UP, digger.moveUp, digger.moveDown)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_UP) {
			digger.move(UP, digger.moveUp, math.Mod(float64(CELLS_OFFSET+digger.offsetX), CELL_WIDTH), LEFT, digger.moveLeft, digger.moveRight)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_DOWN) {
			digger.move(DOWN, digger.moveDown, math.Mod(float64(CELLS_OFFSET+digger.offsetX), CELL_WIDTH), LEFT, digger.moveLeft, digger.moveRight)
		}
	}
}

func (digger *Digger) move(
	dir api.Direction, moveFunc api.DirectionMoveFunc, mod float64,
	perpendicularDir api.Direction, perpendicularMoveFunc1 api.DirectionMoveFunc, perpendicularMoveFunc2 api.DirectionMoveFunc) {
	if digger.direction == dir {
		moveFunc()
	} else if digger.direction == Opposite(dir) {
		digger.direction = dir
	} else {
		if mod != 0 {
			if digger.direction == perpendicularDir {
				perpendicularMoveFunc1()
			} else {
				perpendicularMoveFunc2()
			}
		} else {
			digger.direction = dir
		}
	}
}

func (digger *Digger) moveRight() {
	if digger.offsetX < CELLS_OFFSET+CELL_WIDTH*(CELLS_HORIZONTAL-1) {
		digger.direction = RIGHT
		digger.offsetX += 1
	}
}
func (digger *Digger) moveLeft() {
	if digger.offsetX > CELLS_OFFSET {
		digger.direction = LEFT
		digger.offsetX -= 1
	}
}
func (digger *Digger) moveUp() {
	if digger.offsetY > FIELD_OFFSET_Y+CELLS_OFFSET {
		digger.direction = UP
		digger.offsetY -= 1
	}
}
func (digger *Digger) moveDown() {
	if digger.offsetY < FIELD_OFFSET_Y+CELLS_OFFSET+CELL_HEIGHT*(CELLS_VERTICAL-1) {
		digger.direction = DOWN
		digger.offsetY += 1
	}
}

func (digger Digger) getHitBox() (int32, int32, int32, int32) {
	return digger.offsetX + 2, digger.offsetY + 2, digger.offsetX + CELL_WIDTH - 2, digger.offsetY + CELL_HEIGHT - 2
}

/**
 * VIEW
 */

func (digger Digger) Render() {
	dstRect := sdl.Rect{digger.offsetX, digger.offsetY, CELL_WIDTH, CELL_HEIGHT}
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

	//draw hit box for debug purposes
	//x1, y1, x2, y2 := digger.getHitBox()
	//ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
	//ctx.RendererIns.DrawRect(&sdl.Rect{x1, y1, x2 - x1, y2 - y1})

	ctx.RendererIns.CopyEx(digger.sprites[digger.spritePointer], nil, &dstRect, angle,
		&sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)

	digger.eatField()
}

func (digger *Digger) eatField() {
	field := digger.scene.field

	if digger.direction == RIGHT { //RIGHT
		field.drawEatRight(digger.offsetX, digger.offsetY)
	} else if digger.direction == LEFT { //LEFT
		field.drawEatLeft(digger.offsetX, digger.offsetY)
	} else if digger.direction == UP { //UP
		field.drawEatUp(digger.offsetX, digger.offsetY)
	} else if digger.direction == DOWN { //DOWN
		field.drawEatDown(digger.offsetX, digger.offsetY)
	}
}
