package loop

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type EventLoopImpl struct {
}

func NewEventLoop() *EventLoopImpl {
	return &EventLoopImpl{}
}

func (eventLoop EventLoopImpl) Run() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			if t.State == sdl.PRESSED {
				ctx.PressedKeysCodesSetIns[t.Keysym.Sym] = time.Now().UnixNano()
			} else { // if t.State == sdl.RELEASED {
				delete(ctx.PressedKeysCodesSetIns, t.Keysym.Sym)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				ctx.WindowIns.SaveWindowState()
			}
			break

		case *sdl.QuitEvent:
			ctx.LoopIns.Stop()
			break
		}
	}
	ctx.DeviceIns.ProcessKeyActions()
}
