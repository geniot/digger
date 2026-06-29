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

func (a *Application) onResize() {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())
	a.sourceRect = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), float32(-SCREEN_LOGICAL_HEIGHT))
	ratioX := screenWidth / float32(SCREEN_LOGICAL_WIDTH)
	ratioY := screenHeight / float32(SCREEN_LOGICAL_HEIGHT)
	resizeRatio := If(ratioX < ratioY, ratioX, ratioY)
	a.destRect = rl.NewRectangle(
		(screenWidth-(float32(SCREEN_LOGICAL_WIDTH)*resizeRatio))*0.5,
		(screenHeight-(float32(SCREEN_LOGICAL_HEIGHT)*resizeRatio))*0.5,
		float32(SCREEN_LOGICAL_WIDTH)*resizeRatio,
		float32(SCREEN_LOGICAL_HEIGHT)*resizeRatio,
	)
}

func (a *Application) ProcessInput() {
	if rl.IsWindowResized() {
		a.onResize()
	}
	a.scenes[a.currentSceneIndex].ProcessInput()
}

func (a *Application) Update(tick int64) {
	a.scenes[a.currentSceneIndex].Update(tick)
}

func (a *Application) Render() {
	a.scenes[a.currentSceneIndex].Render(a.drawTarget)
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTexturePro(a.drawTarget.Texture,
		a.sourceRect,
		a.destRect,
		ZERO_VECTOR2, 0, rl.White)
	rl.EndDrawing()
}

func NewApplication() *Application {

	app := Application{}

	// the order of these calls matters
	rl.SetTraceLogLevel(rl.LogTrace)
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable) //should be set before window initialization!
	rl.InitWindow(winWidth, winHeight, "Digger")
	rl.SetWindowMonitor(0) //used for testing on multiple monitors
	rl.InitAudioDevice()

	setDefaultTextStyle()

	// scenes
	app.scenes = make(map[int]Scene)
	app.scenes[gameSceneKey] = NewGameScene(&app)
	app.currentSceneIndex = gameSceneKey

	//debug
	//app.currentSceneIndex = controlsSceneKey

	app.drawTarget = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	app.onResize()

	return &app
}

func (a *Application) ShouldExit() bool {
	return rl.WindowShouldClose() || a.scenes[a.currentSceneIndex].ShouldExit()
}

func (a *Application) Exit() {
	rl.CloseWindow()
}
