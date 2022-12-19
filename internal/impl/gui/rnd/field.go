package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Field struct {
	background     *sdl.Surface
	horizontalBlob *sdl.Surface
	verticalBlob   *sdl.Surface
	endLeftBlob    *sdl.Surface
	endRightBlob   *sdl.Surface
	endUpBlob      *sdl.Surface
	endDownBlob    *sdl.Surface
	scene          *Scene
}

func NewField(scn *Scene) *Field {
	horizontalB, _ := img.LoadRW(resources.GetResource("blob1.png"), true)
	verticalB, _ := img.LoadRW(resources.GetResource("blob2.png"), true)
	endLeftB, _ := img.LoadRW(resources.GetResource("blob3.png"), true)
	endRightB, _ := img.LoadRW(resources.GetResource("blob4.png"), true)
	endUpB, _ := img.LoadRW(resources.GetResource("blob5.png"), true)
	endDownB, _ := img.LoadRW(resources.GetResource("blob6.png"), true)

	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)

	defer bgrTile.Free()

	info, _ := ctx.RendererIns.GetInfo()
	bgrSurface, _ := sdl.CreateRGBSurfaceWithFormat(0,
		glb.SCREEN_LOGICAL_WIDTH,
		glb.SCREEN_LOGICAL_HEIGHT,
		32, uint32(info.TextureFormats[0]))

	for i := 0; i < glb.SCREEN_LOGICAL_WIDTH/int(bgrTile.W); i++ {
		for j := 0; j < glb.SCREEN_LOGICAL_HEIGHT/int(bgrTile.H); j++ {
			bgrTile.Blit(
				&sdl.Rect{0, 0, bgrTile.W, bgrTile.H},
				bgrSurface,
				&sdl.Rect{int32(i * int(bgrTile.W)), int32(j * int(bgrTile.H)), bgrTile.W, bgrTile.H})
		}
	}

	return &Field{bgrSurface, horizontalB, verticalB, endLeftB, endRightB, endUpB, endDownB, scn}
}

func (field Field) Render() {
	bgrTexture, _ := ctx.RendererIns.CreateTextureFromSurface(field.background)
	defer bgrTexture.Destroy()
	ctx.RendererIns.Copy(bgrTexture, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) eatField(tunnelSurface *sdl.Surface, x1 int32, y1 int32, endSurface *sdl.Surface, x2 int32, y2 int32, sourceRect *sdl.Rect) {
	targetTunnelRect := sdl.Rect{field.scene.digger.offsetX + x1, field.scene.digger.offsetY + y1, glb.CELL_WIDTH, glb.CELL_HEIGHT}
	tunnelSurface.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{field.scene.digger.offsetX + x2, field.scene.digger.offsetY + y2, glb.CELL_WIDTH, glb.CELL_HEIGHT}
	endSurface.Blit(nil, field.background, &targetEndRect)
}

func (field *Field) Step(n uint64) {
	sourceHorizontalTunnelRect := sdl.Rect{int32(math.Mod(float64(field.scene.digger.offsetX), float64(field.horizontalBlob.W))), 0, 1, field.horizontalBlob.H}
	sourceVerticalTunnelRect := sdl.Rect{0, int32(math.Mod(float64(field.scene.digger.offsetY), float64(field.verticalBlob.H))), field.verticalBlob.W, 1}

	if field.scene.digger.direction == glb.RIGHT { //RIGHT

		field.eatField(
			field.horizontalBlob,
			glb.CELL_WIDTH-field.horizontalBlob.W,
			-glb.CELL_HEIGHT,
			field.endRightBlob,
			glb.CELL_WIDTH-field.endRightBlob.W+2,
			-glb.CELL_HEIGHT,
			&sourceHorizontalTunnelRect)

	} else if field.scene.digger.direction == glb.LEFT { //LEFT

		field.eatField(
			field.horizontalBlob,
			field.horizontalBlob.W,
			-glb.CELL_HEIGHT,
			field.endLeftBlob,
			-2,
			-glb.CELL_HEIGHT,
			&sourceHorizontalTunnelRect)

	} else if field.scene.digger.direction == glb.UP { //UP

		field.eatField(
			field.verticalBlob,
			0,
			-glb.CELL_HEIGHT+field.verticalBlob.H,
			field.endUpBlob,
			0,
			-glb.CELL_HEIGHT-field.endUpBlob.H+2,
			&sourceVerticalTunnelRect)

	} else if field.scene.digger.direction == glb.DOWN { //DOWN

		field.eatField(
			field.verticalBlob,
			0,
			-field.verticalBlob.H,
			field.endDownBlob,
			0,
			-3,
			&sourceVerticalTunnelRect)

	}

}

func (field Field) eatEmerald(emerald *Emerald) {
	oX := int32(glb.CELLS_OFFSET_X + emerald.cellX*glb.CELL_WIDTH)
	oY := int32(glb.CELLS_OFFSET_Y + (emerald.cellY-1)*glb.CELL_HEIGHT)
	targetRect := sdl.Rect{
		oX,
		oY,
		glb.CELL_WIDTH, glb.CELL_HEIGHT}
	emerald.textureMask.Blit(nil, field.background, &targetRect)
}
