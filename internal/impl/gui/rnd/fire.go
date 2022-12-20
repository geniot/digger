package rnd

import (
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type Fire struct {
	offsetX int32
	offsetY int32

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture

	direction api.Direction

	scene *Scene
}

/**
 * INIT
 */

func NewFire(digger *Digger, scn *Scene) *Fire {
	fr := &Fire{}
	fr.scene = scn
	fr.spritePointer = 0
	fr.spritePointerInc = 1
	fr.sprites = []*sdl.Texture{resources.LoadTexture("cfire1.png"), resources.LoadTexture("cfire2.png"), resources.LoadTexture("cfire3.png")}
	fr.offsetX = digger.offsetX
	fr.offsetY = digger.offsetY
	fr.direction = digger.direction
	if fr.direction == RIGHT {
		fr.offsetX += CELL_WIDTH / 2
	} else if fr.direction == LEFT {
		fr.offsetX -= CELL_WIDTH / 2
	} else if fr.direction == UP {
		fr.offsetY -= CELL_WIDTH / 2
	} else if fr.direction == DOWN {
		fr.offsetY += CELL_WIDTH / 2
	}
	return fr
}

/**
 * MODEL
 */

func (fire *Fire) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		fire.spritePointer += fire.spritePointerInc
		if fire.spritePointer == len(fire.sprites)-1 || fire.spritePointer == 0 {
			fire.spritePointerInc = -fire.spritePointerInc
		}
	}
	if n%FIRE_SPEED_RATE == 0 {
		if fire.direction == UP {
			fire.offsetY -= 1
		} else if fire.direction == DOWN {
			fire.offsetY += 1
		} else if fire.direction == LEFT {
			fire.offsetX -= 1
		} else if fire.direction == RIGHT {
			fire.offsetX += 1
		}
	}
}

func (fire *Fire) Destroy() {
	for i := 0; i < len(fire.sprites); i++ {
		fire.sprites[i].Destroy()
	}
}

/**
 * VIEW
 */

func (fire Fire) Render() {
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

	ctx.RendererIns.CopyEx(fire.sprites[fire.spritePointer], nil, &dstRect, angle,
		&sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, flip)
}
