package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Bag struct {
	offsetX         int32
	offsetY         int32
	width           int32
	height          int32
	texture         *sdl.Texture
	collisionObject *resolv.Object
	scene           *Scene
}

/**
 * INIT
 */

func NewBag(cX int, cY int, scn *Scene) *Bag {
	bg := &Bag{}
	bg.scene = scn
	bg.texture = resources.LoadTexture("csbag.png")

	bg.offsetX = int32(CELLS_OFFSET + cX*CELL_WIDTH)
	bg.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cY*CELL_HEIGHT)
	bg.width = 16
	bg.height = 15

	bg.collisionObject = resolv.NewObject(float64(bg.offsetX+2), float64(bg.offsetY+3), float64(bg.width), float64(bg.height), BAG_COLLISION_TAG)
	bg.collisionObject.Data = bg
	scn.collisionSpace.Add(bg.collisionObject)

	return bg
}

/**
 * MODEL
 */

func (bag *Bag) getHitBox() *sdl.Rect {
	return &sdl.Rect{bag.offsetX + 2, bag.offsetY + 3, bag.width, bag.height}
}

func (bag *Bag) Destroy() {
	bag.texture.Destroy()
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
