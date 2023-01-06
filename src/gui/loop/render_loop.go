package loop

import (
	"github.com/geniot/digger/src/ctx"
	"github.com/geniot/digger/src/glb"
)

type RenderLoopImpl struct {
}

func NewRenderLoop() *RenderLoopImpl {
	return &RenderLoopImpl{}
}

func (renderLoop RenderLoopImpl) Run() {
	ctx.RendererIns.SetDrawColor(glb.BGR_COLOR[0], glb.BGR_COLOR[1], glb.BGR_COLOR[2], glb.BGR_COLOR[3])
	ctx.RendererIns.Clear()
	ctx.SceneIns.Render()
	ctx.RendererIns.Present()
}
