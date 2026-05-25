package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Application struct {
	camera                *rl.Camera
	scenes                map[int]Scene
	currentSceneIndex     int
	colorTextures         map[int]rl.Texture2D
	selectedColorTextures map[int]rl.Texture2D
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[a.currentSceneIndex].ShouldExit()
}

func (a *Application) Update() {
	a.scenes[a.currentSceneIndex].Update(a.camera)
}

func (a *Application) Exit() {
	rl.CloseWindow()
}

func NewApplication() *Application {

	app := Application{}

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogTrace)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	rl.InitWindow(winWidth, winHeight, "TrimUI Digger")
	rl.SetWindowMonitor(0) //used for testing on multiple monitors
	rl.InitAudioDevice()
	rl.SetClipPlanes(0.5, 100) //these values are found by trial and error for TrimUI, see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	setDefaultTextStyle()

	//camera
	app.camera = &rl.Camera3D{}
	app.camera.Position = rl.NewVector3(10, 10, 10)
	app.camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	app.camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	app.camera.Fovy = 40.0
	app.camera.Projection = rl.CameraPerspective

	// textures
	app.colorTextures = make(map[int]rl.Texture2D)
	app.selectedColorTextures = make(map[int]rl.Texture2D)
	//prepareTextures(app.colorTextures, false)
	//prepareTextures(app.selectedColorTextures, true)

	// scenes
	app.scenes = make(map[int]Scene)
	//app.scenes[menuSceneKey] = NewMenuScene(&app)
	app.scenes[gameSceneKey] = NewGameScene(&app)
	//app.scenes[tutorialSceneKey] = NewTutorialScene(&app)
	//app.scenes[controlsSceneKey] = NewControlsScene(&app)
	app.currentSceneIndex = gameSceneKey

	//debug
	//app.currentSceneIndex = controlsSceneKey

	return &app
}
