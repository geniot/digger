package gui

type PhysicsLoop struct {
	application *Application
}

func NewPhysicsLoop(app *Application) *PhysicsLoop {
	return &PhysicsLoop{app}
}

func (physicsLoop PhysicsLoop) Run() {

}
