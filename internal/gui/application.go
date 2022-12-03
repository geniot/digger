package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	loop   *Loop
	window *Window
	config *Config
}

func NewApplication() Application {
	return Application{nil, nil, nil}
}

func (app Application) Start() {
	sdl.Init(sdl.INIT_EVERYTHING)
	println("init")
	defer println("free")

	cnf := NewConfig(&app)
	app.config = &cnf
	wnd := NewWindow(&app)
	app.window = &wnd
	lp := NewLoop(&app)
	app.loop = &lp

	app.loop.Start()
}
