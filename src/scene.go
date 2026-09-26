package main

type Scene interface {
	ProcessInput()
	Update(tick int64)
	Render()
	ShouldExit() bool
}
