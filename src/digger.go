package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Digger struct {
	scene            *GameScene
	posX             int32
	posY             int32
	innerOffsetX     int32
	innerOffsetY     int32
	direction        Direction
	shouldMove       bool
	spritePointer    int
	spritePointerInc int
	sprites          []*TextureImage
}

func NewDigger(scene *GameScene) *Digger {
	digger := &Digger{}
	digger.scene = scene

	digger.sprites = make([]*TextureImage, 3)
	digger.sprites[0] = NewTextureImage("cldig1.png")
	digger.sprites[1] = NewTextureImage("cldig2.png")
	digger.sprites[2] = NewTextureImage("cldig3.png")

	//same for all levels
	cellX := int32(7)
	cellY := int32(9)
	digger.innerOffsetX = int32(0)
	digger.innerOffsetY = int32(1)
	digger.spritePointer = 0
	digger.spritePointerInc = 1
	digger.posX = FIELD_OFFSET_X + cellX*CELL_WIDTH - digger.innerOffsetX + CELL_WIDTH/2
	digger.posY = FIELD_OFFSET_Y + cellY*CELL_HEIGHT - digger.innerOffsetY + CELL_HEIGHT/2
	digger.direction = RIGHT
	digger.shouldMove = false
	return digger
}

func (digger *Digger) Update(tick int64) {
	if tick%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer, digger.spritePointerInc = GetNextSpritePointerAndInc(digger.spritePointer, digger.spritePointerInc, len(digger.sprites))
	}
	if tick%DIGGER_SPEED == 0 && digger.shouldMove {
		if digger.direction == RIGHT {
			digger.handleMove(RIGHT, UP, DOWN, (FIELD_OFFSET_Y+digger.posY)%CELL_HEIGHT)
		} else if digger.direction == LEFT {
			digger.handleMove(LEFT, UP, DOWN, (FIELD_OFFSET_Y+digger.posY)%CELL_HEIGHT)
		} else if digger.direction == UP {
			digger.handleMove(UP, LEFT, RIGHT, (digger.posX)%CELL_WIDTH)
		} else if digger.direction == DOWN {
			digger.handleMove(DOWN, LEFT, RIGHT, (digger.posX)%CELL_WIDTH)
		}
	}
}

func (digger *Digger) Render(drawTarget rl.RenderTexture2D) {
	sourceRec := rl.NewRectangle(0, 0, float32(IfInt(digger.direction == LEFT, CELL_WIDTH, -CELL_WIDTH)), float32(CELL_HEIGHT))
	destRec := rl.NewRectangle(float32(digger.posX), float32(digger.posY), float32(CELL_WIDTH), float32(CELL_HEIGHT))
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexturePro(
		digger.sprites[digger.spritePointer].texture,
		sourceRec,
		destRec,
		CELL_CENTER_VECTOR2,
		float32(IfInt(digger.direction == LEFT || digger.direction == RIGHT, 0, IfInt(digger.direction == UP, 270, 90))),
		rl.White)
	rl.EndTextureMode()
}

func (digger *Digger) handleMove(dir1 Direction, dir2 Direction, dir3 Direction, mod int32) {
	if digger.direction == dir1 {
		if cM, _ := digger.canMoveShouldTurn(dir1); cM {
			digger.move(dir1)
		}
	} else if digger.direction == Opposite(dir1) {
		digger.direction = dir1
	} else {
		if mod != 0 {
			if digger.direction == dir2 {
				cM, sT := digger.canMoveShouldTurn(dir2)
				if cM {
					digger.move(dir2)
				}
				if sT {
					digger.direction = Opposite(dir2)
				}
			} else if digger.direction == dir3 {
				cM, sT := digger.canMoveShouldTurn(dir3)
				if cM {
					digger.move(dir3)
				}
				if sT {
					digger.direction = Opposite(dir3)
				}
			}
		} else {
			digger.direction = dir1
		}
	}
}

func (digger *Digger) canMoveShouldTurn(dir Direction) (bool, bool) {
	if !digger.scene.field.isWithinBounds(dir, digger.posX, digger.posY) {
		return false, false
	}
	return true, false
}

func (digger *Digger) move(dir Direction) {
	digger.direction = dir
	x := If(dir == RIGHT, 1, If(dir == LEFT, -1, 0))
	y := If(dir == DOWN, 1, If(dir == UP, -1, 0))
	digger.posX += int32(x)
	digger.posY += int32(y)
	//digger.collisionObject.X = float64(digger.offsetX + digger.innerOffsetX)
	//digger.collisionObject.Y = float64(digger.offsetY + digger.innerOffsetY)
	//digger.collisionObject.Update()
}
