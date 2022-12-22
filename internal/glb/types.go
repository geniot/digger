package glb

//https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

type Direction int64
type FireState int64
type BagState int64

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

const (
	FIRE_MOVING FireState = iota
	FIRE_STOPPED
	FIRE_FINISHED
)

const (
	BAG_SET BagState = iota
	BAG_PUSHED
	BAG_MOVING
	BAG_SHAKING
	BAG_FALLING
	BAG_GOLD
)
