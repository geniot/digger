package rnd

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Fire struct {
	offsetX      int32
	offsetY      int32
	width        int32
	height       int32
	innerOffsetX int32
	innerOffsetY int32

	spritePointer    int
	spritePointerInc int

	spriteExplosionPointer    int
	spriteExplosionPointerInc int

	direction Direction
	state     FireState

	collisionObject *resolv.Object

	soundChannel int

	scene *Scene
}

/**
 * INIT
 */

func NewFire(digger *Digger, scn *Scene) *Fire {
	fr := &Fire{}
	fr.scene = scn

	fr.width = 8
	fr.height = 8

	fr.spritePointer = 0
	fr.spritePointerInc = 1

	fr.spriteExplosionPointer = 0
	fr.spriteExplosionPointerInc = 1

	fr.offsetX = digger.offsetX
	fr.offsetY = digger.offsetY
	fr.innerOffsetX = 6
	fr.innerOffsetY = 6
	fr.direction = digger.direction

	fr.state = FIRE_MOVING

	if fr.direction == RIGHT {
		fr.offsetX += CELL_WIDTH / 2
	} else if fr.direction == LEFT {
		fr.offsetX -= CELL_WIDTH / 2
	} else if fr.direction == UP {
		fr.offsetY -= CELL_WIDTH / 2
	} else if fr.direction == DOWN {
		fr.offsetY += CELL_WIDTH / 2
	}

	fr.collisionObject = resolv.NewObject(float64(fr.offsetX+fr.innerOffsetX), float64(fr.offsetY+fr.innerOffsetY), float64(fr.width), float64(fr.height), FIRE_COLLISION_TAG)
	fr.collisionObject.Data = fr
	scn.collisionSpace.Add(fr.collisionObject)

	fr.soundChannel, _ = scn.media.soundFire.Play(-1, 0)

	return fr
}

func (fire *Fire) getHitBox() *sdl.Rect {
	return &sdl.Rect{X: fire.offsetX + fire.innerOffsetX, Y: fire.offsetY + fire.innerOffsetY, W: fire.width, H: fire.height}
}

/**
 * MODEL
 */

func (fire *Fire) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		if fire.state == FIRE_MOVING {
			fire.spritePointer, fire.spritePointerInc = GetNextSpritePointerAndInc(fire.spritePointer, fire.spritePointerInc, len(fire.scene.media.fireSprites))
		} else {
			fire.spriteExplosionPointer += fire.spriteExplosionPointerInc
			if fire.spriteExplosionPointer == len(fire.scene.media.fireSpritesExplosion) {
				fire.state = FIRE_FINISHED
			}
		}
	}
	if n%FIRE_SPEED == 0 && fire.state == FIRE_MOVING {
		if fire.direction == UP {
			if fire.canMove(UP) {
				fire.offsetY -= 1
				fire.collisionObject.Y = float64(fire.offsetY + fire.innerOffsetY)
				fire.collisionObject.Update()
			} else {
				fire.state = FIRE_STOPPED
			}
		} else if fire.direction == DOWN {
			if fire.canMove(DOWN) {
				fire.offsetY += 1
				fire.collisionObject.Y = float64(fire.offsetY + fire.innerOffsetY)
				fire.collisionObject.Update()
			} else {
				fire.state = FIRE_STOPPED
			}
		} else if fire.direction == LEFT {
			if fire.canMove(LEFT) {
				fire.offsetX -= 1
				fire.collisionObject.X = float64(fire.offsetX + fire.innerOffsetX)
				fire.collisionObject.Update()
			} else {
				fire.state = FIRE_STOPPED
			}
		} else if fire.direction == RIGHT {
			if fire.canMove(RIGHT) {
				fire.offsetX += 1
				fire.collisionObject.X = float64(fire.offsetX + fire.innerOffsetX)
				fire.collisionObject.Update()
			} else {
				fire.state = FIRE_STOPPED
			}
		}
		if fire.collidesWithField() {
			fire.state = FIRE_STOPPED
		}
	}

	if fire.state == FIRE_FINISHED {
		fire.Destroy()
	}
}

func (fire *Fire) collidesWithField() bool {
	fld := fire.scene.field

	if fire.direction == UP {
		if fld.isPointField(fire.offsetX+fire.innerOffsetX+fire.width/2, fire.offsetY+fire.innerOffsetY) {
			return true
		}
	} else if fire.direction == DOWN {
		if fld.isPointField(fire.offsetX+fire.innerOffsetX+fire.width/2, fire.offsetY+fire.innerOffsetY+fire.height) {
			return true
		}
	} else if fire.direction == LEFT {
		if fld.isPointField(fire.offsetX+fire.innerOffsetX, fire.offsetY+fire.innerOffsetY+fire.height/2) {
			return true
		}
	} else if fire.direction == RIGHT {
		if fld.isPointField(fire.offsetX+fire.innerOffsetX+fire.width, fire.offsetY+fire.innerOffsetY+fire.height/2) {
			return true
		}
	}
	return false
}

func (fire *Fire) canMove(dir Direction) bool {
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	if collision := fire.collisionObject.Check(float64(x), float64(y)); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if _, ok1 := collision.Objects[i].Data.(*Emerald); ok1 {
				return false
			} else if _, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				return false
			} else if monster, ok3 := collision.Objects[i].Data.(*Monster); ok3 {
				fire.offsetX = monster.offsetX
				fire.offsetY = monster.offsetY
				monster.Destroy()
				return false
			}
		}
	}
	return true
}

func (fire *Fire) Destroy() {
	fire.scene.fire = nil
	fire.scene.media.soundExplode.Play(fire.soundChannel, 0)
	fire.scene.collisionSpace.Remove(fire.collisionObject)
}

/**
 * VIEW
 */

func (fire *Fire) Render() {
	dstRect := sdl.Rect{fire.offsetX, fire.offsetY, CELL_WIDTH, CELL_HEIGHT}
	flip := sdl.FLIP_NONE
	if fire.direction == RIGHT {
		flip = sdl.FLIP_HORIZONTAL
	}
	angle := 0.0
	if fire.direction == UP {
		angle = 90
	}
	if fire.direction == DOWN {
		angle = 270
	}

	ctx.RendererIns.CopyEx(If(fire.state == FIRE_MOVING, fire.scene.media.fireSprites[fire.spritePointer], fire.scene.media.fireSpritesExplosion[fire.spriteExplosionPointer]),
		nil, &dstRect, angle,
		&sdl.Point{X: CELL_WIDTH / 2, Y: CELL_HEIGHT / 2}, flip)

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(fire.getHitBox())
	}
}
