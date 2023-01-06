package rnd

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/geniot/digger/src/res"
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
	sprites          []*sdl.Texture

	spriteExplPointer    int
	spriteExplPointerInc int
	spritesExpl          []*sdl.Texture

	direction Direction
	state     FireState

	collisionObject *resolv.Object

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
	fr.sprites = []*sdl.Texture{res.LoadTexture("cfire1.png"), res.LoadTexture("cfire2.png"), res.LoadTexture("cfire3.png")}
	fr.spriteExplPointer = 0
	fr.spriteExplPointerInc = 1
	fr.spritesExpl = []*sdl.Texture{res.LoadTexture("cexp1.png"), res.LoadTexture("cexp2.png"), res.LoadTexture("cexp3.png")}

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
			fire.spritePointer, fire.spritePointerInc = GetNextSpritePointerAndInc(fire.spritePointer, fire.spritePointerInc, len(fire.sprites))
		} else {
			fire.spriteExplPointer += fire.spriteExplPointerInc
			if fire.spriteExplPointer == len(fire.spritesExpl) {
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
			if _, ok := collision.Objects[i].Data.(*Emerald); ok {
				return false
			} else if _, ok = collision.Objects[i].Data.(*Bag); ok {
				return false
			}
		}
	}
	return true
}

func (fire *Fire) Destroy() {
	for i := 0; i < len(fire.sprites); i++ {
		fire.sprites[i].Destroy()
	}
	for i := 0; i < len(fire.spritesExpl); i++ {
		fire.spritesExpl[i].Destroy()
	}
	fire.scene.fire = nil
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

	ctx.RendererIns.CopyEx(If(fire.state == FIRE_MOVING, fire.sprites[fire.spritePointer], fire.spritesExpl[fire.spriteExplPointer]), nil, &dstRect, angle,
		&sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(fire.getHitBox())
	}
}
