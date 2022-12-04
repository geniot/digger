package gui

import (
	"geniot.com/geniot/digger/internal/model"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	loop   *Loop
	window *Window
	config *Config
	scene  *model.Scene
}

func NewApplication() Application {
	return Application{nil, nil, nil, nil}
}

func (app Application) Start() {
	sdl.Init(sdl.INIT_EVERYTHING)

	scn := model.NewScene()
	app.scene = &scn
	cnf := NewConfig(&app)
	app.config = &cnf
	wnd := NewWindow(&app)
	app.window = &wnd
	lp := NewLoop(&app)
	app.loop = &lp

	app.loop.Start()
}
