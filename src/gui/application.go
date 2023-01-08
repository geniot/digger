package gui

import (
	"github.com/geniot/digger/src/ctx"
	"github.com/geniot/digger/src/dev"
	"github.com/geniot/digger/src/glb"
	loop2 "github.com/geniot/digger/src/gui/loop"
	"github.com/geniot/digger/src/gui/rnd"
	"github.com/geniot/digger/src/res"
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

	ctx.LoopIns = loop2.NewLoop()
	ctx.EventLoopIns = loop2.NewEventLoop()
	ctx.PhysicsLoopIns = loop2.NewPhysicsLoop()
	ctx.RenderLoopIns = loop2.NewRenderLoop()

	ctx.FontIns, _ = ttf.OpenFontRW(res.GetImage(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)

	ctx.SceneIns = rnd.NewScene()

	ctx.LoopIns.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) Stop() {
	ctx.FontIns.Close()
	ctx.DeviceIns.Stop()
}
