package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Field struct {
	horizontalBlob *sdl.Surface
	verticalBlob   *sdl.Surface
	endLeftBlob    *sdl.Surface
	endRightBlob   *sdl.Surface
	endUpBlob      *sdl.Surface
	endDownBlob    *sdl.Surface
	background     *sdl.Surface
	scene          *Scene
}

/**
 * INIT
 */

func NewField(scn *Scene) *Field {
	fld := &Field{}
	fld.scene = scn

	fld.horizontalBlob, _ = img.LoadRW(resources.GetResource("blob1.png"), true)
	fld.verticalBlob, _ = img.LoadRW(resources.GetResource("blob2.png"), true)
	fld.endLeftBlob, _ = img.LoadRW(resources.GetResource("blob3.png"), true)
	fld.endRightBlob, _ = img.LoadRW(resources.GetResource("blob4.png"), true)
	fld.endUpBlob, _ = img.LoadRW(resources.GetResource("blob5.png"), true)
	fld.endDownBlob, _ = img.LoadRW(resources.GetResource("blob6.png"), true)

	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)
	defer bgrTile.Free()
	info, _ := ctx.RendererIns.GetInfo()
	fld.background, _ = sdl.CreateRGBSurfaceWithFormat(0,
		SCREEN_LOGICAL_WIDTH,
		SCREEN_LOGICAL_HEIGHT,
		32, uint32(info.TextureFormats[0]))

	for i := 0; i < SCREEN_LOGICAL_WIDTH/int(bgrTile.W); i++ {
		for j := 0; j < SCREEN_LOGICAL_HEIGHT/int(bgrTile.H); j++ {
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
	ctx.RendererIns.Copy(bgrTexture, nil, &sdl.Rect{0, FIELD_OFFSET_Y, SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) drawEat(tunnelSurface *sdl.Surface, x1 int32, y1 int32, endSurface *sdl.Surface, x2 int32, y2 int32, sourceRect *sdl.Rect) {
	targetTunnelRect := sdl.Rect{x1, y1, CELL_WIDTH, CELL_HEIGHT}
	tunnelSurface.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{x2, y2, CELL_WIDTH, CELL_HEIGHT}
	endSurface.Blit(nil, field.background, &targetEndRect)
}
