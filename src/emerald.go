package main

type Emerald struct {
	emeraldsPool *EmeraldsPool
}

func NewEmerald(emeraldsPool *EmeraldsPool) *Emerald {
	emerald := &Emerald{}
	emerald.emeraldsPool = emeraldsPool
	return emerald
}
