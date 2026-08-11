package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TICK               float64 = 1.0 / 200.0
	SPRITE_UPDATE_RATE         = 18
	DIGGER_SPEED               = 4 //less is faster
	FIELD_WIDTH                = 320
	FIELD_HEIGHT               = 186

	SCREEN_LOGICAL_WIDTH  = int32(320)
	SCREEN_LOGICAL_HEIGHT = int32(240)
	CELLS_HORIZONTAL      = int32(15)
	CELLS_VERTICAL        = int32(10)
	CELL_WIDTH            = int32(20)
	CELL_HEIGHT           = int32(18)
	FIELD_OFFSET_X        = int32(10)
	FIELD_OFFSET_Y        = int32(2)
	DIGGER_INNER_OFFSET_X = int32(0)
	DIGGER_INNER_OFFSET_Y = int32(1)

	menuSceneKey = iota
	gameSceneKey
	controlsSceneKey
)

var (
	TransparentYellow = rl.NewColor(253, 249, 0, 100)
	TransparentRed    = rl.NewColor(230, 41, 55, 100)
	TransparentBlue   = rl.NewColor(41, 41, 253, 100)
)

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

// TSP button codes
const (
	noCode = iota
	upCode
	rightCode
	downCode
	leftCode
	xCode
	aCode
	bCode
	yCode
	l1Code
	l2Code
	r1Code
	r2Code
	selectCode
	menuCode
	startCode
)

const (
	winHeight = 720
	winWidth  = 1280
	gamePadId = int32(0)
)

var (
	//go:embed res/*
	resList embed.FS
)

var (
	ZERO_VECTOR2        = rl.Vector2{}
	CELL_CENTER_VECTOR2 = rl.Vector2{X: float32(CELL_WIDTH / 2), Y: float32(CELL_HEIGHT / 2)}
)

var (
	keysToDirectionsMap = map[int32]Direction{
		rl.KeyLeft:  LEFT,
		rl.KeyRight: RIGHT,
		rl.KeyUp:    UP,
		rl.KeyDown:  DOWN,
	}
)
