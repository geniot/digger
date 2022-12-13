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
	background   *sdl.Surface
	tunnelBlob   *sdl.Surface
	endLeftBlob  *sdl.Surface
	endRightBlob *sdl.Surface
	scene        *Scene
}

func NewField(scn *Scene) *Field {
	tunnelB, _ := img.LoadRW(resources.GetResource("blob1.png"), true)
	endLeftB, _ := img.LoadRW(resources.GetResource("blob2.png"), true)
	endRightB, _ := img.LoadRW(resources.GetResource("blob3.png"), true)
	bgrTile, _ := img.LoadRW(resources.GetResource("cback1.png"), true)

	defer bgrTile.Free()

	bgrSurface, _ := sdl.CreateRGBSurfaceWithFormat(0,
		glb.SCREEN_LOGICAL_WIDTH,
		glb.SCREEN_LOGICAL_HEIGHT,
		32, ctx.SurfaceIns.Format.Format)

	for i := 0; i < glb.SCREEN_LOGICAL_WIDTH/int(bgrTile.W); i++ {
		for j := 0; j < glb.SCREEN_LOGICAL_HEIGHT/int(bgrTile.H); j++ {
			bgrTile.Blit(
				&sdl.Rect{0, 0, bgrTile.W, bgrTile.H},
				bgrSurface,
				&sdl.Rect{int32(i * int(bgrTile.W)), int32(j * int(bgrTile.H)), bgrTile.W, bgrTile.H})
		}
	}

	return &Field{bgrSurface, tunnelB, endLeftB, endRightB, scn}
}

func (field Field) Render() {
	bgrTexture, _ := ctx.RendererIns.CreateTextureFromSurface(field.background)
	defer bgrTexture.Destroy()
	ctx.RendererIns.Copy(bgrTexture, nil, &sdl.Rect{0, glb.FIELD_OFFSET_Y, glb.SCREEN_LOGICAL_WIDTH, glb.SCREEN_LOGICAL_HEIGHT})
}

func (field *Field) Step(n uint64) {
	sourceTunnelRect := sdl.Rect{int32(math.Mod(float64(field.scene.digger.offsetX), float64(field.tunnelBlob.W))), 0, 1, field.tunnelBlob.H}
	if field.scene.digger.direction == glb.RIGHT {
		targetTunnelRect := sdl.Rect{field.scene.digger.offsetX + glb.CELL_WIDTH - field.tunnelBlob.W, field.scene.digger.offsetY - glb.CELL_HEIGHT, glb.CELL_WIDTH, glb.CELL_HEIGHT}
		targetEndRect := sdl.Rect{field.scene.digger.offsetX + glb.CELL_WIDTH - field.endRightBlob.W, field.scene.digger.offsetY - glb.CELL_HEIGHT, glb.CELL_WIDTH, glb.CELL_HEIGHT}
		field.tunnelBlob.Blit(&sourceTunnelRect, field.background, &targetTunnelRect)
		field.endRightBlob.Blit(nil, field.background, &targetEndRect)
	} else if field.scene.digger.direction == glb.LEFT {
		targetTunnelRect := sdl.Rect{field.scene.digger.offsetX + field.tunnelBlob.W, field.scene.digger.offsetY - glb.CELL_HEIGHT, glb.CELL_WIDTH, glb.CELL_HEIGHT}
		targetEndRect := sdl.Rect{field.scene.digger.offsetX, field.scene.digger.offsetY - glb.CELL_HEIGHT, glb.CELL_WIDTH, glb.CELL_HEIGHT}
		field.tunnelBlob.Blit(&sourceTunnelRect, field.background, &targetTunnelRect)
		field.endLeftBlob.Blit(nil, field.background, &targetEndRect)
	}

}
