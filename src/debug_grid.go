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
	debugGrid.sourceRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), -float32(SCREEN_LOGICAL_HEIGHT)) //see https://github.com/raysan5/raylib/issues/3803
	debugGrid.destRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), float32(SCREEN_LOGICAL_HEIGHT))

	debugGrid.texture = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	rl.BeginTextureMode(debugGrid.texture)
	for y := int32(0); y < CELLS_VERTICAL+1; y += 1 {
		rl.DrawLine(0, y*CELL_HEIGHT+FIELD_OFFSET_Y, FIELD_WIDTH, y*CELL_HEIGHT+FIELD_OFFSET_Y, rl.White)

	}
	for x := int32(0); x < CELLS_HORIZONTAL+1; x += 1 {
		rl.DrawLine(x*CELL_WIDTH+FIELD_OFFSET_X, 0, x*CELL_WIDTH+FIELD_OFFSET_X, FIELD_HEIGHT, rl.White)
	}
	rl.DrawRectangleLines(0, 0, SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT, rl.Yellow)
	rl.EndTextureMode()

	return debugGrid
}

func (debugGrid *DebugGrid) Update(_ int64) {
}

func (debugGrid *DebugGrid) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(debugGrid.texture.Texture, debugGrid.sourceRec, debugGrid.destRec, ZERO_VECTOR2, 0, rl.White)
	//rl.DrawFPS(5, 5)
	rl.EndTextureMode()
}
