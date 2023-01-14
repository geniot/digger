package rnd

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/mix"
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

	spriteShakePointer int
	spriteGoldPointer  int

	startFallFromY int32
	startHold      uint64
	soundChannel   int

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

	bg.spriteShakePointer = 0
	bg.spriteGoldPointer = 0

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

	bg.soundChannel = -1

	return bg
}

/**
 * MODEL
 */

func (bag *Bag) getHitBox() *sdl.Rect {
	if bag.state == BAG_GOLD || bag.state == BAG_GOLD_FALLING {
		return &sdl.Rect{X: bag.offsetX + bag.innerOffsetX + 1, Y: bag.offsetY + bag.innerOffsetY + 6, W: bag.width - 3, H: bag.height - 6}
	} else {
		return &sdl.Rect{X: bag.offsetX + bag.innerOffsetX, Y: bag.offsetY + bag.innerOffsetY, W: bag.width, H: bag.height}
	}
}

func (bag *Bag) getFallBox() *sdl.Rect {
	return &sdl.Rect{X: bag.offsetX + bag.innerOffsetX*2, Y: bag.offsetY + bag.innerOffsetY + CELL_HEIGHT, W: bag.width - bag.innerOffsetX*2, H: bag.height - bag.innerOffsetY}
}

func (bag *Bag) Destroy() {
	bag.scene.collisionSpace.Remove(bag.collisionObject)
	bag.scene.bags.Remove(bag)
}

func (bag *Bag) Step(n uint64) {
	switch bag.state {
	case BAG_SET:
		if n%SPRITE_UPDATE_RATE == 0 {
			if bag.hasHollowSpaceUnder() {
				bag.state = BAG_HOLD
				bag.startHold = n
			}
		}
	case BAG_HOLD:
		if n-bag.startHold > HOLD_WAIT_STEPS && !bag.isOnHold() {
			bag.state = BAG_WOBBLE
			bag.soundChannel, _ = bag.scene.media.soundWobble.Play(bag.soundChannel, 0)
		}
	case BAG_PUSHED:
		bag.state = BAG_MOVING
		if (CELLS_OFFSET+bag.offsetX)%CELL_WIDTH == 0 && bag.hasHollowSpaceUnder() {
			bag.startFall()
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
						bag.startFall()
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
						bag.startFall()
					} else {
						bag.state = BAG_SET
					}
				}
			}
		}
	case BAG_WOBBLE:
		if n%(SPRITE_UPDATE_RATE*3) == 0 {
			bag.spriteShakePointer += 1
			if bag.spriteShakePointer >= len(bag.scene.media.bagSpritesShakeFrameSequence) {
				bag.spriteShakePointer = 0
				bag.startFall()
			}
		}

	case BAG_FALLING:
		if n%BAG_FALL_SPEED == 0 {
			if bag.canFall() {
				bag.fall()
			} else {
				mix.HaltChannel(bag.soundChannel)
				if bag.state == BAG_FALLING {
					if bag.offsetY-bag.startFallFromY > CELL_HEIGHT {
						bag.turnToGold(BAG_GOLD)
					} else {
						if bag.state != BAG_GOLD {
							bag.state = BAG_SET
						}
					}
				}
			}
		}
	case BAG_GOLD_FALLING:
		if n%BAG_FALL_SPEED == 0 {
			if bag.canFall() {
				bag.fall()
			} else {
				bag.state = BAG_GOLD
			}
		}
	case BAG_GOLD:
		if n%(SPRITE_UPDATE_RATE*4) == 0 {
			if bag.spriteGoldPointer < len(bag.scene.media.bagSpritesGoldFrameSequence)-1 {
				bag.spriteGoldPointer += 1
			}
		}
	}

}

func (bag *Bag) startFall() {
	bag.state = BAG_FALLING
	bag.startFallFromY = bag.offsetY
	bag.soundChannel, _ = bag.scene.media.soundFall.Play(bag.soundChannel, 0)
}

