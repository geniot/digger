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

	spriteShakePointer        int
	spritesShakeFrameSequence []int
	spritesShake              []*sdl.Texture

	spriteGoldPointer        int
	spritesGoldFrameSequence []int
	spritesGold              []*sdl.Texture

	startFallFromY int32

	texture         *sdl.Texture
	textureFall     *sdl.Texture
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
	bg.textureFall = resources.LoadTexture("cfbag.png")

	bg.spriteShakePointer = 0
	bg.spritesShake = []*sdl.Texture{resources.LoadTexture("csbag.png"), resources.LoadTexture("clbag.png"), resources.LoadTexture("crbag.png")}
	bg.spritesShakeFrameSequence = []int{0, 1, 2, 1, 2}

	bg.spriteGoldPointer = 0
	bg.spritesGold = []*sdl.Texture{resources.LoadTexture("cgold1.png"), resources.LoadTexture("cgold2.png"), resources.LoadTexture("cgold3.png")}
	bg.spritesGoldFrameSequence = []int{0, 1, 2}

	bg.offsetX = int32(CELLS_OFFSET + cX*CELL_WIDTH)
	bg.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cY*CELL_HEIGHT)
	bg.width = 16
	bg.height = 15
	bg.innerOffsetX = 2
	bg.innerOffsetY = 3

	bg.state = BAG_SET
	bg.moveAttempts = 0

	hitBox := bg.getHitBox()
	bg.collisionObject = resolv.NewObject(float64(hitBox.X), float64(hitBox.Y), float64(hitBox.W), float64(hitBox.H), BAG_COLLISION_TAG)
	bg.collisionObject.Data = bg
	scn.collisionSpace.Add(bg.collisionObject)

	return bg
}

/**
 * MODEL
 */

func (bag *Bag) getHitBox() *sdl.Rect {
	if bag.state == BAG_GOLD {
		return &sdl.Rect{bag.offsetX + bag.innerOffsetX + 1, bag.offsetY + bag.innerOffsetY + 6, bag.width - 3, bag.height - 6}
	} else {
		return &sdl.Rect{bag.offsetX + bag.innerOffsetX, bag.offsetY + bag.innerOffsetY, bag.width, bag.height}
	}
}

func (bag *Bag) getFallBox() *sdl.Rect {
	return &sdl.Rect{bag.offsetX + bag.innerOffsetX*2, bag.offsetY + bag.innerOffsetY + CELL_HEIGHT, bag.width - bag.innerOffsetX*2, bag.height - bag.innerOffsetY}
}

func (bag *Bag) Destroy() {
	bag.texture.Destroy()
	bag.textureFall.Destroy()
	for i := 0; i < len(bag.spritesShake); i++ {
		bag.spritesShake[i].Destroy()
	}
	for i := 0; i < len(bag.spritesGold); i++ {
		bag.spritesGold[i].Destroy()
	}
	bag.scene.collisionSpace.Remove(bag.collisionObject)
	bag.scene.bags.Remove(bag)
}

