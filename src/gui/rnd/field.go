package rnd

import (
	"github.com/geniot/digger/src/chs"
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/geniot/digger/src/res"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Field struct {
	background        *sdl.Surface
	backgroundTexture *sdl.Texture
	isChanged         bool
	scene             *Scene
}

/**
 * INIT
 */

func NewField(scn *Scene) *Field {
	fld := &Field{}
	fld.scene = scn

	bgrTile, _ := img.LoadRW(res.GetImage("cback1.png"), true)
	defer bgrTile.Free()
	info, _ := ctx.RendererIns.GetInfo()
	fld.background, _ = sdl.CreateRGBSurfaceWithFormat(0,
		SCREEN_LOGICAL_WIDTH,
		SCREEN_LOGICAL_HEIGHT,
		32, uint32(info.TextureFormats[0]))

	for i := 0; i < SCREEN_LOGICAL_WIDTH/int(bgrTile.W); i++ {
		for j := 0; j < SCREEN_LOGICAL_HEIGHT/int(bgrTile.H); j++ {
			bgrTile.Blit(
				&sdl.Rect{W: bgrTile.W, H: bgrTile.H},
				fld.background,
				&sdl.Rect{X: int32(i * int(bgrTile.W)), Y: int32(j * int(bgrTile.H)), W: bgrTile.W, H: bgrTile.H})
		}
	}
	fld.isChanged = true
	return fld
}

/**
 * VIEW
 */

