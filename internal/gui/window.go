package gui

import (
	"fmt"
	"geniot.com/geniot/digger/internal/model"
	"github.com/robfig/cron/v3"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Window struct {
	application *Application
	sdlWindow   *sdl.Window
}

func NewWindow(app *Application) Window {
	wnd, _ := sdl.CreateWindow(
		model.WINDOW_TITLE,
		int32(app.config.Get(model.WINDOW_XPOS_KEY)),
		int32(app.config.Get(model.WINDOW_YPOS_KEY)),
		int32(app.config.Get(model.WINDOW_WIDTH_KEY)),
		int32(app.config.Get(model.WINDOW_HEIGHT_KEY)),
		app.config.Get(model.WINDOW_STATE_KEY))

	w := Window{app, wnd}

	w.Redraw()

	c := cron.New()
	_, err := c.AddFunc("@every 16ms", w.Redraw) //fps=60
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Start()

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
	surface, _ := window.sdlWindow.GetSurface()
	surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))
	window.sdlWindow.UpdateSurface()
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
