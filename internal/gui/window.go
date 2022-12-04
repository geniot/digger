package gui

import (
	"geniot.com/geniot/digger/internal/model"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Window struct {
	application *Application
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer
}

func NewWindow(app *Application) Window {
	wnd, _ := sdl.CreateWindow(
		model.WINDOW_TITLE,
		int32(app.config.Get(model.WINDOW_XPOS_KEY)),
		int32(app.config.Get(model.WINDOW_YPOS_KEY)),
		int32(app.config.Get(model.WINDOW_WIDTH_KEY)),
		int32(app.config.Get(model.WINDOW_HEIGHT_KEY)),
		app.config.Get(model.WINDOW_STATE_KEY))

	rnd, _ := sdl.CreateRenderer(wnd, -1,
		sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	//sdl.RENDERER_ACCELERATED)
	w := Window{app, wnd, rnd}

	sdl.AddEventWatchFunc(w.resizingEventWatcher, nil)

	return w
}

func (window Window) resizingEventWatcher(event sdl.Event, data interface{}) bool {
	switch t := event.(type) {
	case *sdl.WindowEvent:
		if t.Event == sdl.WINDOWEVENT_RESIZED {
			window.Redraw()
		}
		break
	}
	return false
}

func (window Window) Redraw() {
	window.sdlRenderer.SetDrawColor(255, 255, 255, 255)
	window.sdlRenderer.Clear()

	width, height := window.sdlWindow.GetSize()
	factor := int32(10)
	rects := []sdl.Rect{{
		width / factor,
		height / factor,
		width - (width/factor)*2,
		height - (height/factor)*2}}

	window.sdlRenderer.SetDrawColor(0, 0, 255, 255)
	window.sdlRenderer.DrawRects(rects)
	window.sdlRenderer.Present()
}

func (window Window) OnBeforeClose() {
	window.SaveWindowState()
}

func (window Window) SaveWindowState() {
	width, height := window.sdlWindow.GetSize()
	xPos, yPos := window.sdlWindow.GetPosition()
	windowState := window.sdlWindow.GetFlags()
	window.application.config.Set(model.WINDOW_STATE_KEY, strconv.FormatInt(int64(windowState), 10))

	if windowState&sdl.WINDOW_MAXIMIZED <= 0 {
		window.application.config.Set(model.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(width), 10))
		window.application.config.Set(model.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(height), 10))
		window.application.config.Set(model.WINDOW_XPOS_KEY, strconv.FormatInt(int64(xPos), 10))
		window.application.config.Set(model.WINDOW_YPOS_KEY, strconv.FormatInt(int64(yPos), 10))
	}

	window.application.config.Save()
}
