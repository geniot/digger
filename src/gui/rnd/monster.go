package rnd

import (
	"github.com/beefsack/go-astar"
	"github.com/geniot/digger/src/chs"
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Monster struct {
	offsetX int32
	offsetY int32
	width   int32
	height  int32

	direction Direction
	state     MonsterState

	innerOffsetX int32
	innerOffsetY int32

	spritePointer    int
	spritePointerInc int

	collisionObject *resolv.Object

	processedTimeStamp int64

	chasePath  []astar.Pather
	points     []Point
	pointsSize int32

	killerBag *Bag
	scene     *Scene
}

/**
 * INIT
 */

func NewMonster(scn *Scene) *Monster {
	mns := &Monster{}
	mns.scene = scn

	mns.innerOffsetX = 2
	mns.innerOffsetY = 2
	mns.width = 16
	mns.height = 16

	mns.collisionObject = resolv.NewObject(
		float64(mns.offsetX+mns.innerOffsetX),
		float64(mns.offsetY+mns.innerOffsetY),
		float64(mns.width),
		float64(mns.height),
		MONSTER_COLLISION_TAG)
	mns.collisionObject.Data = mns
	scn.collisionSpace.Add(mns.collisionObject)

	//same for all levels
	cellX := 14
	cellY := 0
	mns.spritePointer = 0
	mns.spritePointerInc = 1

	mns.offsetX = int32(CELLS_OFFSET + cellX*CELL_WIDTH)
	mns.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cellY*CELL_HEIGHT)
	mns.direction = RIGHT
	mns.state = MONSTER_NOBBIN

	mns.updateChasePath()
	mns.points = make([]Point, CELL_WIDTH) //reserving some extra

	return mns
}

func (monster *Monster) setPoints(x1 int32, y1 int32, x2 int32, y2 int32) {
	monster.pointsSize = int32(If(x1 == x2, math.Abs(float64(y1-y2)), math.Abs(float64(x1-x2))) + 1)
	monster.pointsSize = If(monster.pointsSize > int32(len(monster.points)), int32(len(monster.points)), monster.pointsSize)
	for i := int32(0); i < monster.pointsSize; i++ {
		if x1 == x2 {
			inc := If(y2 > y1, i, i*-1)
			monster.points[i].X = x1
			monster.points[i].Y = y1 + inc
		} else {
			inc := If(x2 > x1, i, i*-1)
			monster.points[i].X = x1 + inc
			monster.points[i].Y = y1
		}
	}
}

func (monster *Monster) setTilePoints(tile1 *chs.ChaseTile, tile2 *chs.ChaseTile) {
	monster.setPoints(
		int32(CELLS_OFFSET+tile1.X*CELL_WIDTH/2),
		int32(FIELD_OFFSET_Y+CELLS_OFFSET+tile1.Y*CELL_HEIGHT/2),
		int32(CELLS_OFFSET+tile2.X*CELL_WIDTH/2),
		int32(FIELD_OFFSET_Y+CELLS_OFFSET+tile2.Y*CELL_HEIGHT/2),
	)
}

func (monster *Monster) Step(n uint64) {
	switch monster.state {
	case MONSTER_NOBBIN, MONSTER_HOBBIN:
		if n%SPRITE_UPDATE_RATE == 0 {
			monster.spritePointer, monster.spritePointerInc = GetNextSpritePointerAndInc(
				monster.spritePointer,
				monster.spritePointerInc,
				If(monster.state == MONSTER_NOBBIN, len(monster.scene.media.monsterSpritesNobbin), len(monster.scene.media.monsterSpritesHobbin)))
		}
		if monster.scene.digger.state == DIGGER_ALIVE {
			if n%MONSTER_SPEED == 0 {
				if monster.chasePath != nil {
					dir := NONE

					for i := len(monster.chasePath) - 1; i > 0; i-- {
						thisTile := monster.chasePath[i].(*chs.ChaseTile)
						nextTile := monster.chasePath[i-1].(*chs.ChaseTile)
						monster.setTilePoints(thisTile, nextTile)
						dir = monster.getDir()
						if dir != NONE {
							if monster.canMove(dir) {
								monster.move(dir)
							}
							break
						}
					}
					if dir == NONE { //path exists, but we need to get to the first point first
						thisTile := monster.chasePath[len(monster.chasePath)-1].(*chs.ChaseTile)
						monster.setPoints(
							monster.offsetX,
							monster.offsetY,
							int32(CELLS_OFFSET+thisTile.X*CELL_WIDTH/2),
							int32(FIELD_OFFSET_Y+CELLS_OFFSET+thisTile.Y*CELL_HEIGHT/2))
						dir = monster.getDir()
						if dir != NONE && monster.canMove(dir) {
							monster.move(dir)
						}
					}
				}
			}
			if n%CHASE_PATH_UPDATE_RATE == 0 {
				monster.updateChasePath()
			}
		}
	case MONSTER_DIE:
		if monster.killerBag != nil && monster.killerBag.state == BAG_FALLING { //fall with the bag
			if monster.killerBag.offsetY > monster.offsetY {
				monster.offsetY = monster.killerBag.offsetY
				monster.collisionObject.Y = float64(monster.offsetY + monster.innerOffsetY)
				monster.collisionObject.Update()
			}
		} else {
			monster.Destroy()
		}
	}
}

