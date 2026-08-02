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

func (moveGrid *MoveGrid) Update(_ int64) {
}

func (moveGrid *MoveGrid) Render(drawTarget rl.RenderTexture2D) {
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(moveGrid.texture.Texture, moveGrid.sourceRec, moveGrid.destRec, ZERO_VECTOR2, 0, rl.White)
	//rl.DrawFPS(5, 5)
	rl.EndTextureMode()
}

func (moveGrid *MoveGrid) getDiggerStartPos() (int32, int32) {
	cellX := int32(7)
	cellY := int32(9)
	posX := FIELD_OFFSET_X + CELL_WIDTH/2 + cellX*CELL_WIDTH
	posY := FIELD_OFFSET_Y + CELL_HEIGHT/2 + cellY*CELL_HEIGHT
	if !moveGrid.dots[posX][posY] {
		panic("digger start pos should be valid")
	}
	return posX, posY
}

func (moveGrid *MoveGrid) move(x int32, y int32, actualDirection Direction, requestedDirection Direction) (int32, int32, Direction) {
	direction := requestedDirection
	m := map[Direction][9]int32{
		LEFT:  {int32(LEFT), -1, 0, int32(UP), 0, -1, int32(DOWN), 0, 1},
		RIGHT: {int32(RIGHT), 1, 0, int32(UP), 0, -1, int32(DOWN), 0, 1},
		UP:    {int32(UP), 0, -1, int32(LEFT), -1, 0, int32(RIGHT), 1, 0},
		DOWN:  {int32(DOWN), 0, 1, int32(LEFT), -1, 0, int32(RIGHT), 1, 0},
	}
	values := m[requestedDirection]
	if requestedDirection == Direction(values[0]) {
		if moveGrid.dots[x+values[1]][y+values[2]] {
			x += values[1]
			y += values[2]
		} else {
			if moveGrid.dots[x-values[1]][y-values[2]] {
				direction = Direction(values[0])
			} else {
				if actualDirection == Direction(values[3]) && moveGrid.dots[x+values[4]][y+values[5]] {
					x += values[4]
					y += values[5]
					direction = actualDirection
				}
				if actualDirection == Direction(values[6]) && moveGrid.dots[x+values[7]][y+values[8]] {
					x += values[7]
					y += values[8]
					direction = actualDirection
				}
			}
		}
	}

	//LEFT,-1,0,UP,0,-1,DOWN,0,1
	//if requestedDirection == LEFT {
	//	if moveGrid.dots[x-1][y] {
	//		x += -1
	//		y += 0
	//	} else {
	//		if moveGrid.dots[x+1][y] {
	//			direction = LEFT
	//		} else {
	//			if actualDirection == UP && moveGrid.dots[x][y-1] {
	//				y += -1
	//				direction = actualDirection
	//			}
	//			if actualDirection == DOWN && moveGrid.dots[x][y+1] {
	//				y += 1
	//				direction = actualDirection
	//			}
	//		}
	//	}
	//}
	//if requestedDirection == RIGHT {
	//	if moveGrid.dots[x+1][y] {
	//		x += 1
	//	} else {
	//		if moveGrid.dots[x-1][y] {
	//			direction = RIGHT
	//		} else {
	//			if actualDirection == UP && moveGrid.dots[x][y-1] {
	//				y -= 1
	//				direction = actualDirection
	//			}
	//			if actualDirection == DOWN && moveGrid.dots[x][y+1] {
	//				y += 1
	//				direction = actualDirection
	//			}
	//		}
	//	}
	//}
	//if requestedDirection == UP {
	//	if moveGrid.dots[x][y-1] {
	//		y -= 1
	//	} else {
	//		if moveGrid.dots[x][y+1] {
	//			direction = UP
	//		} else {
	//			if actualDirection == LEFT && moveGrid.dots[x-1][y] {
	//				x -= 1
	//				direction = actualDirection
	//			}
	//			if actualDirection == RIGHT && moveGrid.dots[x+1][y] {
	//				x += 1
	//				direction = actualDirection
	//			}
	//		}
	//	}
	//}
	//if requestedDirection == DOWN {
	//	if moveGrid.dots[x][y+1] {
	//		y += 1
	//	} else {
	//		if moveGrid.dots[x][y-1] {
	//			direction = DOWN
	//		} else {
	//			if actualDirection == LEFT && moveGrid.dots[x-1][y] {
	//				x -= 1
	//				direction = actualDirection
	//			}
	//			if actualDirection == RIGHT && moveGrid.dots[x+1][y] {
	//				x += 1
	//				direction = actualDirection
	//			}
	//		}
	//	}
	//}
	return x, y, direction
}
