package rnd

import (
	. "github.com/geniot/digger/src/chs"
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/geniot/digger/src/res"
	"github.com/veandco/go-sdl2/sdl"
)

type DebugGrid struct {
	texture      *sdl.Texture
	textureChase *sdl.Texture
	textureBag   *sdl.Texture
	scene        *Scene
}

func NewDebugGrid(scn *Scene) *DebugGrid {
	dg := &DebugGrid{}
	dg.texture = res.LoadTexture("dbg_field.png")
	dg.textureChase = res.LoadTexture("dbg_chase.png")
	dg.textureBag = res.LoadTexture("dbg_bag.png")
	dg.scene = scn
	return dg
}

func (debugGrid *DebugGrid) Render() {
	//whole area
	ctx.RendererIns.SetDrawColor(255, 0, 0, 255)
	DrawRectLines(&sdl.Rect{W: SCREEN_LOGICAL_WIDTH - 1, H: SCREEN_LOGICAL_HEIGHT - 1})
	//stats
	ctx.RendererIns.SetDrawColor(0, 0, 255, 255)
	DrawRectLines(&sdl.Rect{X: 1, Y: 1, W: SCREEN_LOGICAL_WIDTH - 3, H: 20 - 2})

	ctx.RendererIns.SetDrawColor(0, 255, 0, 255)
	//horizontal lines
	for y := 0; y < CELLS_VERTICAL+1; y++ {
		x1 := int32(CELLS_OFFSET)
		y1 := int32(FIELD_OFFSET_Y + CELLS_OFFSET + y*CELL_HEIGHT)
		x2 := int32(CELLS_OFFSET + CELL_WIDTH*CELLS_HORIZONTAL)
		y2 := int32(FIELD_OFFSET_Y + CELLS_OFFSET + y*CELL_HEIGHT)
		ctx.RendererIns.DrawLine(x1, y1, x2, y2)
	}
	//vertical lines
	for x := 0; x < CELLS_HORIZONTAL+1; x++ {
		ctx.RendererIns.DrawLine(
			int32(CELLS_OFFSET+x*CELL_WIDTH), FIELD_OFFSET_Y+CELLS_OFFSET,
			int32(CELLS_OFFSET+x*CELL_WIDTH), FIELD_OFFSET_Y+CELLS_OFFSET+CELL_HEIGHT*CELLS_VERTICAL)
	}

	for monster := range debugGrid.scene.monsters.Iter() {
		for _, p := range monster.chasePath {
			pT := p.(*ChaseTile)
			ctx.RendererIns.Copy(debugGrid.textureChase, nil, &sdl.Rect{
				X: int32(CELLS_OFFSET + CELL_WIDTH/2 + pT.X*CELL_WIDTH/2 - 3),
				Y: int32(FIELD_OFFSET_Y + CELLS_OFFSET + CELL_HEIGHT/2 + pT.Y*CELL_HEIGHT/2 - 3),
				W: CELL_WIDTH / 3, H: CELL_HEIGHT / 3})
		}
	}

	for y := 0; y < CELLS_VERTICAL*2-1; y++ {
		for x := 0; x < CELLS_HORIZONTAL*2-1; x++ {
			tile := debugGrid.scene.chaseWorld.Tile(x, y)
			if tile.Kind == KindField {
				ctx.RendererIns.Copy(debugGrid.texture, nil, &sdl.Rect{
					X: int32(CELLS_OFFSET + CELL_WIDTH/2 + x*CELL_WIDTH/2 - 3),
					Y: int32(FIELD_OFFSET_Y + CELLS_OFFSET + CELL_HEIGHT/2 + y*CELL_HEIGHT/2 - 3),
					W: CELL_WIDTH / 3, H: CELL_HEIGHT / 3})
			} else if tile.Kind == KindBag {
				ctx.RendererIns.Copy(debugGrid.textureBag, nil, &sdl.Rect{
					X: int32(CELLS_OFFSET + CELL_WIDTH/2 + x*CELL_WIDTH/2 - 3),
					Y: int32(FIELD_OFFSET_Y + CELLS_OFFSET + CELL_HEIGHT/2 + y*CELL_HEIGHT/2 - 3),
					W: CELL_WIDTH / 3, H: CELL_HEIGHT / 3})
			}
		}
	}

}
