package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type Digger struct {
	width   int32
	height  int32
	offsetX int32
	offsetY int32

	spritePointer    int
	spritePointerInc int
	sprites          []*sdl.Texture
}

func NewDigger() *Digger {
	spts := []*sdl.Texture{loadTexture("cldig1.png"), loadTexture("cldig2.png"), loadTexture("cldig3.png")}
	return &Digger{16, 15, glb.SCREEN_LOGICAL_WIDTH / 2, glb.SCREEN_LOGICAL_HEIGHT / 2, 0, 1, spts}
}

func (digger Digger) Render() {
	dstRect := sdl.Rect{digger.offsetX, digger.offsetY, digger.width, digger.height}
	ctx.RendererIns.CopyEx(digger.sprites[digger.spritePointer], nil, &dstRect, 0, &sdl.Point{0, 0}, sdl.FLIP_HORIZONTAL)
}

func (digger *Digger) Step(n uint64) {
	if n%8 == 0 {
		digger.spritePointer += digger.spritePointerInc
		if digger.spritePointer == len(digger.sprites)-1 || digger.spritePointer == 0 {
			digger.spritePointerInc = -digger.spritePointerInc
		}
	}

	if n%2 == 0 {
		if ctx.PressedKeysCodesSetIns.Contains(glb.GCW_BUTTON_LEFT) {
			digger.offsetX -= 1
		}
		if ctx.PressedKeysCodesSetIns.Contains(glb.GCW_BUTTON_RIGHT) {
			digger.offsetX += 1
		}
	}

}
