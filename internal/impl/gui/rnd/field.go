package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"geniot.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type Field struct {
	background  *sdl.Texture
	scene       *Scene
	blackPixels *[20 * 20 * 4]uint32
}

func NewField(scn *Scene) *Field {
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

	pixels := [20 * 20 * 4]uint32{}
	for x := 0; x < len(pixels); x += 4 {
		pixels[x] = 0
		pixels[x+1] = 0
		pixels[x+2] = 0
		pixels[x+3] = 0
	}

	return &Field{bgrTexture, scn, &pixels}
}

func (field Field) Render() {
	ctx.RendererIns.Copy(field.background, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) Step(n uint64) {
	diggerRect := sdl.Rect{field.scene.digger.offsetX, field.scene.digger.offsetY - 20, 20, 20}
	field.background.Update(&diggerRect, unsafe.Pointer(&field.blackPixels[0]), 20*4)

}
