package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type Bag struct {
	cellX   int
	cellY   int
	texture *sdl.Texture
	scene   *Scene
}

/**
 * INIT
 */

func NewBag(cX int, cY int, scn *Scene) *Bag {
	em := &Bag{}
	em.scene = scn
	em.texture = resources.LoadTexture("csbag.png")
	em.cellX = cX
	em.cellY = cY
	return em
}

/**
 * MODEL
 */

func (bag *Bag) Step(n uint64) {
}

func (bag *Bag) Destroy() {
	bag.texture.Destroy()
}

/**
 * VIEW
 */

func (bag Bag) Render() {
	oX := int32(glb.CELLS_OFFSET + bag.cellX*glb.CELL_WIDTH)
	oY := int32(glb.FIELD_OFFSET_Y + glb.CELLS_OFFSET + bag.cellY*glb.CELL_HEIGHT)
	ctx.RendererIns.Copy(bag.texture, nil, &sdl.Rect{oX, oY, glb.CELL_WIDTH, glb.CELL_HEIGHT})

	//debug
	//x1, y1, x2, y2 := bag.getHitBox()
	//ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
	//ctx.RendererIns.DrawRect(&sdl.Rect{x1, y1, x2 - x1, y2 - y1})
}
