package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Application struct {
	scenes            map[int]Scene
	drawTarget        rl.RenderTexture2D
	currentSceneIndex int
	sourceRect        rl.Rectangle
	destRect          rl.Rectangle
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[a.currentSceneIndex].ShouldExit()
}

func (a *Application) Update() {
	if rl.IsWindowResized() {
		a.updateRects()
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	a.scenes[a.currentSceneIndex].Update(a.drawTarget)

	rl.DrawTexturePro(a.drawTarget.Texture,
		a.sourceRect,
		a.destRect,
		rl.Vector2{}, 0, rl.White)
	rl.EndDrawing()
}

func (a *Application) Exit() {
	rl.CloseWindow()
}

func (a *Application) updateRects() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	a.sourceRect = rl.NewRectangle(0, float32(-SCREEN_LOGICAL_HEIGHT), float32(SCREEN_LOGICAL_WIDTH), float32(-SCREEN_LOGICAL_HEIGHT))
	ratioX := screenWidth / float32(SCREEN_LOGICAL_WIDTH)
	ratioY := screenHeight / float32(SCREEN_LOGICAL_HEIGHT)
	resizeRatio := If(ratioX < ratioY, ratioX, ratioY)
	a.destRect = rl.NewRectangle(
		(screenWidth-(SCREEN_LOGICAL_WIDTH*resizeRatio))*0.5,
		(screenHeight-(SCREEN_LOGICAL_HEIGHT*resizeRatio))*0.5,
		SCREEN_LOGICAL_WIDTH*resizeRatio,
		SCREEN_LOGICAL_HEIGHT*resizeRatio,
	)
}

func NewApplication() *Application {

	app := Application{}

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogTrace)
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable) //should be set before window initialization!
	rl.InitWindow(winWidth, winHeight, "TrimUI Digger")
	rl.SetWindowMonitor(0) //used for testing on multiple monitors
	rl.InitAudioDevice()
	rl.SetClipPlanes(0.5, 100) //these values are found by trial and error for TrimUI, see https://github.com/raysan5/raylib/issues/4917
	rl.DisableBackfaceCulling()

	setDefaultTextStyle()

	// scenes
	app.scenes = make(map[int]Scene)
	//app.scenes[menuSceneKey] = NewMenuScene(&app)
	app.scenes[gameSceneKey] = NewGameScene(&app)
	//app.scenes[tutorialSceneKey] = NewTutorialScene(&app)
	//app.scenes[controlsSceneKey] = NewControlsScene(&app)
	app.currentSceneIndex = gameSceneKey

	//debug
	//app.currentSceneIndex = controlsSceneKey

	app.drawTarget = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	rl.BeginTextureMode(app.drawTarget)
	rl.ClearBackground(rl.Black)
	rl.EndTextureMode()

	app.updateRects()

	return &app
}
