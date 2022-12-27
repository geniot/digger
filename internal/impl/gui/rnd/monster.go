package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
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
	spritesNobbin    []*sdl.Texture
	spritesHobbin    []*sdl.Texture

	collisionObject *resolv.Object

	processedTimeStamp int64

	scene *Scene
}

/**
 * INIT
 */

func NewMonster(scn *Scene) *Monster {
	mns := &Monster{}
	mns.scene = scn

	mns.spritesNobbin = []*sdl.Texture{
		resources.LoadTexture("cnob1.png"),
		resources.LoadTexture("cnob2.png"),
		resources.LoadTexture("cnob3.png")}

	mns.spritesHobbin = []*sdl.Texture{
		resources.LoadTexture("clhob1.png"),
		resources.LoadTexture("clhob2.png"),
		resources.LoadTexture("clhob3.png")}

	mns.innerOffsetX = 2
	mns.innerOffsetY = 2
	mns.width = 16
	mns.height = 16

	mns.collisionObject = resolv.NewObject(float64(mns.offsetX+mns.innerOffsetX), float64(mns.offsetY+mns.innerOffsetY), float64(mns.width), float64(mns.height), DIGGER_COLLISION_TAG)
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

	return mns
}

func (monster *Monster) Step(n uint64) {
	if n%SPRITE_UPDATE_RATE == 0 {
		monster.spritePointer, monster.spritePointerInc = GetNextSpritePointerAndInc(
			monster.spritePointer,
			monster.spritePointerInc,
			If(monster.state == MONSTER_NOBBIN, len(monster.spritesNobbin), len(monster.spritesHobbin)))
	}
}

func (monster *Monster) getHitBox() *sdl.Rect {
	return &sdl.Rect{monster.offsetX + monster.innerOffsetX, monster.offsetY + monster.innerOffsetY, monster.width, monster.height}
}

func (monster *Monster) Render() {
	switch monster.state {
	case MONSTER_NOBBIN:
		ctx.RendererIns.CopyEx(
			monster.spritesNobbin[monster.spritePointer],
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