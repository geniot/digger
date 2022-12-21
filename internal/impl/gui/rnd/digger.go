package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Digger struct {
	offsetX   int32
	offsetY   int32
	width     int32
	height    int32
	direction api.Direction

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture

	collisionObject *resolv.Object

	processedTimeStamp int64

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
	dg.width = 16
	dg.height = 16
	dg.direction = RIGHT
	dg.spritePointer = 0
	dg.spritePointerInc = 1

	dg.collisionObject = resolv.NewObject(float64(dg.offsetX+2), float64(dg.offsetY+2), float64(dg.width), float64(dg.height), DIGGER_COLLISION_TAG)
	dg.collisionObject.Data = dg
	scn.collisionSpace.Add(dg.collisionObject)

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
		if _, ok := ctx.PressedKeysCodesSetIns[GCW_BUTTON_RIGHT]; ok {
			digger.move(RIGHT, digger.moveRight, (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT, UP, digger.moveUp, digger.moveDown)
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_LEFT]; ok {
			digger.move(LEFT, digger.moveLeft, (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT, UP, digger.moveUp, digger.moveDown)
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_UP]; ok {
			digger.move(UP, digger.moveUp, (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH, LEFT, digger.moveLeft, digger.moveRight)
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_DOWN]; ok {
			digger.move(DOWN, digger.moveDown, (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH, LEFT, digger.moveLeft, digger.moveRight)
		}
	}

	if p, ok := ctx.PressedKeysCodesSetIns[GCW_BUTTON_A]; ok && p != digger.processedTimeStamp {
		digger.processedTimeStamp = p
		digger.fire()
	}

	//for e := digger.scene.emeralds.Front(); e != nil; e = e.Next() {
	//	if Collide(digger.getHitBox(), e.Value.(*Emerald).getHitBox()) {
	//		e.Value.(*Emerald).Destroy()
	//		digger.scene.emeralds.Remove(e)
	//	}
	//}
}

func (digger *Digger) fire() {
	digger.scene.fire = NewFire(digger, digger.scene)
}

func (digger *Digger) move(
	dir api.Direction, moveFunc api.DirectionMoveFunc, mod int32,
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
		if digger.canMove(1, 0) {
			digger.offsetX += 1
			digger.collisionObject.X = float64(digger.offsetX + 2)
		}
	}
}
func (digger *Digger) moveLeft() {
	if digger.offsetX > CELLS_OFFSET {
		digger.direction = LEFT
		if digger.canMove(-1, 0) {
			digger.offsetX -= 1
			digger.collisionObject.X = float64(digger.offsetX + 2)
		}
	}
}
func (digger *Digger) moveUp() {
	if digger.offsetY > FIELD_OFFSET_Y+CELLS_OFFSET {
		digger.direction = UP
		if digger.canMove(0, -1) {
			digger.offsetY -= 1
			digger.collisionObject.Y = float64(digger.offsetY + 2)
		}
	}
}
func (digger *Digger) moveDown() {
	if digger.offsetY < FIELD_OFFSET_Y+CELLS_OFFSET+CELL_HEIGHT*(CELLS_VERTICAL-1) {
		digger.direction = DOWN
		if digger.canMove(0, 1) {
			digger.offsetY += 1
			digger.collisionObject.Y = float64(digger.offsetY + 2)
		}
	}
}

func (digger *Digger) canMove(x float64, y float64) bool {
	if collision := digger.collisionObject.Check(x, y); collision != nil {
		if em, ok := collision.Objects[0].Data.(*Emerald); ok {
			em.Destroy()
			digger.scene.collisionSpace.Remove(collision.Objects[0])
			digger.scene.emeralds.Remove(em)
		} else if _, ok = collision.Objects[0].Data.(*Bag); ok {
			return false
		}
	}
	return true
}

func (digger *Digger) getHitBox() *sdl.Rect {
	return &sdl.Rect{digger.offsetX + 2, digger.offsetY + 2, digger.width, digger.height}
}

/**
 * VIEW
 */

func (digger *Digger) Render() {
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

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(digger.getHitBox())
	}

	ctx.RendererIns.CopyEx(digger.sprites[digger.spritePointer], nil, &dstRect, angle,
		&sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)

	digger.eatField()
}

func (digger *Digger) eatField() {
	field := digger.scene.field

	if digger.direction == RIGHT { //RIGHT
		for i := 0; i <= CELL_WIDTH/2; i++ {
			field.drawEatRight(digger.offsetX-int32(i), digger.offsetY)
		}
	} else if digger.direction == LEFT { //LEFT
		for i := CELL_WIDTH / 2; i >= 0; i-- {
			field.drawEatLeft(digger.offsetX+int32(i), digger.offsetY)
		}
	} else if digger.direction == UP { //UP
		for i := CELL_WIDTH / 2; i >= 0; i-- {
			field.drawEatUp(digger.offsetX, digger.offsetY+int32(i))
		}
	} else if digger.direction == DOWN { //DOWN
		for i := 0; i <= CELL_WIDTH/2; i++ {
			field.drawEatDown(digger.offsetX, digger.offsetY-int32(i))
		}
	}
}
