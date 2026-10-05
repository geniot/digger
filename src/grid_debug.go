package main

import rl "github.com/gen2brain/raylib-go/raylib"

type DebugGrid struct {
	scene     *GameScene
	texture   rl.RenderTexture2D
	sourceRec rl.Rectangle
	destRec   rl.Rectangle
}

func NewDebugGrid(scene *GameScene) *DebugGrid {
	debugGrid := &DebugGrid{}
	debugGrid.scene = scene
	debugGrid.sourceRec = rl.NewRectangle(0, 0, float32(FieldWidth), -float32(FieldHeight)) //see https://github.com/raysan5/raylib/issues/3803
	debugGrid.destRec = rl.NewRectangle(0, 0, float32(FieldWidth), float32(FieldHeight))

	debugGrid.texture = rl.LoadRenderTexture(FieldWidth, FieldHeight)
	rl.BeginTextureMode(debugGrid.texture)
	for y := range CellsVertical + 1 {
		rl.DrawLine(0, y*CellHeight+FieldOffsetY, FieldWidth, y*CellHeight+FieldOffsetY, rl.White)

	}
	for x := range CellsHorizontal + 1 {
		rl.DrawLine(x*CellWidth+FieldOffsetX, 0, x*CellWidth+FieldOffsetX, FieldHeight, rl.White)
	}
	rl.DrawRectangleLinesEx(rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(FieldWidth),
		Height: float32(FieldHeight),
	}, 1, rl.Yellow)
	rl.EndTextureMode()

	return debugGrid
}

func (debugGrid *DebugGrid) Update(_ int64) {
}

func (debugGrid *DebugGrid) Render() {
	//rl.DrawTexturePro(debugGrid.texture.Texture, debugGrid.sourceRec, debugGrid.destRec, ZERO_VECTOR2, 0, rl.White)
	//rl.DrawFPS(5, 5)
}
