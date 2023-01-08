package rnd

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/geniot/digger/src/res"
	"github.com/solarlune/resolv"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Emerald struct {
	offsetX      int32
	offsetY      int32
	width        int32
	height       int32
	innerOffsetX int32
	innerOffsetY int32

	texture         *sdl.Texture
	textureMask     *sdl.Surface
	collisionObject *resolv.Object
	scene           *Scene
}

/**
 * INIT
 */

func NewEmerald(cX int, cY int, scn *Scene) *Emerald {
	em := &Emerald{}
	em.scene = scn
	em.texture = res.LoadTexture("emerald.png")
	em.textureMask, _ = img.LoadRW(res.GetImage("emerald_mask.png"), true)

	em.offsetX = int32(CELLS_OFFSET + cX*CELL_WIDTH)
	em.offsetY = int32(FIELD_OFFSET_Y + CELLS_OFFSET + cY*CELL_HEIGHT)
	em.width = 10
	em.height = 8
	em.innerOffsetX = 5
	em.innerOffsetY = 7

	em.collisionObject = resolv.NewObject(float64(em.offsetX+em.innerOffsetX), float64(em.offsetY+em.innerOffsetY),
		float64(em.width), float64(em.height), EMERALD_COLLISION_TAG)
	em.collisionObject.Data = em
	scn.collisionSpace.Add(em.collisionObject)

	return em
}

/**
 * MODEL
 */

func (emerald *Emerald) getHitBox() *sdl.Rect {
	return &sdl.Rect{X: emerald.offsetX + emerald.innerOffsetX, Y: emerald.offsetY + emerald.innerOffsetY, W: emerald.width, H: emerald.height}
}

func (emerald *Emerald) Destroy() {
	emerald.scene.soundEat()
	emerald.scene.field.eatEmerald(emerald)
	emerald.textureMask.Free()
	emerald.texture.Destroy()
	emerald.scene.collisionSpace.Remove(emerald.collisionObject)
	emerald.scene.emeralds.Remove(emerald)
}

/**
 * VIEW
 */

func (emerald *Emerald) Render() {
	ctx.RendererIns.Copy(emerald.texture, nil, &sdl.Rect{X: emerald.offsetX, Y: emerald.offsetY, W: CELL_WIDTH, H: CELL_HEIGHT})

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(emerald.getHitBox())
	}

}
