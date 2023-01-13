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

	chasePath []astar.Pather

	scene *Scene
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

	return mns
}

func (monster *Monster) getPoints(x1 int32, y1 int32, x2 int32, y2 int32) []Point {
	size := int32(If(x1 == x2, math.Abs(float64(y1-y2)), math.Abs(float64(x1-x2))) + 1)
	points := make([]Point, size)
	for i := int32(0); i < size; i++ {
		point := Point{}
		if x1 == x2 {
			inc := If(y2 > y1, i, i*-1)
			point.X = x1
			point.Y = y1 + inc
		} else {
			inc := If(x2 > x1, i, i*-1)
			point.X = x1 + inc
			point.Y = y1
		}
		points[i] = point
	}
	return points
}

func (monster *Monster) getTilePoints(tile1 *chs.ChaseTile, tile2 *chs.ChaseTile) []Point {
	return monster.getPoints(
		int32(CELLS_OFFSET+tile1.X*CELL_WIDTH/2),
		int32(FIELD_OFFSET_Y+CELLS_OFFSET+tile1.Y*CELL_HEIGHT/2),
		int32(CELLS_OFFSET+tile2.X*CELL_WIDTH/2),
		int32(FIELD_OFFSET_Y+CELLS_OFFSET+tile2.Y*CELL_HEIGHT/2),
	)
}

func (monster *Monster) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		monster.spritePointer, monster.spritePointerInc = GetNextSpritePointerAndInc(
			monster.spritePointer,
			monster.spritePointerInc,
			If(monster.state == MONSTER_NOBBIN, len(monster.scene.media.monsterSpritesNobbin), len(monster.scene.media.monsterSpritesHobbin)))
	}
	if n%DIGGER_SPEED == 0 {
		if monster.chasePath != nil {
			isOffsetUpdated := false

			for i := len(monster.chasePath) - 1; i > 0; i-- {
				thisTile := monster.chasePath[i].(*chs.ChaseTile)
				nextTile := monster.chasePath[i-1].(*chs.ChaseTile)
				points := monster.getTilePoints(thisTile, nextTile)
				isOffsetUpdated = monster.move(points)
				if isOffsetUpdated {
					break
				}
			}
			if !isOffsetUpdated { //path exists, but we need to get to the first point first
				thisTile := monster.chasePath[len(monster.chasePath)-1].(*chs.ChaseTile)
				points := monster.getPoints(
					monster.offsetX,
					monster.offsetY,
					int32(CELLS_OFFSET+thisTile.X*CELL_WIDTH/2),
					int32(FIELD_OFFSET_Y+CELLS_OFFSET+thisTile.Y*CELL_HEIGHT/2))
				monster.move(points)
			}
		}
	}
	if n%CHASE_PATH_UPDATE_RATE == 0 {
		monster.updateChasePath()
	}
}

func (monster *Monster) move(points []Point) bool {
	for k := 0; k < len(points)-1; k++ {
		point := points[k]
		nextPoint := points[k+1]
		if point.X == monster.offsetX && point.Y == monster.offsetY {
			monster.offsetX = nextPoint.X
			monster.offsetY = nextPoint.Y
			monster.collisionObject.X = float64(monster.offsetX + monster.innerOffsetX)
			monster.collisionObject.Y = float64(monster.offsetY + monster.innerOffsetY)
			monster.collisionObject.Update()
			return true
		}
	}
	return false
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
		monster.chasePath = path
	}
}

func (monster *Monster) getHitBox() *sdl.Rect {
	return &sdl.Rect{X: monster.offsetX + monster.innerOffsetX, Y: monster.offsetY + monster.innerOffsetY, W: monster.width, H: monster.height}
}

func (monster *Monster) Render() {
	switch monster.state {
	case MONSTER_NOBBIN:
		ctx.RendererIns.CopyEx(
			monster.scene.media.monsterSpritesNobbin[monster.spritePointer],
			nil,
			&sdl.Rect{X: monster.offsetX, Y: monster.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT},
			0.0,
			&sdl.Point{X: CELL_WIDTH / 2, Y: CELL_HEIGHT / 2},
			sdl.FLIP_NONE)
	}

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(monster.getHitBox())
	}

}
