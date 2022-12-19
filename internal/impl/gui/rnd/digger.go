package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Digger struct {
	offsetX   int32
	offsetY   int32
	cellX     int
	cellY     int
	direction api.Direction

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture
	scene            *Scene
}

func NewDigger(scn *Scene) *Digger {
	spts := []*sdl.Texture{loadTexture("cldig1.png"), loadTexture("cldig2.png"), loadTexture("cldig3.png")}
	cX := 5
	cY := 5
	oX := int32(CELLS_OFFSET_X + cX*CELL_WIDTH)
	oY := int32(CELLS_OFFSET_Y + cY*CELL_HEIGHT)

	return &Digger{
		oX, oY,
		cX, cY,
		RIGHT, 0, 1, spts, scn}
}

func (digger Digger) getHitBox() (int32, int32, int32, int32) {
	return digger.offsetX + 2, digger.offsetY + 2, digger.offsetX + CELL_WIDTH - 2, digger.offsetY + CELL_HEIGHT - 2
}

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
}

func (digger *Digger) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer += digger.spritePointerInc
		if digger.spritePointer == len(digger.sprites)-1 || digger.spritePointer == 0 {
			digger.spritePointerInc = -digger.spritePointerInc
		}
	}

	if n%DIGGER_SPEED_RATE == 0 {
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_RIGHT) {
			digger.move(RIGHT, digger.moveRight, math.Mod(float64(CELLS_OFFSET_Y+digger.offsetY), CELL_HEIGHT), UP, digger.moveUp, digger.moveDown)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_LEFT) {
			digger.move(LEFT, digger.moveLeft, math.Mod(float64(CELLS_OFFSET_Y+digger.offsetY), CELL_HEIGHT), UP, digger.moveUp, digger.moveDown)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_UP) {
			digger.move(UP, digger.moveUp, math.Mod(float64(CELLS_OFFSET_X+digger.offsetX), CELL_WIDTH), LEFT, digger.moveLeft, digger.moveRight)
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_DOWN) {
			digger.move(DOWN, digger.moveDown, math.Mod(float64(CELLS_OFFSET_X+digger.offsetX), CELL_WIDTH), LEFT, digger.moveLeft, digger.moveRight)
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
	if digger.offsetX < CELLS_OFFSET_X+CELL_WIDTH*(CELLS_HORIZONTAL-1) {
		digger.direction = RIGHT
		digger.offsetX += 1
	}
}
func (digger *Digger) moveLeft() {
	if digger.offsetX > CELLS_OFFSET_X {
		digger.direction = LEFT
		digger.offsetX -= 1
	}
}
func (digger *Digger) moveUp() {
	if digger.offsetY > CELLS_OFFSET_Y {
		digger.direction = UP
		digger.offsetY -= 1
	}
}
func (digger *Digger) moveDown() {
	if digger.offsetY < CELLS_OFFSET_Y+CELL_HEIGHT*(CELLS_VERTICAL-1) {
		digger.direction = DOWN
		digger.offsetY += 1
	}
}
