package loop

import (
	"geniot.com/geniot/digger/internal/ctx"
)

type RenderLoopImpl struct {
}

func NewRenderLoop() *RenderLoopImpl {
	return &RenderLoopImpl{}
}

func (renderLoop RenderLoopImpl) Run() {
	ctx.RendererIns.SetDrawColor(211, 211, 211, 255)
	ctx.RendererIns.Clear()
	ctx.SceneIns.Render()
	ctx.RendererIns.Present()
}