func (bag *Bag) Step(n uint64) {
	switch bag.state {
	case BAG_SET:
		if n%SPRITE_UPDATE_RATE == 0 {
			if bag.hasHollowSpaceUnder() {
				bag.state = BAG_SHAKING
			}
		}
	case BAG_PUSHED:
		bag.state = BAG_MOVING
		if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH == 0 && bag.hasHollowSpaceUnder() {
			bag.state = BAG_FALLING
			bag.startFallFromY = bag.offsetY
		} else {
			if bag.canMove(bag.pushDir) {
				bag.move()
			}
		}
	case BAG_MOVING:
		if n%BAG_PUSH_SPEED == 0 {
			if bag.canMove(bag.pushDir) {
				if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH != 0 {
					bag.move()
				} else {
					if bag.hasHollowSpaceUnder() {
						bag.state = BAG_FALLING
						bag.startFallFromY = bag.offsetY
					} else {
						bag.state = BAG_SET
					}
				}
			} else {
				if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH != 0 {
					if bag.moveAttempts > 10 {
						bag.moveAttempts = 0
						bag.pushDir = Opposite(bag.pushDir)
					}
				} else {
					if bag.hasHollowSpaceUnder() {
						bag.state = BAG_FALLING
						bag.startFallFromY = bag.offsetY
					} else {
						bag.state = BAG_SET
					}
				}
			}
		}
	case BAG_SHAKING:
		if n%(SPRITE_UPDATE_RATE*4) == 0 {
			bag.spriteShakePointer += 1
			if bag.spriteShakePointer >= len(bag.spritesShakeFrameSequence) {
				bag.spriteShakePointer = 0
				bag.state = BAG_FALLING
				bag.startFallFromY = bag.offsetY
			}
		}

	case BAG_FALLING:
		if n%BAG_FALL_SPEED == 0 {
			if bag.canFall() {
				bag.offsetY += 1
				bag.scene.field.drawEatUp(bag.offsetX, bag.offsetY+bag.height-4)
				bag.collisionObject.Y = float64(bag.offsetY + bag.innerOffsetY)
				bag.collisionObject.Update()
			} else {
				if bag.offsetY-bag.startFallFromY > CELL_HEIGHT {
					bag.turnToGold()
				} else {
					if bag.state != BAG_GOLD {
						bag.state = BAG_SET
					}
				}
			}
		}
	case BAG_GOLD:
		if n%(SPRITE_UPDATE_RATE*4) == 0 {
			if bag.spriteGoldPointer < len(bag.spritesGoldFrameSequence)-1 {
				bag.spriteGoldPointer += 1
			}
		}
	}

}

func (bag *Bag) turnToGold() {
	bag.state = BAG_GOLD
	hitBox := bag.getHitBox()
	bag.collisionObject.X = float64(hitBox.X)
	bag.collisionObject.Y = float64(hitBox.Y)
	bag.collisionObject.W = float64(hitBox.W)
	bag.collisionObject.H = float64(hitBox.H)
	bag.collisionObject.Update()
}

func (bag *Bag) hasHollowSpaceUnder() bool {
	fB := bag.getFallBox()
	return !bag.scene.field.isPointField(fB.X, fB.Y+fB.H/2) ||
		!bag.scene.field.isPointField(fB.X+fB.W, fB.Y+fB.H/2) ||
		!bag.scene.field.isPointField(fB.X+fB.W/2, fB.Y+fB.H)
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
	if bag.state == BAG_SHAKING {
		dstRect := sdl.Rect{bag.offsetX, bag.offsetY, CELL_WIDTH, CELL_HEIGHT}
		ctx.RendererIns.CopyEx(bag.spritesShake[bag.spritesShakeFrameSequence[bag.spriteShakePointer]], nil, &dstRect, 0, &sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, sdl.FLIP_NONE)
	} else if bag.state == BAG_GOLD {
		dstRect := sdl.Rect{bag.offsetX, bag.offsetY, CELL_WIDTH, CELL_HEIGHT}
		ctx.RendererIns.CopyEx(bag.spritesGold[bag.spritesGoldFrameSequence[bag.spriteGoldPointer]], nil, &dstRect, 0, &sdl.Point{CELL_WIDTH / 2, CELL_HEIGHT / 2}, sdl.FLIP_NONE)
	} else {
		ctx.RendererIns.Copy(If(bag.state == BAG_FALLING, bag.textureFall, bag.texture), nil, &sdl.Rect{bag.offsetX, bag.offsetY, CELL_WIDTH, CELL_HEIGHT})
	}

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(bag.getHitBox())
		ctx.RendererIns.SetDrawColor(0, 0, 255, 255)
		DrawRectLines(bag.getFallBox())
	}
}

func (bag *Bag) push(dir Direction) {
	if dir == UP {
		bag.state = BAG_HOLD
	} else if dir == LEFT || dir == RIGHT {
		bag.pushDir = dir
		bag.state = BAG_PUSHED
	}
}

func (bag *Bag) canFall() bool {
	if collision := bag.collisionObject.Check(0, 1); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if em, ok1 := collision.Objects[i].Data.(*Emerald); ok1 {
				em.Destroy()
			} else if dg, ok3 := collision.Objects[i].Data.(*Digger); ok3 {
				dg.kill()
			} else if bg, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				bag.turnToGold()
				bg.turnToGold()
				return false
			}
		}
	}
	if (bag.offsetY-FIELD_OFFSET_Y-CELLS_OFFSET)%CELL_HEIGHT == 0 {
		return bag.hasHollowSpaceUnder()
	} else {
		return true
	}
}
