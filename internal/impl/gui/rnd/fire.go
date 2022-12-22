package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
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

	direction  api.Direction
	isMoving   bool
	isFinished bool

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
	fr.sprites = []*sdl.Texture{resources.LoadTexture("cfire1.png"), resources.LoadTexture("cfire2.png"), resources.LoadTexture("cfire3.png")}
	fr.spriteExplPointer = 0
	fr.spriteExplPointerInc = 1
	fr.spritesExpl = []*sdl.Texture{resources.LoadTexture("cexp1.png"), resources.LoadTexture("cexp2.png"), resources.LoadTexture("cexp3.png")}

	fr.offsetX = digger.offsetX
	fr.offsetY = digger.offsetY
	fr.innerOffsetX = 6
	fr.innerOffsetY = 6
	fr.direction = digger.direction

	fr.isMoving = true
	fr.isFinished = false

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
	return &sdl.Rect{fire.offsetX + fire.innerOffsetX, fire.offsetY + fire.innerOffsetY, fire.width, fire.height}
}

/**
 * MODEL
 */

func (fire *Fire) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		if fire.isMoving {
			fire.spritePointer += fire.spritePointerInc
			if fire.spritePointer == len(fire.sprites)-1 || fire.spritePointer == 0 {
				fire.spritePointerInc = -fire.spritePointerInc
			}
		} else {
			fire.spriteExplPointer += fire.spriteExplPointerInc
			if fire.spriteExplPointer == len(fire.spritesExpl) {
				fire.isFinished = true
			}
		}
	}
	if n%FIRE_SPEED_RATE == 0 && fire.isMoving {
		if fire.direction == UP {
			if fire.canMove(UP) {
				fire.offsetY -= 1
				fire.collisionObject.Y = float64(fire.offsetY + fire.innerOffsetY)
			} else {
				fire.isMoving = false
			}
		} else if fire.direction == DOWN {
			if fire.canMove(DOWN) {
				fire.offsetY += 1
				fire.collisionObject.Y = float64(fire.offsetY + fire.innerOffsetY)
			} else {
				fire.isMoving = false
			}
		} else if fire.direction == LEFT {
			if fire.canMove(LEFT) {
				fire.offsetX -= 1
				fire.collisionObject.X = float64(fire.offsetX + fire.innerOffsetX)
			} else {
				fire.isMoving = false
			}
		} else if fire.direction == RIGHT {
			if fire.canMove(RIGHT) {
				fire.offsetX += 1
				fire.collisionObject.X = float64(fire.offsetX + fire.innerOffsetX)
			} else {
				fire.isMoving = false
			}
		}
		if fire.scene.field.collide(fire.getHitBox(), fire.direction) {
			fire.isMoving = false
		}
	}

	if fire.isFinished {
		fire.Destroy()
	}
}

func (fire *Fire) canMove(dir api.Direction) bool {
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	if collision := fire.collisionObject.Check(float64(x), float64(y)); collision != nil {
		if _, ok := collision.Objects[0].Data.(*Emerald); ok {
			return false
		} else if _, ok = collision.Objects[0].Data.(*Bag); ok {
			return false
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

	ctx.RendererIns.CopyEx(If(fire.isMoving, fire.sprites[fire.spritePointer], fire.spritesExpl[fire.spriteExplPointer]), nil, &dstRect, angle,
		&sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(fire.getHitBox())
	}
}
