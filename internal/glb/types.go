package glb

type Direction int64
type FireState int64

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

const (
	MOVING FireState = iota
	STOPPED
	FINISHED
)
