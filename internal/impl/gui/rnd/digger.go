package rnd

import (
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
	direction Direction

	innerOffsetX int32
	innerOffsetY int32

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

	dg.innerOffsetX = 2
	dg.innerOffsetY = 2

	dg.collisionObject = resolv.NewObject(float64(dg.offsetX+dg.innerOffsetX), float64(dg.offsetY+dg.innerOffsetY), float64(dg.width), float64(dg.height), DIGGER_COLLISION_TAG)
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
			if digger.direction == RIGHT {
				if cM, _ := digger.canMoveShouldTurn(RIGHT); cM {
					digger.move(RIGHT)
				}
			} else if digger.direction == LEFT {
				digger.direction = RIGHT
			} else {
				if (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT != 0 {
					if digger.direction == UP {
						cM, sT := digger.canMoveShouldTurn(UP)
						if cM {
							digger.move(UP)
						}
						if sT {
							digger.direction = DOWN
						}
					} else if digger.direction == DOWN {
						cM, sT := digger.canMoveShouldTurn(DOWN)
						if cM {
							digger.move(DOWN)
						}
						if sT {
							digger.direction = UP
						}
					}
				} else {
					digger.direction = RIGHT
				}
			}
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_LEFT]; ok {
			if digger.direction == LEFT {
				if cM, _ := digger.canMoveShouldTurn(LEFT); cM {
					digger.move(LEFT)
				}
			} else if digger.direction == RIGHT {
				digger.direction = LEFT
			} else {
				if (FIELD_OFFSET_Y+CELLS_OFFSET+digger.offsetY)%CELL_HEIGHT != 0 {
					if digger.direction == UP {
						cM, sT := digger.canMoveShouldTurn(UP)
						if cM {
							digger.move(UP)
						}
						if sT {
							digger.direction = DOWN
						}
					} else {
						cM, sT := digger.canMoveShouldTurn(DOWN)
						if cM {
							digger.move(DOWN)
						}
						if sT {
							digger.direction = UP
						}
					}
				} else {
					digger.direction = LEFT
				}
			}
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_UP]; ok {
			if digger.direction == UP {
				if cM, _ := digger.canMoveShouldTurn(UP); cM {
					digger.move(UP)
				}
			} else if digger.direction == DOWN {
				digger.direction = UP
			} else {
				if (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH != 0 {
					if digger.direction == LEFT {
						cM, sT := digger.canMoveShouldTurn(LEFT)
						if cM {
							digger.move(LEFT)
						}
						if sT {
							digger.direction = RIGHT
						}
					} else {
						cM, sT := digger.canMoveShouldTurn(RIGHT)
						if cM {
							digger.move(RIGHT)
						}
						if sT {
							digger.direction = LEFT
						}
					}
				} else {
					digger.direction = UP
				}
			}
		} else if _, ok = ctx.PressedKeysCodesSetIns[GCW_BUTTON_DOWN]; ok {
			if digger.direction == DOWN {
				if cM, _ := digger.canMoveShouldTurn(DOWN); cM {
					digger.move(DOWN)
				}
			} else if digger.direction == UP {
				digger.direction = DOWN
			} else {
				if (CELLS_OFFSET+digger.offsetX)%CELL_WIDTH != 0 {
					if digger.direction == LEFT {
						cM, sT := digger.canMoveShouldTurn(LEFT)
						if cM {
							digger.move(LEFT)
						}
						if sT {
							digger.direction = RIGHT
						}
					} else {
						cM, sT := digger.canMoveShouldTurn(RIGHT)
						if cM {
							digger.move(RIGHT)
						}
						if sT {
							digger.direction = LEFT
						}
					}
				} else {
					digger.direction = DOWN
				}
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
	x := If(dir == RIGHT, int32(1), If(dir == LEFT, int32(-1), 0))
	y := If(dir == DOWN, int32(1), If(dir == UP, int32(-1), 0))
	digger.offsetX += x
	digger.offsetY += y
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
				em.Destroy()
			} else if bag, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				bag.push(dir)
				return false, !bag.canMove(dir)
			}
		}
	}
	return true, false
}

func (digger *Digger) getHitBox() *sdl.Rect {
	return &sdl.Rect{digger.offsetX + digger.innerOffsetX, digger.offsetY + digger.innerOffsetY, digger.width, digger.height}
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
