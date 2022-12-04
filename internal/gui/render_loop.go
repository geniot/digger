package gui

type RenderLoop struct {
	application *Application
}

func NewRenderLoop(app *Application) *RenderLoop {
	return &RenderLoop{app}
}

func (renderLoop RenderLoop) Run() {
	window := renderLoop.application.window
	window.sdlRenderer.SetDrawColor(255, 255, 255, 255)
	window.sdlRenderer.Clear()
	renderLoop.application.scene.Render(window.sdlRenderer)
	window.sdlRenderer.Present()
}
