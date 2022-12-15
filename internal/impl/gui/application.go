package gui

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/internal/impl/dev"
	"github.com/geniot/digger/internal/impl/gui/loop"
	"github.com/geniot/digger/internal/impl/gui/rnd"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/ttf"
)

type ApplicationImpl struct {
}

func NewApplication() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (app *ApplicationImpl) Start() {
	ctx.DeviceIns = dev.NewDevice()
	ctx.ConfigIns = NewConfig()
	ctx.WindowIns = NewWindow()

	ctx.LoopIns = loop.NewLoop()
	ctx.EventLoopIns = loop.NewEventLoop()
	ctx.PhysicsLoopIns = loop.NewPhysicsLoop()
	ctx.RenderLoopIns = loop.NewRenderLoop()

	ctx.SceneIns = rnd.NewScene()

	ctx.FontIns, _ = ttf.OpenFontRW(resources.GetResource(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)

	ctx.LoopIns.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) Stop() {
	ctx.FontIns.Close()
	ctx.DeviceIns.Stop()
}
