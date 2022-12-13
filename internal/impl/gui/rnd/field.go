package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"geniot.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Field struct {
	background *sdl.Surface
	hBlob      *sdl.Surface
	scene      *Scene
}

func NewField(scn *Scene) *Field {
	hB, _ := img.LoadRW(resources.GetResource("hfblob.png"), true)
	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)
	defer bgrTile.Free()
	bgrSurface, _ := sdl.CreateRGBSurfaceWithFormat(0,
		glb.SCREEN_LOGICAL_WIDTH,
		glb.SCREEN_LOGICAL_HEIGHT,
		32, ctx.SurfaceIns.Format.Format)
	for i := 0; i < glb.SCREEN_LOGICAL_WIDTH/20; i++ {
		for j := 0; j < glb.SCREEN_LOGICAL_HEIGHT/4; j++ {
			bgrTile.Blit(
				&sdl.Rect{0, 0, 20, 4},
				bgrSurface,
				&sdl.Rect{int32(i * 20), int32(j * 4), 20, 4})
		}
	}

	return &Field{bgrSurface, hB, scn}
}

func (field Field) Render() {
	bgrTexture, _ := ctx.RendererIns.CreateTextureFromSurface(field.background)
	defer bgrTexture.Destroy()
	ctx.RendererIns.Copy(bgrTexture, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) Step(n uint64) {
	blobRect := sdl.Rect{int32(math.Mod(float64(field.scene.digger.offsetX), 4)), 0, 1, 20}
	diggerRect := sdl.Rect{field.scene.digger.offsetX + 20, field.scene.digger.offsetY - 20, 20, 20}
	field.hBlob.Blit(&blobRect, field.background, &diggerRect)

}
