package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
	"geniot.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Digger struct {
	width   int32
	height  int32
	offsetX int32
	offsetY int32
	texture *sdl.Texture
}

func NewDigger() Digger {
	surface, _ := img.LoadRW(resources.GetResource("cldig1.png"), true)
	defer surface.Free()
	txt, err := ctx.RendererIns.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return Digger{surface.W, surface.H, glb.SCREEN_LOGICAL_WIDTH / 2, glb.SCREEN_LOGICAL_HEIGHT / 2, txt}
}

func (digger Digger) Render() {
	dstRect := sdl.Rect{digger.offsetX, digger.offsetY, digger.width, digger.height}
	ctx.RendererIns.Copy(digger.texture, nil, &dstRect)
}