func (monster *Monster) getDir() Direction {
	for k := int32(0); k < monster.pointsSize-1; k++ {
		point := monster.points[k]
		nextPoint := monster.points[k+1]
		if point.X == monster.offsetX && point.Y == monster.offsetY {
			if nextPoint.X != point.X {
				if nextPoint.X > point.X {
					return RIGHT
				} else {
					return LEFT
				}
			} else {
				if nextPoint.Y > point.Y {
					return DOWN
				} else {
					return UP
				}
			}

		}
	}
	return NONE
}

func (monster *Monster) move(dir Direction) {
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	monster.offsetX += int32(x)
	monster.offsetY += int32(y)
	monster.collisionObject.X = float64(monster.offsetX + monster.innerOffsetX)
	monster.collisionObject.Y = float64(monster.offsetY + monster.innerOffsetY)
	monster.collisionObject.Update()
}

func (monster *Monster) updateChasePath() {
	fromX := int((monster.offsetX-CELLS_OFFSET+CELL_WIDTH/2)/CELL_WIDTH) * 2
	fromY := int((monster.offsetY-CELLS_OFFSET-FIELD_OFFSET_Y+CELL_HEIGHT/2)/CELL_WIDTH) * 2

	toX := int((monster.scene.digger.offsetX-CELLS_OFFSET+CELL_WIDTH/2)/CELL_WIDTH) * 2
	toY := int((monster.scene.digger.offsetY-CELLS_OFFSET-FIELD_OFFSET_Y+CELL_HEIGHT/2)/CELL_HEIGHT) * 2

	from := monster.scene.chaseWorld.Tile(fromX, fromY)
	to := monster.scene.chaseWorld.Tile(toX, toY)

	path, _, _ := astar.Path(from, to)
	if path != nil {
		if monster.chasePath == nil {
			monster.chasePath = path
		} else {
			//we don't change monster direction, it looks too automatic
			if monster.getPathDir(path) == monster.getDir() {
				monster.chasePath = path
			}
		}

	}
}

func (monster *Monster) getPathDir(path []astar.Pather) Direction {
	if len(path) > 1 {
		lastTile := path[len(path)-1].(*chs.ChaseTile)
		nextTile := path[len(path)-2].(*chs.ChaseTile)
		if lastTile.X == nextTile.X {
			if nextTile.Y > lastTile.Y {
				return DOWN
			} else {
				return UP
			}
		} else {
			if nextTile.X > lastTile.X {
				return RIGHT
			} else {
				return LEFT
			}
		}
	} else {
		return NONE
	}
}

func (monster *Monster) getHitBox() *sdl.Rect {
	return &sdl.Rect{X: monster.offsetX + monster.innerOffsetX, Y: monster.offsetY + monster.innerOffsetY, W: monster.width, H: monster.height}
}

func (monster *Monster) Render() {
	switch monster.state {
	case MONSTER_NOBBIN:
		ctx.RendererIns.Copy(monster.scene.media.monsterSpritesNobbin[monster.spritePointer],
			nil, &sdl.Rect{X: monster.offsetX, Y: monster.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	case MONSTER_DIE:
		ctx.RendererIns.Copy(monster.scene.media.monsterSpriteNobbinDead,
			nil, &sdl.Rect{X: monster.offsetX, Y: monster.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})
	}

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(monster.getHitBox())
	}

}

func (monster *Monster) Destroy() {
	monster.scene.collisionSpace.Remove(monster.collisionObject)
	monster.scene.monsters.Remove(monster)
}

func (monster *Monster) canMove(dir Direction) bool {
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	if collision := monster.collisionObject.Check(float64(x), float64(y)); collision != nil {
		for i := 0; i < len(collision.Objects); i++ {
			if digger, ok1 := collision.Objects[i].Data.(*Digger); ok1 {
				digger.kill()
				return false
			} else if bag, ok2 := collision.Objects[i].Data.(*Bag); ok2 {
				bag.push(dir, monster)
				return false
			}
		}
	}
	return true
}

func (monster *Monster) kill() {
	monster.state = MONSTER_DIE
}
