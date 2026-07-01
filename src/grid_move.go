package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MoveGrid struct {
	scene     *GameScene
	texture   rl.RenderTexture2D
	sourceRec rl.Rectangle
	destRec   rl.Rectangle
	dots      [FIELD_WIDTH][FIELD_HEIGHT]bool
}

func NewMoveGrid(scene *GameScene) *MoveGrid {
	moveGrid := &MoveGrid{}
	moveGrid.scene = scene
	moveGrid.sourceRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), -float32(SCREEN_LOGICAL_HEIGHT)) //see https://github.com/raysan5/raylib/issues/3803
	moveGrid.destRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), float32(SCREEN_LOGICAL_HEIGHT))

	moveGrid.texture = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	rl.BeginTextureMode(moveGrid.texture)

	moveGrid.dots = [FIELD_WIDTH][FIELD_HEIGHT]bool{}

	for x := FIELD_OFFSET_X + CELL_WIDTH/2; x < FIELD_WIDTH-FIELD_OFFSET_X-CELL_WIDTH/2; x += 1 {
		moveGrid.dots[x][FIELD_OFFSET_Y+CELL_HEIGHT/2] = true

	}
	for x := int32(0); x < FIELD_WIDTH; x += 1 {
		for y := int32(0); y < FIELD_HEIGHT; y += 1 {
			if moveGrid.dots[x][y] {
				rl.DrawPixel(x, y, rl.Red)
			}
		}
	}
	rl.EndTextureMode()

	return moveGrid
}

func (moveGrid *MoveGrid) Update(_ int64) {
}

func (moveGrid *MoveGrid) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(moveGrid.texture.Texture, moveGrid.sourceRec, moveGrid.destRec, ZERO_VECTOR2, 0, rl.White)
	//rl.DrawFPS(5, 5)
	rl.EndTextureMode()
}
