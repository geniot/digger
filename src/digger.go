package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Digger struct {
	scene              *GameScene
	posX               int32
	posY               int32
	width              int32
	height             int32
	direction          Direction
	requestedDirection Direction
	shouldMove         bool
	spritePointer      int
	spritePointerInc   int
	leftSprites        []*TextureImage
	rightSprites       []*TextureImage
	upSprites          []*TextureImage
	downSprites        []*TextureImage
}

func NewDigger(scene *GameScene) *Digger {
	digger := &Digger{}
	digger.scene = scene

	digger.leftSprites = digger.initSprites(0, false, false)
	digger.rightSprites = digger.initSprites(0, true, false)
	digger.upSprites = digger.initSprites(90, false, false)
	digger.downSprites = digger.initSprites(90, true, true)

	digger.spritePointer = 0
	digger.spritePointerInc = 1
	digger.width = 16
	digger.height = 16
	digger.posX, digger.posY = scene.moveGrid.getDiggerStartPos()
	digger.direction = RIGHT
	digger.shouldMove = false
	return digger
}

func (digger *Digger) initSprites(degrees int32, flipHorizontal bool, flipVertical bool) []*TextureImage {
	sprites := make([]*TextureImage, 3)
	sprites[0] = NewTextureImage("cldig1.png", degrees, flipHorizontal, flipVertical)
	sprites[1] = NewTextureImage("cldig2.png", degrees, flipHorizontal, flipVertical)
	sprites[2] = NewTextureImage("cldig3.png", degrees, flipHorizontal, flipVertical)
	return sprites
}

func (digger *Digger) Update(tick int64) {
	if tick%SPRITE_UPDATE_RATE == 0 {
		digger.spritePointer, digger.spritePointerInc = GetNextSpritePointerAndInc(digger.spritePointer, digger.spritePointerInc, len(digger.leftSprites))
	}
	if tick%DIGGER_SPEED == 0 && digger.shouldMove {
		posX, posY, dir := digger.scene.moveGrid.move(digger.posX, digger.posY, digger.direction, digger.requestedDirection)
		if digger.posX != posX || digger.posY != posY || digger.direction != dir {
			digger.posX, digger.posY, digger.direction = posX, posY, dir
			if dir == RIGHT {
				digger.scene.field.draw(float32(digger.posX-digger.posX%4+4), float32(digger.posY-CELL_HEIGHT/2+1), digger.scene.field.rightBlob)
			} else if dir == LEFT {
				digger.scene.field.draw(float32(digger.posX-digger.posX%4-CELL_WIDTH/2-2), float32(digger.posY-CELL_HEIGHT/2+1), digger.scene.field.leftBlob)
			} else if dir == UP {
				digger.scene.field.draw(float32(digger.posX-CELL_WIDTH/2), float32(digger.posY-CELL_HEIGHT/2-digger.posY%3+1), digger.scene.field.upBlob)
			} else if dir == DOWN {
				digger.scene.field.draw(float32(digger.posX-CELL_WIDTH/2), float32(digger.posY+CELL_HEIGHT/2-digger.posY%3-2), digger.scene.field.downBlob)
			}
		}
	}
}

func (digger *Digger) getSprites() []*TextureImage {
	switch digger.direction {
	case RIGHT:
		return digger.rightSprites
	case LEFT:
		return digger.leftSprites
	case UP:
		return digger.upSprites
	case DOWN:
		return digger.downSprites
	default:
		return digger.rightSprites
	}
}

func (digger *Digger) getCollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(digger.posX + (CELL_WIDTH-digger.width)/2 - CELL_WIDTH/2 - DIGGER_INNER_OFFSET_X + collisionOffsetsMap[digger.direction].X),
		Y:      float32(digger.posY + (CELL_WIDTH-digger.height)/2 - CELL_HEIGHT/2 - DIGGER_INNER_OFFSET_Y + collisionOffsetsMap[digger.direction].Y),
		Width:  float32(digger.width + collisionSizeMap[digger.direction].W),
		Height: float32(digger.height + collisionSizeMap[digger.direction].H),
	}
}

var (
	renderOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 0},
		RIGHT: {0, 0},
		UP:    {-1, 0},
		DOWN:  {1, 1},
	}
	collisionOffsetsMap = map[Direction]Pos{
		LEFT:  {0, 2},
		RIGHT: {0, 2},
		UP:    {0, 0},
		DOWN:  {2, 1},
	}
	collisionSizeMap = map[Direction]WidthHeight{
		LEFT:  {0, -2},
		RIGHT: {0, -2},
		UP:    {-2, 0},
		DOWN:  {-2, 0},
	}
)

func (digger *Digger) Render(drawTarget rl.RenderTexture2D) {
	sprites := digger.getSprites()
	rl.BeginTextureMode(drawTarget)
	rl.DrawTexture(
		sprites[digger.spritePointer].texture,
		digger.posX-CELL_WIDTH/2-DIGGER_INNER_OFFSET_X+renderOffsetsMap[digger.direction].X,
		digger.posY-CELL_HEIGHT/2-DIGGER_INNER_OFFSET_Y+renderOffsetsMap[digger.direction].Y,
		rl.White)
	//rl.DrawRectangleLinesEx(digger.getCollisionRec(), 1.0, TransparentYellow)
	rl.EndTextureMode()
}
