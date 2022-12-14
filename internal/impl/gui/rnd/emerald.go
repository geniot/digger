package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"geniot.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Emerald struct {
	cellX       int
	cellY       int
	texture     *sdl.Texture
	textureMask *sdl.Surface
	scene       *Scene
}

func NewEmerald(scn *Scene) *Emerald {
	txt := loadTexture("emerald.png")
	txtMask, _ := img.LoadRW(resources.GetResource("emerald_mask.png"), true)
	cX := rand.Intn(glb.CELLS_HORIZONTAL)
	cY := rand.Intn(glb.CELLS_VERTICAL)
	return &Emerald{cX, cY, txt, txtMask, scn}
}

func (emerald Emerald) getHitBox() (int32, int32, int32, int32) {
	oX := int32(glb.CELLS_OFFSET_X + emerald.cellX*glb.CELL_WIDTH)
	oY := int32(glb.CELLS_OFFSET_Y + emerald.cellY*glb.CELL_HEIGHT)
	return oX + 5, oY + 7, oX + glb.CELL_WIDTH - 5, oY + glb.CELL_HEIGHT - 5
}

func (emerald Emerald) Render() {

	oX := int32(glb.CELLS_OFFSET_X + emerald.cellX*glb.CELL_WIDTH)
	oY := int32(glb.CELLS_OFFSET_Y + emerald.cellY*glb.CELL_HEIGHT)
	ctx.RendererIns.Copy(emerald.texture, nil, &sdl.Rect{oX, oY, glb.CELL_WIDTH, glb.CELL_HEIGHT})

	//debug
	//x1, y1, x2, y2 := emerald.getHitBox()
	//ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
	//ctx.RendererIns.DrawRect(&sdl.Rect{x1, y1, x2 - x1, y2 - y1})
}

func (emerald *Emerald) Step(n uint64) {

}

func (emerald *Emerald) Destroy() {
	emerald.textureMask.Free()
	emerald.texture.Destroy()
}
