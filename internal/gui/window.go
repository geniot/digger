package gui

import "github.com/veandco/go-sdl2/sdl"

type Window interface {
	Redraw()
}

type window struct {
	wnd     *sdl.Window
	surface *sdl.Surface
}

func NewWindow() Window {
	wnd, _ := sdl.CreateWindow(
		"digger",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		300, 300,
		sdl.WINDOW_SHOWN)
	surface, _ := wnd.GetSurface()

	return window{wnd, surface}
}

func (w window) Redraw() {
	w.surface.FillRect(nil, sdl.MapRGB(w.surface.Format, 16, 16, 16))
	w.wnd.UpdateSurface()
}
