package gui

import (
	"geniot.com/geniot/digger/internal/model"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Window struct {
	application *Application
	sdlWindow   *sdl.Window
	sdlSurface  *sdl.Surface
}

func NewWindow(app *Application) Window {
	wnd, _ := sdl.CreateWindow(
		model.WINDOW_TITLE,
		int32(app.config.Get(model.WINDOW_XPOS_KEY)),
		int32(app.config.Get(model.WINDOW_YPOS_KEY)),
		int32(app.config.Get(model.WINDOW_WIDTH_KEY)),
		int32(app.config.Get(model.WINDOW_HEIGHT_KEY)),
		app.config.Get(model.WINDOW_STATE_KEY))
	surface, _ := wnd.GetSurface()

	return Window{app, wnd, surface}
}

func (window Window) Redraw() {
	surface, _ := window.sdlWindow.GetSurface()
	window.sdlSurface = surface

	window.sdlSurface.FillRect(nil, sdl.MapRGB(window.sdlSurface.Format, 16, 16, 16))
	window.sdlWindow.UpdateSurface()
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
