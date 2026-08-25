package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TICK             float64 = 1.0 / 200.0
	SpriteUpdateRate         = 18
	DiggerSpeed              = 4 //less is faster
	FireSpeed                = 3
	FieldWidth               = 320
	FieldHeight              = 186

	ScreenLogicalWidth  = int32(320)
	ScreenLogicalHeight = int32(240)
	CellsHorizontal     = int32(15)
	CellsVertical       = int32(10)
	CellWidth           = int32(20)
	CellHeight          = int32(18)
	FieldOffsetX        = int32(10)
	FieldOffsetY        = int32(2)

	menuSceneKey = iota
	gameSceneKey
	controlsSceneKey
)

var (
	TransparentYellow = rl.NewColor(253, 249, 0, 100)
	TransparentRed    = rl.NewColor(230, 41, 55, 100)
	TransparentBlue   = rl.NewColor(41, 41, 253, 100)
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
	ZERO_VECTOR2        = rl.Vector2{}
	CELL_CENTER_VECTOR2 = rl.Vector2{X: float32(CellWidth / 2), Y: float32(CellHeight / 2)}
)
