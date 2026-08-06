package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MoveGrid struct {
	scene     *GameScene
	texture   rl.RenderTexture2D
	sourceRec rl.Rectangle
	destRec   rl.Rectangle
	dots      [FIELD_WIDTH][FIELD_HEIGHT]bool
	moveMap   map[Direction]Pos
	dirMap    map[Direction][3]int32
}

func NewMoveGrid(scene *GameScene) *MoveGrid {
	moveGrid := &MoveGrid{}
	moveGrid.scene = scene

	moveGrid.moveMap = map[Direction]Pos{
		LEFT:  {-1, 0},
		RIGHT: {1, 0},
		UP:    {0, -1},
		DOWN:  {0, 1},
	}
	moveGrid.dirMap = map[Direction][3]int32{
		LEFT:  {int32(LEFT), int32(UP), int32(DOWN)},
		RIGHT: {int32(RIGHT), int32(UP), int32(DOWN)},
		UP:    {int32(UP), int32(LEFT), int32(RIGHT)},
		DOWN:  {int32(DOWN), int32(LEFT), int32(RIGHT)},
	}

	moveGrid.sourceRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), -float32(SCREEN_LOGICAL_HEIGHT)) //see https://github.com/raysan5/raylib/issues/3803
	moveGrid.destRec = rl.NewRectangle(0, 0, float32(SCREEN_LOGICAL_WIDTH), float32(SCREEN_LOGICAL_HEIGHT))

	moveGrid.texture = rl.LoadRenderTexture(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)
	moveGrid.dots = [FIELD_WIDTH][FIELD_HEIGHT]bool{}

	for y := FIELD_OFFSET_Y + CELL_HEIGHT/2; y < FIELD_HEIGHT-FIELD_OFFSET_Y-CELL_HEIGHT/2; y += CELL_HEIGHT {
		for x := FIELD_OFFSET_X + CELL_WIDTH/2; x < FIELD_WIDTH-FIELD_OFFSET_X-CELL_WIDTH/2; x += 1 {
			moveGrid.dots[x][y] = true
		}
	}
	for x := FIELD_OFFSET_X + CELL_WIDTH/2; x < FIELD_WIDTH-CELL_WIDTH/2; x += CELL_WIDTH {
		for y := FIELD_OFFSET_Y + CELL_HEIGHT/2; y < FIELD_HEIGHT-FIELD_OFFSET_Y-CELL_HEIGHT/2-1; y += 1 {
			moveGrid.dots[x][y] = true
		}
	}

	rl.BeginTextureMode(moveGrid.texture)
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

func (mg *MoveGrid) Update(_ int64) {
}

func (mg *MoveGrid) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(mg.texture.Texture, mg.sourceRec, mg.destRec, ZERO_VECTOR2, 0, rl.White)
	//rl.DrawFPS(5, 5)
	rl.EndTextureMode()
}

func (mg *MoveGrid) getDiggerStartPos() (int32, int32) {
	cellX := int32(7)
	cellY := int32(9)
	posX := FIELD_OFFSET_X + CELL_WIDTH/2 + cellX*CELL_WIDTH
	posY := FIELD_OFFSET_Y + CELL_HEIGHT/2 + cellY*CELL_HEIGHT
	if !mg.dots[posX][posY] {
		panic("digger start pos should be valid")
	}
	return posX, posY
}

func (mg *MoveGrid) canMove(x1, x2, x3, y1, y2, y3 int32) bool {
	limit := CELL_WIDTH / 3
	for i := int32(0); i < limit; i++ {
		if mg.dots[x1-x2*i+x3][y1-y2*i+y3] {
			return true
		}
	}
	return false
}

func (mg *MoveGrid) move(x int32, y int32, actualDirection Direction, requestedDirection Direction) (int32, int32, Direction) {
	direction := requestedDirection
	d0 := Direction(mg.dirMap[requestedDirection][0])
	d1 := Direction(mg.dirMap[requestedDirection][1])
	d2 := Direction(mg.dirMap[requestedDirection][2])
	m0 := mg.moveMap[d0]
	m1 := mg.moveMap[d1]
	m2 := mg.moveMap[d2]

	if requestedDirection == d0 {
		if mg.dots[x+m0.X][y+m0.Y] { //most expected situation
			x += m0.X
			y += m0.Y
		} else {
			if mg.dots[x-m0.X][y-m0.Y] { //border cases
				direction = d0
			} else {
				if actualDirection == d1 {
					if mg.canMove(x, m1.X, m0.X, y, m1.Y, m0.Y) {
						x -= m1.X
						y -= m1.Y
						direction = d2
					} else if mg.dots[x+m1.X][y+m1.Y] {
						x += m1.X
						y += m1.Y
						direction = d1
					}
				}
				if actualDirection == d2 {
					if mg.canMove(x, m2.X, m0.X, y, m2.Y, m0.Y) {
						x -= m2.X
						y -= m2.Y
						direction = d1
					} else if mg.dots[x+m2.X][y+m2.Y] {
						x += m2.X
						y += m2.Y
						direction = d2
					}
				}
			}
		}
	}
	return x, y, direction
}
