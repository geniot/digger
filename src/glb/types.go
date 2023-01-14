package glb

import "github.com/veandco/go-sdl2/sdl"

//https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

type Direction int64

type DiggerState int64
type MonsterState int64
type FireState int64
type BagState int64

type Point struct {
	X, Y int32
}

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

const (
	DIGGER_ALIVE DiggerState = iota
	DIGGER_DIE
	DIGGER_GRAVE
)

const (
	MONSTER_NOBBIN MonsterState = iota
	MONSTER_HOBBIN
	MONSTER_DIE
)

const (
	FIRE_MOVING FireState = iota
	FIRE_STOPPED
	FIRE_FINISHED
)

const (
	BAG_SET BagState = iota
	BAG_PUSHED
	BAG_HOLD
	BAG_MOVING
	BAG_WOBBLE
	BAG_FALLING
	BAG_GOLD_FALLING
	BAG_GOLD
)

type SurfTexture struct {
	W int32
	H int32
	T *sdl.Texture
	S *sdl.Surface
}
