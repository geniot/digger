package rnd

import (
	"geniot.com/geniot/digger/internal/api"
	"geniot.com/geniot/digger/internal/ctx"
	. "geniot.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type Digger struct {
	width     int32
	height    int32
	offsetX   int32
	offsetY   int32
	cellX     int
	cellY     int
	direction api.Direction

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture
}

func NewDigger() *Digger {
	spts := []*sdl.Texture{loadTexture("cldig1.png"), loadTexture("cldig2.png"), loadTexture("cldig3.png")}
	cX := 14
	cY := 0
	oX := int32(CELLS_OFFSET_X + cX*CELL_WIDTH)
	oY := int32(CELLS_OFFSET_Y + cY*CELL_HEIGHT)

	return &Digger{16, 15,
		oX, oY,
		cX, cY,
		RIGHT, 0, 1, spts}
}

func (digger Digger) Render() {
	dstRect := sdl.Rect{digger.offsetX, digger.offsetY, digger.width, digger.height}
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

	ctx.RendererIns.CopyEx(digger.sprites[digger.spritePointer], nil, &dstRect, angle,
		&sdl.Point{digger.width / 2, digger.height / 2}, flip)
}

func (digger *Digger) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer += digger.spritePointerInc
		if digger.spritePointer == len(digger.sprites)-1 || digger.spritePointer == 0 {
			digger.spritePointerInc = -digger.spritePointerInc
		}
	}

	if n%DIGGER_SPEED == 0 {
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_RIGHT) {
			if digger.offsetX < CELLS_OFFSET_X+CELL_WIDTH*(CELLS_HORIZONTAL-1) {
				digger.offsetX += 1
			}
			digger.direction = RIGHT
		}
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_LEFT) {
			if digger.offsetX > CELLS_OFFSET_X {
				digger.offsetX -= 1
			}
			digger.direction = LEFT
		}
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_UP) {
			digger.offsetY -= 1
			digger.direction = UP
		}
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_DOWN) {
			digger.offsetY += 1
			digger.direction = DOWN
		}
	}

}
