package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Bag struct {
	offsetX      int32
	offsetY      int32
	width        int32
	height       int32
	innerOffsetX int32
	innerOffsetY int32

	moveAttempts int

	spriteShakePointer    int
	spriteShakePointerInc int
	spritesShake          []*sdl.Texture

	texture         *sdl.Texture
	collisionObject *resolv.Object
	pushDir         Direction
	state           BagState
	scene           *Scene
}

/**
 * INIT
 */

func NewBag(cX int, cY int, scn *Scene) *Bag {
	bg := &Bag{}
	bg.scene = scn
	bg.texture = resources.LoadTexture("csbag.png")

	bg.spriteShakePointer = 0
	bg.spriteShakePointerInc = 1
	bg.spritesShake = []*sdl.Texture{resources.LoadTexture("clbag.png"), resources.LoadTexture("crbag.png")}

	bg.offsetX = int32(CELLS_OFFSET + cX*CELL_WIDTH)
	bg.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cY*CELL_HEIGHT)
	bg.width = 16
	bg.height = 15
	bg.innerOffsetX = 2
	bg.innerOffsetY = 3

	bg.state = BAG_SET
	bg.moveAttempts = 0

	bg.collisionObject = resolv.NewObject(float64(bg.offsetX+bg.innerOffsetX), float64(bg.offsetY+bg.innerOffsetY), float64(bg.width), float64(bg.height), BAG_COLLISION_TAG)
	bg.collisionObject.Data = bg
	scn.collisionSpace.Add(bg.collisionObject)

	return bg
}

/**
 * MODEL
 */

func (bag *Bag) getHitBox() *sdl.Rect {
	return &sdl.Rect{bag.offsetX + bag.innerOffsetX, bag.offsetY + bag.innerOffsetY, bag.width, bag.height}
}

func (bag *Bag) Destroy() {
	bag.texture.Destroy()
}

func (bag *Bag) Step(n uint64) {
	if n%BAG_PUSH_RATE_RATE == 0 {
		if bag.state == BAG_PUSHED {
			bag.state = BAG_MOVING
			if bag.canMove(bag.pushDir) {
				bag.move()
			}
		} else if bag.state == BAG_MOVING {
			if bag.canMove(bag.pushDir) {
				if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH != 0 {
					bag.move()
				} else {
					bag.state = BAG_SET
				}
			} else {
				if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH != 0 {
					if bag.moveAttempts > 10 {
						bag.moveAttempts = 0
						bag.pushDir = Opposite(bag.pushDir)
					}
				} else {
					bag.state = BAG_SET
				}
			}
		}
	}
}

func (bag *Bag) move() {
	bag.moveAttempts = 0
	bag.offsetX += If(bag.pushDir == RIGHT, int32(1), If(bag.pushDir == LEFT, int32(-1), 0))
	bag.collisionObject.X = float64(bag.offsetX + bag.innerOffsetX)
	bag.collisionObject.Update()
}

func (bag *Bag) canMove(dir Direction) bool {
	if dir == UP || dir == DOWN {
		return false
	}
	if !bag.scene.field.isWithinBounds(dir, bag.offsetX, bag.offsetY) {
		return false
	}
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	if collision := bag.collisionObject.Check(float64(x), 0); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if bg, ok1 := collision.Objects[i].Data.(*Bag); ok1 {
				bg.push(dir)
				bag.moveAttempts += 1
				return false
			} else if _, ok2 := collision.Objects[i].Data.(*Digger); ok2 {
				return false
			}
		}
	}
	return true
}

/**
 * VIEW
 */

func (bag *Bag) Render() {
	ctx.RendererIns.Copy(bag.texture, nil, &sdl.Rect{bag.offsetX, bag.offsetY, CELL_WIDTH, CELL_HEIGHT})

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(bag.getHitBox())
	}
}

func (bag *Bag) push(dir Direction) {
	bag.pushDir = dir
	bag.state = BAG_PUSHED
}
