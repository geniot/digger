package gui

import (
	"geniot.com/geniot/digger/internal/model"
	"geniot.com/geniot/digger/internal/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Window interface {
	Redraw()
}

type window struct {
	wnd     *sdl.Window
	surface *sdl.Surface
	config  utils.Config
}

func NewWindow(cfg utils.Config) Window {
	wnd, _ := sdl.CreateWindow(
		model.WINDOW_TITLE,
		cfg.Get(model.WINDOW_XPOS_KEY),
		cfg.Get(model.WINDOW_YPOS_KEY),
		cfg.Get(model.WINDOW_WIDTH_KEY),
		cfg.Get(model.WINDOW_HEIGHT_KEY),
		sdl.WINDOW_SHOWN)
	surface, _ := wnd.GetSurface()

	return window{wnd, surface, cfg}
}

func (w window) Redraw() {
	w.surface.FillRect(nil, sdl.MapRGB(w.surface.Format, 16, 16, 16))
	w.wnd.UpdateSurface()
}