func (field *Field) Render() {
	if field.isChanged {
		if field.backgroundTexture != nil {
			field.backgroundTexture.Destroy()
		}
		field.backgroundTexture, _ = ctx.RendererIns.CreateTextureFromSurface(field.background)
		field.isChanged = false
	}
	ctx.RendererIns.Copy(field.backgroundTexture, nil, &sdl.Rect{Y: FIELD_OFFSET_Y, W: SCREEN_LOGICAL_WIDTH, H: SCREEN_LOGICAL_HEIGHT})
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

func (field *Field) isPointField(x int32, y int32) bool {
	r, g, b, _ := field.background.At(int(x), int(y-FIELD_OFFSET_Y)).RGBA()
	return r != 0 || g != 0 || b != 0
}

func (field *Field) isWithinBounds(dir Direction, offsetX int32, offsetY int32) bool {
	//screen bounds
	switch dir {
	case RIGHT:
		return offsetX < CELLS_OFFSET+CELL_WIDTH*(CELLS_HORIZONTAL-1)
	case LEFT:
		return offsetX > CELLS_OFFSET
	case UP:
		return offsetY > FIELD_OFFSET_Y+CELLS_OFFSET
	case DOWN:
		return offsetY < FIELD_OFFSET_Y+CELLS_OFFSET+CELL_HEIGHT*(CELLS_VERTICAL-1)
	default:
		return true
	}
}

/*
	EAT FIELD
*/

func (field *Field) drawEatRight(x int32, y int32) {
	media := field.scene.media
	sourceRect := &sdl.Rect{X: x % media.fieldHorizontalBlob.W, W: 1, H: media.fieldHorizontalBlob.H}
	targetTunnelRect := sdl.Rect{X: x + CELL_WIDTH - media.fieldHorizontalBlob.W, Y: y - CELL_HEIGHT, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldHorizontalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{X: x + CELL_WIDTH - media.fieldEndRightBlob.W + 2, Y: y - CELL_HEIGHT, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldEndRightBlob.Blit(nil, field.background, &targetEndRect)
	field.updateChaseWorld(targetTunnelRect, targetEndRect)
	field.isChanged = true
}

func (field *Field) drawEatLeft(x int32, y int32) {
	media := field.scene.media
	sourceRect := &sdl.Rect{X: x % media.fieldHorizontalBlob.W, W: 1, H: media.fieldHorizontalBlob.H}
	targetTunnelRect := sdl.Rect{X: x + media.fieldHorizontalBlob.W, Y: y - CELL_HEIGHT, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldHorizontalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{X: x - 2, Y: y - CELL_HEIGHT, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldEndLeftBlob.Blit(nil, field.background, &targetEndRect)
	field.updateChaseWorld(targetTunnelRect, targetEndRect)
	field.isChanged = true
}

func (field *Field) drawEatUp(x int32, y int32) {
	media := field.scene.media
	sourceRect := &sdl.Rect{Y: y % media.fieldVerticalBlob.H, W: media.fieldVerticalBlob.W, H: 1}
	targetTunnelRect := sdl.Rect{X: x, Y: y - CELL_HEIGHT + media.fieldVerticalBlob.H, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldVerticalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{X: x, Y: y - CELL_HEIGHT - media.fieldEndUpBlob.H + 2, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldEndUpBlob.Blit(nil, field.background, &targetEndRect)
	field.updateChaseWorld(targetTunnelRect, targetEndRect)
	field.isChanged = true
}

func (field *Field) drawEatDown(x int32, y int32) {
	media := field.scene.media
	sourceRect := &sdl.Rect{Y: y % media.fieldVerticalBlob.H, W: media.fieldVerticalBlob.W, H: 1}
	targetTunnelRect := sdl.Rect{X: x, Y: y - media.fieldVerticalBlob.H, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldVerticalBlob.Blit(sourceRect, field.background, &targetTunnelRect)
	targetEndRect := sdl.Rect{X: x, Y: y - 3, W: CELL_WIDTH, H: CELL_HEIGHT}
	media.fieldEndDownBlob.Blit(nil, field.background, &targetEndRect)
	field.updateChaseWorld(targetTunnelRect, targetEndRect)
	field.isChanged = true
}

func (field *Field) eatEmerald(emerald *Emerald) {
	targetRect := sdl.Rect{
		X: emerald.offsetX,
		Y: emerald.offsetY - FIELD_OFFSET_Y,
		W: CELL_WIDTH, H: CELL_HEIGHT}
	emerald.scene.media.emeraldTextureMask.Blit(nil, field.background, &targetRect)
	field.updateChaseWorld(targetRect)
	field.isChanged = true
}

// translating rects to our grid, updating grid if necessary
func (field *Field) updateChaseWorld(rects ...sdl.Rect) {
	for i := 0; i < len(rects); i++ {
		rect := rects[i]
		//we change 1-2 cells max with one rect
		x1 := (rect.X - CELLS_OFFSET) / CELL_WIDTH
		y1 := (rect.Y - CELLS_OFFSET) / CELL_WIDTH
		x2 := (rect.X + rect.W - CELLS_OFFSET) / CELL_WIDTH
		y2 := (rect.Y + rect.H - CELLS_OFFSET) / CELL_WIDTH
		field.updateChaseTiles(x1, y1)
		if x2 != x1 || y2 != y1 {
			field.updateChaseTiles(x2, y2)
		}
	}
}

func (field *Field) updateChaseTiles(x int32, y int32) {
	x = If(x >= CELLS_HORIZONTAL, CELLS_HORIZONTAL-1, x)
	y = If(y >= CELLS_VERTICAL, CELLS_VERTICAL-1, y)

	//left link
	if x > 0 {
		tileX := If(x == 1, 1, int((x-1)*2-1))
		tileY := If(y == 0, 0, int(y*2))
		tile := field.scene.chaseWorld.Tile(tileX, tileY)
		if tile.Kind != chs.KindTunnel && field.isLinkOpen(
			CELLS_OFFSET+x*CELL_WIDTH-CELL_WIDTH,
			CELLS_OFFSET+y*CELL_HEIGHT+CELL_HEIGHT/2,
			CELL_WIDTH*CHASE_LINK_OPEN_COEFFICIENT,
			false) {
			tile.Kind = chs.KindTunnel
		}
	}
	//right link
	if x < CELLS_HORIZONTAL {
		tileX := If(x == 0, 1, int((x-1)*2+1))
		tileY := If(y == 0, 0, int(y*2))
		tile := field.scene.chaseWorld.Tile(tileX, tileY)
		if tile.Kind != chs.KindTunnel && field.isLinkOpen(
			CELLS_OFFSET+x*CELL_WIDTH,
			CELLS_OFFSET+y*CELL_HEIGHT+CELL_HEIGHT/2,
			CELL_WIDTH*CHASE_LINK_OPEN_COEFFICIENT,
			false) {
			tile.Kind = chs.KindTunnel
		}
	}
	//up link
	if y > 0 {
		tileX := If(x == 0, 0, int(x*2))
		tileY := If(y <= 1, 0, int((y-1)*2-1))
		tile := field.scene.chaseWorld.Tile(tileX, tileY)
		if tile.Kind != chs.KindTunnel && field.isLinkOpen(
			CELLS_OFFSET+x*CELL_WIDTH+CELL_WIDTH/2,
			CELLS_OFFSET+y*CELL_HEIGHT-CELL_HEIGHT,
			CELL_HEIGHT*CHASE_LINK_OPEN_COEFFICIENT,
			true) {
			tile.Kind = chs.KindTunnel
		}
	}
	//down link
	if y < CELLS_VERTICAL {
		tileX := If(x == 0, 0, int(x*2))
		tileY := If(y <= 1, 1, int((y-1)*2+1))
		tile := field.scene.chaseWorld.Tile(tileX, tileY)
		if tile.Kind != chs.KindTunnel && field.isLinkOpen(
			CELLS_OFFSET+x*CELL_WIDTH+CELL_WIDTH/2,
			CELLS_OFFSET+y*CELL_HEIGHT,
			CELL_HEIGHT*CHASE_LINK_OPEN_COEFFICIENT,
			true) {
			tile.Kind = chs.KindTunnel
		}
	}
}

func (field *Field) isLinkOpen(checkX int32, checkY int32, limit int32, isVertical bool) bool {
	isLinkOpen := true
	for i := int32(0); i < limit; i++ {
		if isVertical {
			if field.isPointField(checkX, checkY+i+FIELD_OFFSET_Y) || field.isPointField(checkX, checkY-i+FIELD_OFFSET_Y) {
				isLinkOpen = false
				break
			}
		} else { //horizontal
			if field.isPointField(checkX+i, checkY+FIELD_OFFSET_Y) || field.isPointField(checkX-i, checkY+FIELD_OFFSET_Y) {
				isLinkOpen = false
				break
			}
		}
	}
	return isLinkOpen
}
