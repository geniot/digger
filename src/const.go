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

	menuSceneKey = iota
	gameSceneKey
	controlsSceneKey
)

var (
	TransparentYellow = rl.NewColor(253, 249, 0, 100)
	TransparentRed    = rl.NewColor(230, 41, 55, 100)
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

	LEVEL_DATA = [][]string{
		{
			"S   B     HHHHS",
			"V  CC  C  V B  ",
			"VB CC  C  V    ",
			"V  CCB CB V CCC",
			"V  CC  C  V CCC",
			"HH CC  C  V CCC",
			" V    B B V    ",
			" HHHH     V    ",
			"C   V     V   C",
			"CC  HHHHHHH  CC",
		},
		{
			"SHHHHH  B B  HS",
			" CC  V       V ",
			" CC  V CCCCC V ",
			"BCCB V CCCCC V ",
			"CCCC V       V ",
			"CCCC V B  HHHH ",
			" CC  V CC V    ",
			" BB  VCCCCV CC ",
			"C    V CC V CC ",
			"CC   HHHHHH    ",
		},
		{
			"SHHHHB B BHHHHS",
			"CC  V C C V BB ",
			"C   V C C V CC ",
			" BB V C C VCCCC",
			"CCCCV C C VCCCC",
			"CCCCHHHHHHH CC ",
			" CC  C V C  CC ",
			" CC  C V C     ",
			"C    C V C    C",
			"CC   C H C   CC",
		},
		{
			"SHBCCCCBCCCCBHS",
			"CV  CCCCCCC  VC",
			"CHHH CCCCC HHHC",
			"C  V  CCC  V  C",
			"   HHH C HHH   ",
			"  B  V B V  B  ",
			"  C  VCCCV  C  ",
			" CCC HHHHH CCC ",
			"CCCCC CVC CCCCC",
			"CCCCC CHC CCCCC",
		},
		{
			"SHHHHHHHHHHHHHS",
			"VBCCCCBVCCCCCCV",
			"VCCCCCCV CCBC V",
			"V CCCC VCCBCCCV",
			"VCCCCCCV CCCC V",
			"V CCCC VBCCCCCV",
			"VCCBCCCV CCCC V",
			"V CCBC VCCCCCCV",
			"VCCCCCCVCCCCCCV",
			"HHHHHHHHHHHHHHH",
		},
		{
			"SHHHHHHHHHHHHHS",
			"VCBCCV V VCCBCV",
			"VCCC VBVBV CCCV",
			"VCCCHH V HHCCCV",
			"VCC V CVC V CCV",
			"VCCHH CVC HHCCV",
			"VC V CCVCC V CV",
			"VCHHBCCVCCBHHCV",
			"VCVCCCCVCCCCVCV",
			"HHHHHHHHHHHHHHH",
		},
		{
			"SHCCCCCVCCCCCHS",
			" VCBCBCVCBCBCV ",
			"BVCCCCCVCCCCCVB",
			"CHHCCCCVCCCCHHC",
			"CCV CCCVCCC VCC",
			"CCHHHCCVCCHHHCC",
			"CCCCV CVC VCCCC",
			"CCCCHH V HHCCCC",
			"CCCCCV V VCCCCC",
			"CCCCCHHHHHCCCCC",
		},
		{
			"HHHHHHHHHHHHHHS",
			"V CCBCCCCCBCC V",
			"HHHCCCCBCCCCHHH",
			"VBV CCCCCCC VBV",
			"VCHHHCCCCCHHHCV",
			"VCCBV CCC VBCCV",
			"VCCCHHHCHHHCCCV",
			"VCCCC V V CCCCV",
			"VCCCCCV VCCCCCV",
			"HHHHHHHHHHHHHHH",
		},
	}
)
