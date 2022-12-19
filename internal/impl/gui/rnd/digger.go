package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
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

	horizontalBlob *sdl.Surface
	verticalBlob   *sdl.Surface
	endLeftBlob    *sdl.Surface
	endRightBlob   *sdl.Surface
	endUpBlob      *sdl.Surface
	endDownBlob    *sdl.Surface

	scene *Scene
}

/**
 * INIT
 */

func NewDigger(scn *Scene) *Digger {
	dg := &Digger{}
	dg.scene = scn

	dg.sprites = []*sdl.Texture{loadTexture("cldig1.png"), loadTexture("cldig2.png"), loadTexture("cldig3.png")}
	dg.horizontalBlob, _ = img.LoadRW(resources.GetResource("blob1.png"), true)
	dg.verticalBlob, _ = img.LoadRW(resources.GetResource("blob2.png"), true)
	dg.endLeftBlob, _ = img.LoadRW(resources.GetResource("blob3.png"), true)
	dg.endRightBlob, _ = img.LoadRW(resources.GetResource("blob4.png"), true)
	dg.endUpBlob, _ = img.LoadRW(resources.GetResource("blob5.png"), true)
	dg.endDownBlob, _ = img.LoadRW(resources.GetResource("blob6.png"), true)

	cellX := 5
	cellY := 5
	dg.offsetX = int32(CELLS_OFFSET_X + cellX*CELL_WIDTH)
	dg.offsetY = int32(CELLS_OFFSET_Y + cellY*CELL_HEIGHT)
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
	sourceHorizontalTunnelRect := sdl.Rect{int32(math.Mod(float64(digger.offsetX), float64(digger.horizontalBlob.W))), 0, 1, digger.horizontalBlob.H}
	sourceVerticalTunnelRect := sdl.Rect{0, int32(math.Mod(float64(digger.offsetY), float64(digger.verticalBlob.H))), digger.verticalBlob.W, 1}

	if digger.direction == RIGHT { //RIGHT

		digger.drawEat(
			digger.horizontalBlob,
			CELL_WIDTH-digger.horizontalBlob.W,
			-CELL_HEIGHT,
			digger.endRightBlob,
			CELL_WIDTH-digger.endRightBlob.W+2,
			-CELL_HEIGHT,
			&sourceHorizontalTunnelRect)

	} else if digger.direction == LEFT { //LEFT

		digger.drawEat(
			digger.horizontalBlob,
			digger.horizontalBlob.W,
			-CELL_HEIGHT,
			digger.endLeftBlob,
			-2,
			-CELL_HEIGHT,
			&sourceHorizontalTunnelRect)

	} else if digger.direction == UP { //UP

		digger.drawEat(
			digger.verticalBlob,
			0,
			-CELL_HEIGHT+digger.verticalBlob.H,
			digger.endUpBlob,
			0,
			-CELL_HEIGHT-digger.endUpBlob.H+2,
			&sourceVerticalTunnelRect)

	} else if digger.direction == DOWN { //DOWN

		digger.drawEat(
			digger.verticalBlob,
			0,
			-digger.verticalBlob.H,
			digger.endDownBlob,
			0,
			-3,
			&sourceVerticalTunnelRect)

	}
}

func (digger *Digger) drawEat(tunnelSurface *sdl.Surface, x1 int32, y1 int32, endSurface *sdl.Surface, x2 int32, y2 int32, sourceRect *sdl.Rect) {
	targetTunnelRect := sdl.Rect{digger.offsetX + x1, digger.offsetY + y1, CELL_WIDTH, CELL_HEIGHT}
	tunnelSurface.Blit(sourceRect, digger.scene.field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{digger.offsetX + x2, digger.offsetY + y2, CELL_WIDTH, CELL_HEIGHT}
	endSurface.Blit(nil, digger.scene.field.background, &targetEndRect)
}
