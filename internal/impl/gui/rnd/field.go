package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Field struct {
	background *sdl.Surface
	scene      *Scene
}

/**
 * INIT
 */

func NewField(scn *Scene) *Field {
	fld := &Field{}
	fld.scene = scn
	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)
	defer bgrTile.Free()
	info, _ := ctx.RendererIns.GetInfo()
	fld.background, _ = sdl.CreateRGBSurfaceWithFormat(0,
		glb.SCREEN_LOGICAL_WIDTH,
		glb.SCREEN_LOGICAL_HEIGHT,
		32, uint32(info.TextureFormats[0]))

	for i := 0; i < glb.SCREEN_LOGICAL_WIDTH/int(bgrTile.W); i++ {
		for j := 0; j < glb.SCREEN_LOGICAL_HEIGHT/int(bgrTile.H); j++ {
			bgrTile.Blit(
				&sdl.Rect{0, 0, bgrTile.W, bgrTile.H},
				fld.background,
				&sdl.Rect{int32(i * int(bgrTile.W)), int32(j * int(bgrTile.H)), bgrTile.W, bgrTile.H})
		}
	}
	return fld
}

/**
 * MODEL
 */

func (field *Field) Step(n uint64) {

}

/**
 * VIEW
 */

func (field Field) Render() {
	bgrTexture, _ := ctx.RendererIns.CreateTextureFromSurface(field.background)
	defer bgrTexture.Destroy()
	ctx.RendererIns.Copy(bgrTexture, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}
