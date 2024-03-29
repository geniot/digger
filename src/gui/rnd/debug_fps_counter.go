package rnd

import (
	"github.com/geniot/digger/src/ctx"
	"github.com/geniot/digger/src/glb"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type DebugFpsCounter struct {
	startTicks     uint32
	frameCount     uint32
	currentSecond  uint32
	currentFPS     uint32
	cachedTextures map[string]*glb.SurfTexture
}

func NewFpsCounter() *DebugFpsCounter {
	dbg := &DebugFpsCounter{}
	dbg.startTicks = sdl.GetTicks()
	dbg.frameCount = 0
	dbg.currentSecond = sdl.GetTicks() / 1000
	dbg.currentFPS = 0
	dbg.cachedTextures = make(map[string]*glb.SurfTexture)
	return dbg
}

func (fpsCounter *DebugFpsCounter) Render() {
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
	txt := "FPS: " + strconv.FormatInt(int64(fpsCounter.currentFPS), 10)
	texture := fpsCounter.cachedTextures[txt]
	if texture != nil {
		ctx.RendererIns.Copy(texture.T, nil, &sdl.Rect{X: 5, Y: 5, W: texture.W, H: texture.H})
	} else {
		fpsCounter.cachedTextures[txt] = glb.DrawText(txt, glb.COLOR_WHITE)
	}
}
