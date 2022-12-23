package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
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
	em.texture = resources.LoadTexture("emerald.png")
	em.textureMask, _ = img.LoadRW(resources.GetResource("emerald_mask.png"), true)

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
	return &sdl.Rect{emerald.offsetX + emerald.innerOffsetX, emerald.offsetY + emerald.innerOffsetY, emerald.width, emerald.height}
}

func (emerald *Emerald) Destroy() {
	emerald.eatField()
	emerald.textureMask.Free()
	emerald.texture.Destroy()
	emerald.scene.collisionSpace.Remove(emerald.collisionObject)
	emerald.scene.emeralds.Remove(emerald)
}

/**
 * VIEW
 */

func (emerald *Emerald) Render() {
	ctx.RendererIns.Copy(emerald.texture, nil, &sdl.Rect{emerald.offsetX, emerald.offsetY, CELL_WIDTH, CELL_HEIGHT})

	if IS_DEBUG_ON {
		ctx.RendererIns.SetDrawColor(255, 255, 255, 255)
		DrawRectLines(emerald.getHitBox())
	}

}

func (emerald *Emerald) eatField() {
	targetRect := sdl.Rect{
		emerald.offsetX,
		emerald.offsetY - FIELD_OFFSET_Y,
		CELL_WIDTH, CELL_HEIGHT}
	emerald.textureMask.Blit(nil, emerald.scene.field.background, &targetRect)
}