func (bag *Bag) fall() {
	bag.offsetY += 1
	bag.scene.field.drawEatUp(bag.offsetX, bag.offsetY+bag.height-4)
	bag.updateCollisionObject()
}

func (bag *Bag) turnToGold(newState BagState) {
	bag.state = newState
	bag.updateCollisionObject()
	if bag.state == BAG_GOLD && bag.scene.digger.state == DIGGER_ALIVE {
		bag.soundChannel, _ = bag.scene.media.soundBagToGold.Play(bag.soundChannel, 0)
	}
}

func (bag *Bag) updateCollisionObject() {
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
				if bg.state == BAG_GOLD { //bag can be pushed over gold
					bg.Destroy()
					return true
				} else {
					bg.push(dir)
					bag.moveAttempts += 1
					return false
				}
			} else if _, ok2 := collision.Objects[i].Data.(*Digger); ok2 {
				return false
			}
		}
	}
	return true
}

/*
*
Checks that the digger is below this bag and is holding it
*/
func (bag *Bag) isOnHold() bool {
	if collision := bag.collisionObject.Check(0, 1); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if _, ok1 := collision.Objects[i].Data.(*Digger); ok1 {
				if _, ok2 := ctx.PressedKeysCodesSetIns[GCW_BUTTON_UP]; ok2 {
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}

/**
 * VIEW
 */

func (bag *Bag) Render() {
	switch bag.state {
	case BAG_SET, BAG_PUSHED, BAG_HOLD, BAG_MOVING:
		ctx.RendererIns.Copy(bag.scene.media.bagTexture, nil, &sdl.Rect{X: bag.offsetX, Y: bag.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	case BAG_WOBBLE:
		dstRect := sdl.Rect{X: bag.offsetX, Y: bag.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT}
		ctx.RendererIns.CopyEx(bag.scene.media.bagSpritesShake[bag.scene.media.bagSpritesShakeFrameSequence[bag.spriteShakePointer]], nil, &dstRect, 0, &sdl.Point{X: CELL_WIDTH / 2, Y: CELL_HEIGHT / 2}, sdl.FLIP_NONE)
	case BAG_FALLING:
		ctx.RendererIns.Copy(bag.scene.media.bagTextureFall, nil, &sdl.Rect{X: bag.offsetX, Y: bag.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	case BAG_GOLD_FALLING, BAG_GOLD:
		ctx.RendererIns.Copy(bag.scene.media.bagSpritesGold[bag.scene.media.bagSpritesGoldFrameSequence[bag.spriteGoldPointer]],
			nil,
			&sdl.Rect{X: bag.offsetX, Y: bag.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	}

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(bag.getHitBox())
		ctx.RendererIns.SetDrawColor(0, 0, 255, 255)
		DrawRectLines(bag.getFallBox())
	}
}

func (bag *Bag) push(dir Direction) {
	switch bag.state {
	case BAG_SET:
		if dir == UP {
			bag.state = BAG_HOLD
		} else if dir == LEFT || dir == RIGHT {
			bag.pushDir = dir
			bag.state = BAG_PUSHED
		}
	case BAG_FALLING:
		if bag.scene.digger.offsetY > bag.offsetY { //if we just pushed the bag there is no kill
			bag.scene.digger.killerBag = bag
			bag.scene.digger.kill()
		}
	case BAG_GOLD:
		bag.scene.media.soundEatGold.Play(-1, 0)
		bag.Destroy()
	}
}

func (bag *Bag) canFall() bool {
	if collision := bag.collisionObject.Check(0, 1); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if em, ok1 := collision.Objects[i].Data.(*Emerald); ok1 {
				em.Destroy()
			} else if dg, ok3 := collision.Objects[i].Data.(*Digger); ok3 {
				dg.killerBag = bag
				dg.kill()
			} else if bg, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				bag.turnToGold(BAG_GOLD_FALLING)
				bg.turnToGold(BAG_GOLD)
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
