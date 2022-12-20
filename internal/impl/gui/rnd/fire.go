package rnd

import (
	"github.com/geniot/digger/internal/api"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type Fire struct {
	offsetX int
	offsetY int

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture

	direction api.Direction

	scene *Scene
}

/**
 * INIT
 */

func NewFire(cX int, cY int, scn *Scene) *Fire {
	fr := &Fire{}
	fr.scene = scn
	fr.sprites = []*sdl.Texture{resources.LoadTexture("cfire1.png"), resources.LoadTexture("cfire2.png"), resources.LoadTexture("cfire3.png")}
	fr.offsetX = cX
	fr.offsetY = cY
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

}
