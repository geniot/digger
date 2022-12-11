package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	. "geniot.com/geniot/digger/internal/glb"
)

type DebugGrid struct {
}

func NewDebugGrid() *DebugGrid {
	return &DebugGrid{}
}

func (debugGrid DebugGrid) Render() {
	ctx.RendererIns.SetDrawColor(0, 255, 0, 255)
	//horizontal lines
	for y := 0; y < CELLS_VERTICAL+1; y++ {
		x1 := int32(CELLS_OFFSET_X)
		y1 := int32(CELLS_OFFSET_Y + y*CELL_HEIGHT - 1)
		x2 := int32(CELLS_OFFSET_X + CELL_WIDTH*CELLS_HORIZONTAL)
		y2 := int32(CELLS_OFFSET_Y + y*CELL_HEIGHT - 1)
		ctx.RendererIns.DrawLine(x1, y1, x2, y2)
	}
	//vertical lines
	for x := 0; x < CELLS_HORIZONTAL+1; x++ {
		ctx.RendererIns.DrawLine(
			int32(CELLS_OFFSET_X+x*CELL_WIDTH), CELLS_OFFSET_Y,
			int32(CELLS_OFFSET_X+x*CELL_WIDTH), CELLS_OFFSET_Y+CELL_HEIGHT*CELLS_VERTICAL)
	}
}

func (debugGrid *DebugGrid) Step(n uint64) {

}
