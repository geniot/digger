package gui

import (
	"geniot.com/geniot/digger/internal/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Application interface {
	Start()
}

type application struct {
	gameLoop GameLoop
	window   Window
	config   utils.Config
}

func NewApplication() Application {
	sdl.Init(sdl.INIT_EVERYTHING)
	cfg := utils.NewConfig()
	return application{
		NewGameLoop(),
		NewWindow(cfg),
		cfg}
}

func (app application) Start() {
	println("init")
	defer println("free")
	app.gameLoop.Start(app.window)
}
