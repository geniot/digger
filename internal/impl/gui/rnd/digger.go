package rnd

import (
	"geniot.com/geniot/digger/internal/api"
	"geniot.com/geniot/digger/internal/ctx"
	. "geniot.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"math"
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
	cX := 0
	cY := 0
	oX := int32(CELLS_OFFSET_X + cX*CELL_WIDTH)
	oY := int32(CELLS_OFFSET_Y + cY*CELL_HEIGHT)

	return &Digger{20, 20,
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

	if n%DIGGER_SPEED_RATE == 0 {
		if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_RIGHT) {
			if digger.direction == RIGHT {
				digger.moveRight()
			} else if digger.direction == LEFT {
				digger.direction = RIGHT
			} else {
				mod := math.Mod(float64(CELLS_OFFSET_Y+digger.offsetY), CELL_HEIGHT)
				if mod != 0 {
					if digger.direction == UP {
						digger.moveUp()
					} else {
						digger.moveDown()
					}
				} else {
					digger.direction = RIGHT
				}
			}
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_LEFT) {
			if digger.direction == LEFT {
				digger.moveLeft()
			} else if digger.direction == RIGHT {
				digger.direction = LEFT
			} else {
				mod := math.Mod(float64(CELLS_OFFSET_Y+digger.offsetY), CELL_HEIGHT)
				if mod != 0 {
					if digger.direction == UP {
						digger.moveUp()
					} else {
						digger.moveDown()
					}
				} else {
					digger.direction = LEFT
				}
			}
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_UP) {
			if digger.direction == UP {
				digger.moveUp()
			} else if digger.direction == DOWN {
				digger.direction = UP
			} else {
				mod := math.Mod(float64(CELLS_OFFSET_X+digger.offsetX), CELL_WIDTH)
				if mod != 0 {
					if digger.direction == LEFT {
						digger.moveLeft()
					} else {
						digger.moveRight()
					}
				} else {
					digger.direction = UP
				}
			}
		} else if ctx.PressedKeysCodesSetIns.Contains(GCW_BUTTON_DOWN) {
			if digger.direction == DOWN {
				digger.moveDown()
			} else if digger.direction == UP {
				digger.direction = DOWN
			} else {
				mod := math.Mod(float64(CELLS_OFFSET_X+digger.offsetX), CELL_WIDTH)
				if mod != 0 {
					if digger.direction == LEFT {
						digger.moveLeft()
					} else {
						digger.moveRight()
					}
				} else {
					digger.direction = DOWN
				}
			}
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
