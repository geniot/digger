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
		app.config.Get(model.WINDOW_XPOS_KEY),
		app.config.Get(model.WINDOW_YPOS_KEY),
		app.config.Get(model.WINDOW_WIDTH_KEY),
		app.config.Get(model.WINDOW_HEIGHT_KEY),
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
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
	window.application.config.Set(model.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(width), 10))
	window.application.config.Set(model.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(height), 10))
	window.application.config.Set(model.WINDOW_XPOS_KEY, strconv.FormatInt(int64(xPos), 10))
	window.application.config.Set(model.WINDOW_YPOS_KEY, strconv.FormatInt(int64(yPos), 10))
	window.application.config.Save()
}
