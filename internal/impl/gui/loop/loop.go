package loop

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/tevino/abool/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type LoopImpl struct {
	isRunning *abool.AtomicBool
}

func NewLoop() *LoopImpl {
	return &LoopImpl{abool.New()}
}

func (loop LoopImpl) Start() {
	loop.isRunning.Set()
	for loop.isRunning.IsSet() {
		ctx.EventLoopIns.Run()
		ctx.PhysicsLoopIns.Run()
		ctx.RenderLoopIns.Run()
		sdl.Delay(16)
	}
}

func (loop LoopImpl) Stop() {
	loop.isRunning.UnSet()
}
