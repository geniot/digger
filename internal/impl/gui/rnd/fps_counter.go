package rnd

import (
	"geniot.com/geniot/digger/internal/glb"
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

func (fpsCounter *FpsCounter) Render() {
	fpsCounter.frameCount += 1
	currentTicks := sdl.GetTicks()
	if currentTicks == 0 {
		return
	}
	fps := 1000 / ((currentTicks - fpsCounter.startTicks) / fpsCounter.frameCount)

	sec := currentTicks / 1000
	if sec > fpsCounter.currentSecond {
		fpsCounter.currentFPS = fps
		fpsCounter.frameCount = 0
		fpsCounter.startTicks = currentTicks
		fpsCounter.currentSecond = sec
	}
	drawText("FPS: "+strconv.FormatInt(int64(fpsCounter.currentFPS), 10), 10, 10, glb.COLOR_BLACK)
}
