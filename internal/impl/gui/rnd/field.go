package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"geniot.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Field struct {
	background *sdl.Texture
}

func NewField() *Field {
	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)
	defer bgrTile.Free()
	bgrSurface, _ := sdl.CreateRGBSurfaceWithFormat(0,
		glb.SCREEN_LOGICAL_WIDTH,
		glb.SCREEN_LOGICAL_HEIGHT,
		32, ctx.SurfaceIns.Format.Format)
	defer bgrSurface.Free()
	for i := 0; i < glb.SCREEN_LOGICAL_WIDTH/20; i++ {
		for j := 0; j < glb.SCREEN_LOGICAL_HEIGHT/4; j++ {
			bgrTile.Blit(
				&sdl.Rect{0, 0, 20, 4},
				bgrSurface,
				&sdl.Rect{int32(i * 20), int32(j * 4), 20, 4})
		}
	}
	bgrTexture, _ := ctx.RendererIns.CreateTextureFromSurface(bgrSurface)
	return &Field{bgrTexture}
}

func (field Field) Render() {
	ctx.RendererIns.Copy(field.background, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) Step(n uint64) {

}
