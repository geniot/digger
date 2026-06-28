//go:build !wasm
// +build !wasm

package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	application := NewApplication()

	accumulator := float64(0)
	lastTime := float64(0)

	tick := int64(0)
	for !application.ShouldExit() {
		t := rl.GetTime()
		dt := t - lastTime
		lastTime = t
		application.ProcessInput()
		for accumulator += dt; accumulator > TICK; accumulator -= TICK {
			application.Update(tick)
			tick++
		}
		application.Render()
	}
	application.Exit()
}
