package gui

type Application interface {
	Start()
}

type application struct {
	gameLoop GameLoop
	window   Window
}

func NewApplication() Application {
	return application{NewGameLoop(), NewWindow()}
}

func (app application) Start() {
	println("init")
	defer println("free")
	app.gameLoop.Start(app.window)
}
