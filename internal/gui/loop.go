package gui

import "github.com/tevino/abool/v2"

type Loop struct {
	application *Application
	isRunning   *abool.AtomicBool
	eventLoop   *EventLoop
	physicsLoop *PhysicsLoop
	renderLoop  *RenderLoop
}

func NewLoop(app *Application) *Loop {
	eventLoop := NewEventLoop(app)
	physicsLoop := NewPhysicsLoop(app)
	renderLoop := NewRenderLoop(app)
	return &Loop{app, abool.New(), eventLoop, physicsLoop, renderLoop}
}

func (loop Loop) Start() {
	loop.isRunning.Set()
	for loop.isRunning.IsSet() {
		loop.eventLoop.Run()
		loop.physicsLoop.Run()
		loop.renderLoop.Run()
	}
}
