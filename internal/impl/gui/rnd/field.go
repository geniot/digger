package rnd

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
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

func (field *Field) drawEatRight(x int32, y int32) {
	sourceRect := &sdl.Rect{int32(math.Mod(float64(x), float64(field.horizontalBlob.W))), 0, 1, field.horizontalBlob.H}
	targetTunnelRect := sdl.Rect{x + CELL_WIDTH - field.horizontalBlob.W, y - CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT}
	field.horizontalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{x + CELL_WIDTH - field.endRightBlob.W + 2, y - CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT}
	field.endRightBlob.Blit(nil, field.background, &targetEndRect)
}

func (field *Field) drawEatLeft(x int32, y int32) {
	sourceRect := &sdl.Rect{int32(math.Mod(float64(x), float64(field.horizontalBlob.W))), 0, 1, field.horizontalBlob.H}
	targetTunnelRect := sdl.Rect{x + field.horizontalBlob.W, y - CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT}
	field.horizontalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{x - 2, y - CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT}
	field.endLeftBlob.Blit(nil, field.background, &targetEndRect)
}

func (field *Field) drawEatUp(x int32, y int32) {
	sourceRect := &sdl.Rect{0, int32(math.Mod(float64(y), float64(field.verticalBlob.H))), field.verticalBlob.W, 1}
	targetTunnelRect := sdl.Rect{x, y - CELL_HEIGHT + field.verticalBlob.H, CELL_WIDTH, CELL_HEIGHT}
	field.verticalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{x, y - CELL_HEIGHT - field.endUpBlob.H + 2, CELL_WIDTH, CELL_HEIGHT}
	field.endUpBlob.Blit(nil, field.background, &targetEndRect)
}

func (field *Field) drawEatDown(x int32, y int32) {
	sourceRect := &sdl.Rect{0, int32(math.Mod(float64(y), float64(field.verticalBlob.H))), field.verticalBlob.W, 1}
	targetTunnelRect := sdl.Rect{x, y - field.verticalBlob.H, CELL_WIDTH, CELL_HEIGHT}
	field.verticalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{x, y - 3, CELL_WIDTH, CELL_HEIGHT}
	field.endDownBlob.Blit(nil, field.background, &targetEndRect)
}

func (field *Field) eatVertical(x int, y int, isUpCont bool, isDownCont bool) {
	oX := int32(CELLS_OFFSET + x*CELL_WIDTH)
	oY := int32(CELLS_OFFSET + y*CELL_HEIGHT + CELL_HEIGHT)

	for i := CELL_HEIGHT / 2; i >= If(isUpCont, -CELL_HEIGHT/2, 0); i-- {
		oY = int32(CELLS_OFFSET + y*CELL_HEIGHT + CELL_HEIGHT + i)
		field.drawEatUp(oX, oY)
	}
	for i := -CELL_HEIGHT / 2; i <= If(isDownCont, CELL_HEIGHT/2, 0); i++ {
		oY = int32(CELLS_OFFSET + y*CELL_HEIGHT + CELL_HEIGHT + i)
		field.drawEatDown(oX, oY)
	}
}

func (field *Field) eatHorizontal(x int, y int, isRightCont bool, isLeftCont bool) {
	oX := int32(CELLS_OFFSET + x*CELL_WIDTH)
	oY := int32(CELLS_OFFSET + y*CELL_HEIGHT + CELL_HEIGHT)

	for i := -CELL_WIDTH / 2; i <= If(isRightCont, CELL_WIDTH/2, 0); i++ {
		oX = int32(CELLS_OFFSET + x*CELL_WIDTH + i)
		field.drawEatRight(oX, oY)
	}
	for i := CELL_WIDTH / 2; i >= If(isLeftCont, -CELL_WIDTH/2, 0); i-- {
		oX = int32(CELLS_OFFSET + x*CELL_WIDTH + i)
		field.drawEatLeft(oX, oY)
	}
}
