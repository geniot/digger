package loop

import (
	"geniot.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type PhysicsLoopImpl struct {
	accumulator       float64
	lastTime, lastFps float64
	frames, fps       int
}

func NewPhysicsLoop() *PhysicsLoopImpl {
	return &PhysicsLoopImpl{}
}

func (physicsLoop *PhysicsLoopImpl) Run() {
	t := float64(sdl.GetTicks()) / 1000
	dt := t - physicsLoop.lastTime
	if dt > 0.2 {
		dt = 0.2
	}
	physicsLoop.lastTime = t
	physicsLoop.frames++
	if t-physicsLoop.lastFps >= 1 {
		physicsLoop.fps = physicsLoop.frames
		physicsLoop.frames = 0
		physicsLoop.lastFps += 1
	}

	for physicsLoop.accumulator += dt; physicsLoop.accumulator > glb.TICK; physicsLoop.accumulator -= glb.TICK {
		//step
	}
}