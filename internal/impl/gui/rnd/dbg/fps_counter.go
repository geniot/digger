package dbg

import (
	"github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/internal/impl/gui/rnd"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type FpsCounter struct {
	startTicks    uint32
	frameCount    uint32
	currentSecond uint32
	currentFPS    uint32
}

func NewFpsCounter() *FpsCounter {
	return &FpsCounter{sdl.GetTicks(), 0, sdl.GetTicks() / 1000, 0}
}

func (fpsCounter *FpsCounter) Step(n uint64) {

}

func (fpsCounter *FpsCounter) Render() {
	fpsCounter.frameCount += 1
	currentTicks := sdl.GetTicks()
	newFps := uint32(0)
	if currentTicks != 0 {
		ticksDelta := currentTicks - fpsCounter.startTicks
		frameDuration := ticksDelta / fpsCounter.frameCount
		if frameDuration != 0 {
			newFps = 1000 / frameDuration
		}
	}

	sec := currentTicks / 1000
	if sec > fpsCounter.currentSecond {
		fpsCounter.currentFPS = newFps
		fpsCounter.frameCount = 0
		fpsCounter.startTicks = currentTicks
		fpsCounter.currentSecond = sec
	}
	rnd.drawText("FPS: "+strconv.FormatInt(int64(fpsCounter.currentFPS), 10), 10, 10, glb.COLOR_WHITE)
}
